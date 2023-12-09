package test

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/paycrest/protocol/ent"
	"github.com/paycrest/protocol/ent/receiveaddress"
	db "github.com/paycrest/protocol/storage"
	"github.com/paycrest/protocol/types"
	"github.com/paycrest/protocol/utils/crypto"

	"github.com/paycrest/protocol/services/contracts"
	cryptoUtils "github.com/paycrest/protocol/utils/crypto"
)

// SetUpTestBlockchain sets up a connection to a local Ethereum blockchain.
func SetUpTestBlockchain() (types.RPCClient, error) {
	// Connect to local ethereum client
	client, err := types.NewEthClient("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client, nil
}

// deployERC20Contract deploys an ERC20 contract with the provided parameters.
// It returns the address of the deployed contract, the transaction object for the deployment,
// an instance of the deployed contract, and any error that occurred.
func DeployERC20Contract(client types.RPCClient) (*common.Address, error) {
	// Prepare the deployment
	auth, err := prepareDeployment(client)
	if err != nil {
		return nil, err
	}

	initialSupply := new(big.Int)
	initialSupply.SetString("2000000000000000000000", 10) // Initial supply of 2,000 tokens (token decimal is 18)

	// Deploy the contract
	address, tx, _, err := contracts.DeployTestToken(auth, client.(bind.ContractBackend), initialSupply)
	if err != nil {
		return nil, err
	}

	_, err = bind.WaitMined(context.Background(), client.(bind.DeployBackend), tx)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}

	client.Commit()

	return &address, nil
}

// DeployEIP4337FactoryContract deploys an EIP-4337 factory contract.
func DeployEIP4337FactoryContract(client types.RPCClient) (common.Address, error) {
	// Prepare the deployment
	auth, err := prepareDeployment(client)
	if err != nil {
		return common.Address{}, err
	}

	// Deploy the contract
	address, tx, _, err := contracts.DeploySimpleAccountFactory(
		auth, client.(bind.ContractBackend), common.HexToAddress("0x8091bDf8fa8762414007C08cF642D33697C8bF51"))
	if err != nil {
		return common.Address{}, err
	}

	_, err = bind.WaitMined(context.Background(), client.(bind.DeployBackend), tx)
	if err != nil {
		log.Fatalf("Tx receipt failed: %v", err)
	}

	client.Commit()

	return address, nil
}

// FundAddressWithTestToken funds an amount of a test ERC20 token from the owner account
func FundAddressWithTestToken(client types.RPCClient, token common.Address, amount *big.Int, address common.Address) error {
	// Get master account
	_, privateKey, _ := crypto.GenerateAccountFromIndex(0)

	// Create a new instance of the TestToken contract
	testToken, err := contracts.NewTestToken(token, client.(bind.ContractBackend))
	if err != nil {
		return err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	// Transfer test token to the address
	tx, err := testToken.Transfer(auth, address, amount)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(context.Background(), client.(bind.DeployBackend), tx)
	if err != nil {
		return err
	}

	return nil
}

// prepareDeployment prepares the deployment of a contract.
func prepareDeployment(client types.RPCClient) (*bind.TransactOpts, error) {
	// Get master account
	fromAddress, privateKey, _ := crypto.GenerateAccountFromIndex(0)

	// Configure the transaction
	ctx := context.Background()

	nonce, err := client.PendingNonceAt(ctx, *fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	callMsg := ethereum.CallMsg{
		From:     *fromAddress,
		To:       nil,
		Gas:      0,             // Set to 0 to let the client estimate
		GasPrice: big.NewInt(0), // Set to 0 for gas price estimation
		Value:    big.NewInt(0),
		Data:     []byte{},
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return nil, err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = gasLimit + uint64(9000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// CreateSmartAccount function generates and saves a new EIP-4337 smart contract account address
func CreateSmartAccount(ctx context.Context, client types.RPCClient) (*ent.ReceiveAddress, error) {

	// Initialize contract factory
	factory, err := DeployEIP4337FactoryContract(client)
	if err != nil {
		return nil, err
	}

	factoryInstance, err := contracts.NewSimpleAccountFactory(factory, client.(bind.ContractBackend))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize factory contract: %w", err)
	}

	// Get master account
	ownerAddress, privateKey, _ := crypto.GenerateAccountFromIndex(0)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	// Create a new big.Int from the hash
	salt := new(big.Int).SetBytes(nonce)
	callOpts := &bind.CallOpts{
		Pending: true,
		From:    auth.From,
		Context: context.Background(),
	}

	// Generate address
	smartAccountAddress, err := factoryInstance.GetAddress(callOpts, *ownerAddress, salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate address: %w", err)
	}
	// Deploy smart account
	createTx, err := factoryInstance.CreateAccount(auth, *ownerAddress, salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate address: %w", err)
	}
	_, err = bind.WaitMined(context.Background(), client.(bind.DeployBackend), createTx)
	if err != nil {
		log.Fatalf("createTx receipt failed: %v", err)
	}
	saltEncrypted, err := cryptoUtils.EncryptPlain([]byte(salt.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt salt: %w", err)
	}

	// Save address in db
	receiveAddress, err := db.Client.ReceiveAddress.
		Create().
		SetAddress(smartAccountAddress.Hex()).
		SetSalt(saltEncrypted).
		SetStatus(receiveaddress.StatusUnused).
		SetValidUntil(time.Now().Add(time.Millisecond * 5)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to save address: %w", err)
	}

	return receiveAddress, nil
}
