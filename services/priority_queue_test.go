package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/paycrest/protocol/ent"
	"github.com/paycrest/protocol/ent/enttest"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/provisionbucket"
	"github.com/ethereum/go-ethereum/crypto"
	db "github.com/paycrest/protocol/storage"
	"github.com/paycrest/protocol/types"
	cryptoUtils "github.com/paycrest/protocol/utils/crypto"
	"github.com/paycrest/protocol/utils/test"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"encoding/hex"
)

var testCtxForPQ = struct {
	publicProvider              *ent.User
	publicProviderProfile       *ent.ProviderProfile
	publicProviderProfileAPIKey *ent.APIKey
	privateProviderProfile      *ent.ProviderProfile
	currency                    *ent.FiatCurrency
	client                      types.RPCClient
	token                       *ent.Token
	minAmount                   decimal.Decimal
	maxAmount                   decimal.Decimal
	bucket                      *ent.ProvisionBucket
}{}

var (
	fromAddress, privateKey, _ = cryptoUtils.GenerateAccountFromIndex(0)
)

func setupForPQ() error {
	// Set up test data
	testCtxForPQ.maxAmount = decimal.NewFromFloat(10000)
	testCtxForPQ.minAmount = decimal.NewFromFloat(1)

	backend, err := test.SetUpTestBlockchain()
	if err != nil {
		return err
	}
	token, err := test.CreateERC20Token(backend, map[string]interface{}{})
	if err != nil {
		return err
	}
	testCtxForPQ.token = token

	user, err := test.CreateTestUser(map[string]interface{}{
		"scope": "provider",
		"email": "providerjohndoe@test.com",
	})
	if err != nil {
		return err
	}
	testCtxForPQ.publicProvider = user

	currency, err := test.CreateTestFiatCurrency(map[string]interface{}{
		"code":        "KES",
		"short_name":  "Shilling",
		"decimals":    2,
		"symbol":      "KSh",
		"name":        "Kenyan Shilling",
		"market_rate": 550.0,
	})
	if err != nil {
		return err
	}
	testCtxForPQ.currency = currency

	publicProviderProfile, err := test.CreateTestProviderProfile(map[string]interface{}{
		"user_id":         testCtxForPQ.publicProvider.ID,
		"currency_id":     currency.ID,
		"host_identifier": "https://example2.com",
	})
	if err != nil {
		return err
	}
	apiKeyService := NewAPIKeyService()
	apiKey, _, err := apiKeyService.GenerateAPIKey(
		context.Background(),
		nil,
		nil,
		publicProviderProfile,
	)
	if err != nil {
		return err
	}
	testCtxForPQ.publicProviderProfileAPIKey = apiKey
	tokens, err := test.AddProviderOrderTokenToProvider(
		map[string]interface{}{
			"fixed_conversion_rate":    decimal.NewFromFloat(100),
			"conversion_rate_type":     "fixed",
			"floating_conversion_rate": decimal.NewFromFloat(1.0),
			"max_order_amount":         decimal.NewFromFloat(1000),
			"min_order_amount":         decimal.NewFromFloat(1.0),
			"tokenSymbol":              token.Symbol,
			"provider":                 publicProviderProfile,
			"addresses": []map[string]string{
				{
					"address": fromAddress.Hex(),
					"network": "localhost",
				},
			},
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("provider tokens", tokens)
	testCtxForPQ.publicProviderProfile = publicProviderProfile

	bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
		"provider_id": publicProviderProfile.ID,
		"min_amount":  decimal.NewFromFloat(1),
		"max_amount":  decimal.NewFromFloat(10000.0),
		"currency_id": currency.ID,
	})
	if err != nil {
		return err
	}
	testCtxForPQ.bucket = bucket

	privateProvider, err := test.CreateTestUser(map[string]interface{}{
		"scope": "provider",
		"email": "private@test.com",
	})
	if err != nil {
		return err
	}

	privateProviderProfile, err := test.CreateTestProviderProfile(map[string]interface{}{
		"currency_id":     currency.ID,
		"visibility_mode": "private",
		"user_id":         privateProvider.ID,
	})
	if err != nil {
		return err
	}
	testCtxForPQ.privateProviderProfile = privateProviderProfile

	_, err = test.CreateTestProvisionBucket(map[string]interface{}{
		"provider_id": privateProviderProfile.ID,
		"min_amount":  testCtxForPQ.minAmount,
		"max_amount":  testCtxForPQ.maxAmount,
		"currency_id": currency.ID,
	})
	if err != nil {
		return err
	}

	// Set up payment order
	_, err = test.CreateTestLockPaymentOrder(map[string]interface{}{
		"provider":   privateProviderProfile,
		"tokenID":    testCtxForPQ.token.ID,
		"gateway_id": "order-12345",
	})
	if err != nil {
		return err
	}
	_, err = test.CreateTestLockPaymentOrder(map[string]interface{}{
		"provider": publicProviderProfile,
		"tokenID":  testCtxForPQ.token.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func TestPriorityQueueTest(t *testing.T) {
	// Set up test database client
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer redisClient.Close()

	db.RedisClient = redisClient
	{

		err := redisClient.FlushAll(context.Background()).Err()
		assert.NoError(t, err)
	}

	db.Client = client

	// Setup test data
	err := setupForPQ()
	assert.NoError(t, err)

	service := NewPriorityQueueService()
	t.Run("TestGetProvisionBuckets", func(t *testing.T) {
		buckets, err := service.GetProvisionBuckets(context.Background())
		assert.NoError(t, err)
		assert.Greater(t, len(buckets), 0)
	})

	t.Run("TestCreatePriorityQueueForBucket", func(t *testing.T) {
		ctx := context.Background()
		bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
			"provider_id": testCtxForPQ.publicProviderProfile.ID,
			"min_amount":  testCtxForPQ.minAmount,
			"max_amount":  testCtxForPQ.maxAmount,
			"currency_id": testCtxForPQ.currency.ID,
		})
		assert.NoError(t, err)

		_bucket, err := db.Client.ProvisionBucket.
			Query().
			Where(provisionbucket.IDEQ(bucket.ID)).
			WithCurrency().
			WithProviderProfiles().
			Only(ctx)
		assert.NoError(t, err)

		service.CreatePriorityQueueForBucket(ctx, _bucket)

		redisKey := fmt.Sprintf("bucket_%s_%s_%s", _bucket.Edges.Currency.Code, testCtxForPQ.minAmount, testCtxForPQ.maxAmount)

		data, err := db.RedisClient.LRange(ctx, redisKey, 0, -1).Result()
		assert.NoError(t, err)
		assert.Equal(t, len(data), 1)
		assert.Contains(t, data[0], testCtxForPQ.publicProviderProfile.ID)
	})

	t.Run("TestProcessBucketQueues", func(t *testing.T) {
		err = service.ProcessBucketQueues()
		assert.NoError(t, err)

		redisKey := fmt.Sprintf("bucket_%s_%s_%s", testCtxForPQ.currency.Code, testCtxForPQ.minAmount, testCtxForPQ.maxAmount)

		data, err := db.RedisClient.LRange(context.Background(), redisKey, 0, -1).Result()
		assert.NoError(t, err)
		assert.Equal(t, len(data), 1)
	})

	t.Run("TestAssignLockPaymentOrder", func(t *testing.T) {

		bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
			"provider_id": testCtxForPQ.privateProviderProfile.ID,
			"min_amount":  testCtxForPQ.minAmount,
			"max_amount":  testCtxForPQ.maxAmount,
			"currency_id": testCtxForPQ.currency.ID,
		})
		assert.NoError(t, err)
		_order, err := test.CreateTestLockPaymentOrder(map[string]interface{}{
			"provider":   testCtxForPQ.publicProviderProfile,
			"tokenID":    testCtxForPQ.token.ID,
			"gateway_id": "order-1",
		})
		assert.NoError(t, err)

		_, err = test.AddProvisionBucketToLockPaymentOrder(_order, bucket.ID)
		assert.NoError(t, err)

		err = db.RedisClient.RPush(context.Background(), fmt.Sprintf("order_exclude_list_%s", _order.ID), testCtxForPQ.publicProviderProfile.ID).Err()
		assert.NoError(t, err)

		order, err := db.Client.LockPaymentOrder.
			Query().
			Where(lockpaymentorder.IDEQ(_order.ID)).
			WithProvisionBucket(func(pb *ent.ProvisionBucketQuery) {
				pb.WithCurrency()
			}).
			WithToken().
			Only(context.Background())

		assert.NoError(t, err)

		err = service.AssignLockPaymentOrder(context.Background(), types.LockPaymentOrderFields{
			ID:                order.ID,
			Token:             testCtxForPQ.token,
			GatewayID:         order.GatewayID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			Memo:              order.Memo,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		})
		assert.NoError(t, err)
	})

	t.Run("TestGetProviderRate", func(t *testing.T) {
		rate, err := service.GetProviderRate(context.Background(), testCtxForPQ.publicProviderProfile, testCtxForPQ.token.Symbol)
		assert.NoError(t, err)
		_rate, ok := rate.Float64()
		assert.True(t, ok)
		assert.Equal(t, _rate, float64(100))
	})

	t.Run("TestSendOrderRequest", func(t *testing.T) {
		// setup httpmock
		httpmock.Activate()
		defer httpmock.Deactivate()

		httpmock.RegisterResponder("POST", testCtxForPQ.publicProviderProfile.HostIdentifier+"/new_order",
			func(r *http.Request) (*http.Response, error) {
				bytes, err := io.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}

				var orderRequestData map[string]interface{}
				if err := json.Unmarshal(bytes, &orderRequestData); err != nil {
					log.Fatal(err)
				}

				messageHash := test.CreateMessageHash(orderRequestData)

				// signature, err := utils.PersonalSign(messageHash.Hex(), privateKey)
				signatureBytes, err := crypto.Sign(messageHash.Bytes(), privateKey)
				if err != nil {
					log.Fatal(err)
				}
				
				responseData := map[string]string{
					"signature": hex.EncodeToString(signatureBytes),
				}
				responseBytes, err := json.Marshal(responseData)
				if err != nil {
					log.Fatal(err)
				}
				return httpmock.NewBytesResponse(http.StatusOK, responseBytes), nil
			},
		)
		
		bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
			"provider_id": testCtxForPQ.privateProviderProfile.ID,
			"min_amount":  testCtxForPQ.minAmount,
			"max_amount":  testCtxForPQ.maxAmount,
			"currency_id": testCtxForPQ.currency.ID,
		})
		assert.NoError(t, err)
		_order, err := test.CreateTestLockPaymentOrder(map[string]interface{}{
			"provider":   testCtxForPQ.publicProviderProfile,
			"tokenID":    testCtxForPQ.token.ID,
			"gateway_id": "order-1234",
		})
		assert.NoError(t, err)

		_, err = test.AddProvisionBucketToLockPaymentOrder(_order, bucket.ID)
		assert.NoError(t, err)

		err = db.RedisClient.RPush(context.Background(), fmt.Sprintf("order_exclude_list_%s", _order.ID), testCtxForPQ.publicProviderProfile.ID).Err()
		assert.NoError(t, err)

		order, err := db.Client.LockPaymentOrder.
			Query().
			Where(lockpaymentorder.IDEQ(_order.ID)).
			WithProvisionBucket(func(pb *ent.ProvisionBucketQuery) {
				pb.WithCurrency()
			}).
			WithToken().
			Only(context.Background())

		assert.NoError(t, err)

		err = service.sendOrderRequest(context.Background(), types.LockPaymentOrderFields{
			ID:                order.ID,
			ProviderID:        testCtxForPQ.publicProviderProfile.ID,
			Token:             testCtxForPQ.token,
			Network:           &ent.Network{Identifier: "localhost"},
			GatewayID:         order.GatewayID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			Memo:              order.Memo,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		})
		assert.NoError(t, err)

		t.Run("TestNotifyProvider", func(t *testing.T) {

			err := service.notifyProvider(
				context.Background(), 
				map[string]interface{}{
						"providerId": testCtxForPQ.publicProviderProfile.ID,
						"data":       "test",
					},
				"localhost",
			)
			assert.NoError(t, err)
		})
	})

	// TODO: move these tests to tasks_test.go
	// t.Run("TestNoErrorFunctions", func(t *testing.T) {

	// 	t.Run("TestReassignUnfulfilledLockOrders", func(t *testing.T) {

	// 		bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
	// 			"provider_id": testCtxForPQ.privateProviderProfile.ID,
	// 			"min_amount":  testCtxForPQ.minAmount,
	// 			"max_amount":  testCtxForPQ.maxAmount,
	// 			"currency_id": testCtxForPQ.currency.ID,
	// 		})
	// 		assert.NoError(t, err)
	// 		_order, err := test.CreateTestLockPaymentOrder(map[string]interface{}{
	// 			"provider": testCtxForPQ.publicProviderProfile,
	// 			"tokenID":  testCtxForPQ.token.ID,
	// 			"status":   lockpaymentorder.StatusProcessing.String(),
	// 		})
	// 		assert.NoError(t, err)

	// 		_, err = test.AddProvisionBucketToLockPaymentOrder(_order, bucket.ID)
	// 		assert.NoError(t, err)

	// 		service.ReassignUnfulfilledLockOrders()

	// 		order, err := db.Client.LockPaymentOrder.
	// 			Query().
	// 			Where(lockpaymentorder.IDEQ(_order.ID)).Only(context.Background())
	// 		assert.NoError(t, err)

	// 		//validate the ReassignUnfulfilledLockOrders updated the UnfulfilledLockOrder
	// 		assert.True(t, _order.UpdatedAt.Before(order.UpdatedAt))
	// 	})

	// 	t.Run("TestReassignStaleOrderRequest", func(t *testing.T) {
	// 		bucket, err := test.CreateTestProvisionBucket(map[string]interface{}{
	// 			"provider_id": testCtxForPQ.privateProviderProfile.ID,
	// 			"min_amount":  testCtxForPQ.minAmount,
	// 			"max_amount":  testCtxForPQ.maxAmount,
	// 			"currency_id": testCtxForPQ.currency.ID,
	// 		})
	// 		assert.NoError(t, err)
	// 		_order, err := test.CreateTestLockPaymentOrder(map[string]interface{}{
	// 			"provider":  testCtxForPQ.privateProviderProfile,
	// 			"tokenID":   testCtxForPQ.token.ID,
	// 			"status":    lockpaymentorder.StatusProcessing.String(),
	// 			"updatedAt": time.Now().Add(-5 * time.Minute),
	// 		})
	// 		assert.NoError(t, err)

	// 		orderKey := fmt.Sprintf("order_exclude_list_%s", _order.ID)
	// 		_, err = db.RedisClient.RPush(context.Background(), orderKey, testCtxForPQ.privateProviderProfile.ID).Result()
	// 		assert.NoError(t, err)

	// 		_, err = test.AddProvisionBucketToLockPaymentOrder(_order, bucket.ID)
	// 		assert.NoError(t, err)

	// 		service.ReassignUnfulfilledLockOrders()

	// 		// Create Channel
	// 		orderRequestChan := make(chan *redis.Message, 1)
	// 		orderRequestChan <- &redis.Message{Payload: _order.ID.String() + "_" + "TEST"}
	// 		service.ReassignStaleOrderRequest(context.Background(), orderRequestChan)

	// 		order, err := db.Client.LockPaymentOrder.
	// 			Query().
	// 			Where(lockpaymentorder.IDEQ(_order.ID)).Only(context.Background())
	// 		assert.NoError(t, err)
	// 		// validate the StaleOrderRequest updated the StaleOrderRequest
	// 		assert.True(t, _order.UpdatedAt.Before(order.UpdatedAt))

	// 		// Close channel
	// 		close(orderRequestChan)
	// 	})
	// })
}
