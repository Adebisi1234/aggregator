package sender

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/paycrest/protocol/config"
	"github.com/paycrest/protocol/ent"
	db "github.com/paycrest/protocol/storage"

	"github.com/paycrest/protocol/ent/network"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/paycrest/protocol/ent/token"
	svc "github.com/paycrest/protocol/services"
	"github.com/paycrest/protocol/types"
	u "github.com/paycrest/protocol/utils"
	"github.com/paycrest/protocol/utils/logger"
	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
)

// SenderController is a controller type for sender endpoints
type SenderController struct {
	indexerService        *svc.IndexerService
	receiveAddressService *svc.ReceiveAddressService
	orderService          *svc.OrderService
}

// NewSenderController creates a new instance of SenderController
func NewSenderController(indexer svc.Indexer) *SenderController {
	var indexerService *svc.IndexerService

	if indexer != nil {
		indexerService = svc.NewIndexerService(indexer)
	} else {
		indexerService = svc.NewIndexerService(nil)
	}

	return &SenderController{
		indexerService:        indexerService,
		receiveAddressService: svc.NewReceiveAddressService(),
		orderService:          svc.NewOrderService(),
	}
}

// CreatePaymentOrder controller creates a payment order
func (ctrl *SenderController) CreatePaymentOrder(ctx *gin.Context) {
	var payload types.NewPaymentOrderPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusBadRequest, "error",
			"Failed to validate payload", u.GetErrorData(err))
		return
	}

	// Get sender profile from the context
	senderCtx, ok := ctx.Get("sender")
	if !ok {
		u.APIResponse(ctx, http.StatusUnauthorized, "error", "Invalid API key", nil)
		return
	}
	sender := senderCtx.(*ent.SenderProfile)

	conf := config.ServerConfig()

	if !sender.IsActive && !conf.Debug {
		u.APIResponse(ctx, http.StatusForbidden, "error", "Your account is not active", nil)
		return
	}

	// Get token from DB
	token, err := db.Client.Token.
		Query().
		Where(
			token.SymbolEQ(payload.Token),
			token.HasNetworkWith(network.IdentifierEQ(payload.Network)),
		).
		WithNetwork().
		Only(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusBadRequest, "error",
			"Provided token is not supported", nil)
		return
	}

	// Generate receive address
	receiveAddress, err := ctrl.receiveAddressService.CreateSmartAccount(ctx, nil, nil)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to initiate payment order", nil)
		return
	}

	// Create payment order and recipient in a transaction
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to initiate payment order", nil)
		return
	}

	// Create payment order
	paymentOrder, err := tx.PaymentOrder.
		Create().
		SetSenderProfile(sender).
		SetAmount(payload.Amount).
		SetAmountPaid(decimal.NewFromInt(0)).
		SetToken(token).
		SetLabel(payload.Label).
		SetRate(payload.Rate).
		SetReceiveAddress(receiveAddress).
		SetReceiveAddressText(receiveAddress.Address).
		Save(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to initiate payment order", nil)
		_ = tx.Rollback()
		return
	}

	// Create payment order recipient
	_, err = tx.PaymentOrderRecipient.
		Create().
		SetInstitution(payload.Recipient.Institution).
		SetAccountIdentifier(payload.Recipient.AccountIdentifier).
		SetAccountName(payload.Recipient.AccountName).
		SetProviderID(payload.Recipient.ProviderID).
		SetMemo(payload.Recipient.Memo).
		SetPaymentOrder(paymentOrder).
		Save(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to initiate payment order", nil)
		_ = tx.Rollback()
		return
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to initiate payment order", nil)
		return
	}

	// Start a background process to index token transfers to the receive address
	go ctrl.indexerService.RunIndexERC20Transfer(ctx, receiveAddress)

	paymentOrderAmount, _ := paymentOrder.Amount.Float64()

	u.APIResponse(ctx, http.StatusCreated, "success", "Payment order initiated successfully",
		&types.ReceiveAddressResponse{
			ID:             paymentOrder.ID,
			Amount:         paymentOrderAmount,
			Network:        token.Edges.Network.Identifier,
			ReceiveAddress: paymentOrder.ReceiveAddressText,
		})
}

// GetPaymentOrderByID controller fetches a payment order by ID
func (ctrl *SenderController) GetPaymentOrderByID(ctx *gin.Context) {
	// Get order ID from the URL
	orderID := ctx.Param("id")

	// Convert order ID to UUID
	id, err := uuid.Parse(orderID)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusBadRequest, "error",
			"Invalid order ID", nil)
		return
	}

	// Get sender profile from the context
	senderCtx, ok := ctx.Get("sender")
	if !ok {
		u.APIResponse(ctx, http.StatusUnauthorized, "error", "Invalid API key", nil)
		return
	}
	sender := senderCtx.(*ent.SenderProfile)

	// Fetch payment order from the database
	paymentOrder, err := db.Client.PaymentOrder.
		Query().
		Where(
			paymentorder.IDEQ(id),
			paymentorder.HasSenderProfileWith(senderprofile.IDEQ(sender.ID)),
		).
		WithRecipient().
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		Only(ctx)

	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusNotFound, "error",
			"Payment order not found", nil)
		return
	}

	u.APIResponse(ctx, http.StatusOK, "success", "The order has been successfully retrieved", &types.PaymentOrderResponse{
		ID:      paymentOrder.ID,
		Amount:  paymentOrder.Amount,
		Rate:    paymentOrder.Rate,
		Network: paymentOrder.Edges.Token.Edges.Network.Identifier,
		Recipient: types.PaymentOrderRecipient{
			Institution:       paymentOrder.Edges.Recipient.Institution,
			AccountIdentifier: paymentOrder.Edges.Recipient.AccountIdentifier,
			AccountName:       paymentOrder.Edges.Recipient.AccountName,
			ProviderID:        paymentOrder.Edges.Recipient.ProviderID,
			Memo:              paymentOrder.Edges.Recipient.Memo,
		},
		Label:     paymentOrder.Label,
		CreatedAt: paymentOrder.CreatedAt,
		UpdatedAt: paymentOrder.UpdatedAt,
		TxHash:    paymentOrder.TxHash,
		Status:    paymentOrder.Status,
	})
}

// GetPaymentOrders controller fetches all payment orders
func (ctrl *SenderController) GetPaymentOrders(ctx *gin.Context) {
	// get page and pageSize query params
	page, pageSize := u.Paginate(ctx)

	// Get sender profile from the context
	senderCtx, ok := ctx.Get("sender")
	if !ok {
		u.APIResponse(ctx, http.StatusUnauthorized, "error", "Invalid API key", nil)
		return
	}
	sender := senderCtx.(*ent.SenderProfile)

	paymentOrderQuery := db.Client.PaymentOrder.
		Query()

	// Filter by status
	statusMap := map[string]paymentorder.Status{
		"initiated": paymentorder.StatusInitiated,
		"pending":   paymentorder.StatusPending,
		"settled":   paymentorder.StatusSettled,
		"cancelled": paymentorder.StatusCancelled,
		"failed":    paymentorder.StatusFailed,
		"refunded":  paymentorder.StatusRefunded,
	}

	statusQueryParam := ctx.Query("status")

	if status, ok := statusMap[statusQueryParam]; ok {
		paymentOrderQuery = paymentOrderQuery.Where(
			paymentorder.HasSenderProfileWith(senderprofile.IDEQ(sender.ID)),
			paymentorder.StatusEQ(status),
		)
	} else {
		paymentOrderQuery = paymentOrderQuery.Where(
			paymentorder.HasSenderProfileWith(senderprofile.IDEQ(sender.ID)),
		)
	}

	// get ordering query param
	ordering := ctx.Query("ordering")
	order := ent.Desc(paymentorder.FieldCreatedAt)
	if ordering == "asc" {
		order = ent.Asc(paymentorder.FieldCreatedAt)
	}

	// Get all payment orders for the sender
	paymentOrders, err := paymentOrderQuery.
		WithRecipient().
		WithToken(func(tq *ent.TokenQuery) {
			tq.WithNetwork()
		}).
		Limit(pageSize).
		Offset(page).
		Order(order).
		All(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		u.APIResponse(ctx, http.StatusInternalServerError, "error",
			"Failed to fetch payment orders", nil)
		return
	}

	var response []types.PaymentOrderResponse

	for _, paymentOrder := range paymentOrders {
		response = append(response, types.PaymentOrderResponse{
			ID:      paymentOrder.ID,
			Amount:  paymentOrder.Amount,
			Rate:    paymentOrder.Rate,
			Network: paymentOrder.Edges.Token.Edges.Network.Identifier,
			Recipient: types.PaymentOrderRecipient{
				Institution:       paymentOrder.Edges.Recipient.Institution,
				AccountIdentifier: paymentOrder.Edges.Recipient.AccountIdentifier,
				AccountName:       paymentOrder.Edges.Recipient.AccountName,
				ProviderID:        paymentOrder.Edges.Recipient.ProviderID,
				Memo:              paymentOrder.Edges.Recipient.Memo,
			},
			Label:     paymentOrder.Label,
			CreatedAt: paymentOrder.CreatedAt,
			UpdatedAt: paymentOrder.UpdatedAt,
			TxHash:    paymentOrder.TxHash,
			Status:    paymentOrder.Status,
		})
	}

	u.APIResponse(ctx, http.StatusOK, "success", "Payment orders retrieved successfully", types.SenderPaymentOrderList{
		TotalRecords: len(response),
		Page:         page + 1,
		PageSize:     pageSize,
		Orders:       response,
	})
}
