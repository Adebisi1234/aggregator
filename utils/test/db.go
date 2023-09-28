package test

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent"
	"github.com/paycrest/paycrest-protocol/ent/lockpaymentorder"
	db "github.com/paycrest/paycrest-protocol/storage"
	"github.com/paycrest/paycrest-protocol/types"
	"github.com/shopspring/decimal"
)

// CreateTestUser creates a test user with default or custom values
func CreateTestUser(overrides map[string]string) (*ent.User, error) {

	// Default payload
	payload := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
		"email":     "johndoe@test.com",
		"password":  "password",
	}

	// Apply overrides
	for key, value := range overrides {
		payload[key] = value
	}

	// Create user
	user, err := db.Client.User.
		Create().
		SetFirstName(payload["firstName"]).
		SetLastName(payload["lastName"]).
		SetEmail(strings.ToLower(payload["email"])).
		SetPassword(payload["password"]).
		Save(context.Background())

	return user, err
}

// CreateTestToken creates a test token with default or custom values
func CreateTestToken(client types.RPCClient, overrides map[string]interface{}) (*ent.Token, error) {

	// Deploy ERC20 token contract
	tokenAddress, err := DeployERC20Contract(client)
	if err != nil {
		return nil, err
	}

	// Default payload
	payload := map[string]interface{}{
		"symbol":           "TST",
		"contract_address": tokenAddress.Hex(),
		"decimals":         18,
		"networkRPC":       "http://localhost:8545",
		"is_enabled":       true,
	}

	// Apply overrides
	for key, value := range overrides {
		payload[key] = value
	}

	// Create Network
	network, err := db.Client.Network.
		Create().
		SetIdentifier("polygon-mumbai").
		SetChainID(1337).
		SetRPCEndpoint(payload["networkRPC"].(string)).
		SetIsTestnet(true).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	// Create token
	token, err := db.Client.Token.
		Create().
		SetSymbol(payload["symbol"].(string)).
		SetContractAddress(payload["contract_address"].(string)).
		SetDecimals(int8(payload["decimals"].(int))).
		SetNetwork(network).
		SetIsEnabled(payload["is_enabled"].(bool)).
		Save(context.Background())

	return token, err
}

// CreateTestLockPaymentOrder creates a test LockPaymentOrder with default or custom values
func CreateTestLockPaymentOrder(overrides map[string]interface{}) (*ent.LockPaymentOrder, error) {

	// Default payload
	payload := map[string]interface{}{
		"id":                 uuid.New(),
		"order_id":           "order-123",
		"amount":             100.50,
		"rate":               750.0,
		"status":             "pending",
		"block_number":       12345,
		"institution":        "Test Bank",
		"account_identifier": "1234567890",
		"account_name":       "Test Account",
	}

	// Apply overrides
	for key, value := range overrides {
		payload[key] = value
	}

	// Create test token
	backend, _ := NewSimulatedBlockchain()
	token, _ := CreateTestToken(backend, nil)

	// Create LockPaymentOrder
	order, err := db.Client.LockPaymentOrder.
		Create().
		SetID(payload["id"].(uuid.UUID)).
		SetOrderID(payload["order_id"].(string)).
		SetAmount(decimal.NewFromFloat(payload["amount"].(float64))).
		SetRate(decimal.NewFromFloat(payload["rate"].(float64))).
		SetStatus(lockpaymentorder.Status(payload["status"].(string))).
		SetBlockNumber(int64(payload["block_number"].(int))).
		SetInstitution(payload["institution"].(string)).
		SetAccountIdentifier(payload["account_identifier"].(string)).
		SetAccountName(payload["account_name"].(string)).
		SetTokenID(token.ID).
		Save(context.Background())

	return order, err
}

// CreateTestLockOrderFulfillment creates a test LockOrderFulfillment with defaults or custom values
func CreateTestLockOrderFulfillment(overrides map[string]interface{}) (*ent.LockOrderFulfillment, error) {

	// Default payload
	payload := map[string]interface{}{
		"id":                uuid.New(),
		"tx_id":             "0x123...",
		"tx_receipt_image":  "https://picsum.photos/200",
		"confirmations":     0,
		"validation_errors": []string{},
	}

	// Apply overrides
	for key, value := range overrides {
		payload[key] = value
	}

	// Create lock order
	order, _ := CreateTestLockPaymentOrder(nil)

	// Create LockOrderFulfillment
	fulfillment, err := db.Client.LockOrderFulfillment.
		Create().
		SetID(payload["id"].(uuid.UUID)).
		SetTxID(payload["tx_id"].(string)).
		SetTxReceiptImage(payload["tx_receipt_image"].(string)).
		SetConfirmations(payload["confirmations"].(int)).
		SetValidationErrors(payload["validation_errors"].([]string)).
		SetOrderID(order.ID).
		Save(context.Background())

	return fulfillment, err
}

// CreateTestValidatorProfile creates a test ValidatorProfile with defaults or custom values
func CreateTestValidatorProfile(overrides map[string]interface{}) (*ent.ValidatorProfile, error) {

	// Default payload
	payload := map[string]interface{}{
		"id":             uuid.New(),
		"wallet_address": "0x000000000000000000000000000000000000dEaD",
		"api_key":        nil,
	}

	// Apply overrides
	for key, value := range overrides {
		payload[key] = value
	}

	// Create ValidatorProfile
	profile, err := db.Client.ValidatorProfile.
		Create().
		SetID(payload["id"].(uuid.UUID)).
		SetWalletAddress(payload["wallet_address"].(string)).
		SetAPIKeyID(payload["api_key"].(uuid.UUID)).
		Save(context.Background())

	return profile, err
}
