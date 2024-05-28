package config

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

// OrderConfiguration type defines payment order configurations
type OrderConfiguration struct {
	OrderFulfillmentValidity         time.Duration
	ReceiveAddressValidity           time.Duration
	OrderRequestValidity             time.Duration
	GatewayContractAddress           common.Address
	GatewayContractAddressTron       string
	TronProApiKey                    string
	EntryPointContractAddress        common.Address
	BucketQueueRebuildInterval       int // in hours
	RefundCancellationCount          int
	PercentDeviationFromExternalRate decimal.Decimal
	PercentDeviationFromMarketRate   decimal.Decimal
	BundlerUrlEthereum               string
	PaymasterUrlEthereum             string
	BundlerUrlPolygon                string
	PaymasterUrlPolygon              string
	BundlerUrlBase                   string
	PaymasterUrlBase                 string
	BundlerUrlBSC                    string
	PaymasterUrlBSC                  string
	BundlerUrlArbitrum               string
	PaymasterUrlArbitrum             string
	ActiveAAService                  string
}

// OrderConfig sets the order configuration
func OrderConfig() *OrderConfiguration {
	viper.SetDefault("RECEIVE_ADDRESS_VALIDITY", 30)
	viper.SetDefault("ORDER_REQUEST_VALIDITY", 120)
	viper.SetDefault("ORDER_FULFILLMENT_VALIDITY", 10)
	viper.SetDefault("BUCKET_QUEUE_REBUILD_INTERVAL", 1)
	viper.SetDefault("REFUND_CANCELLATION_COUNT", 3)
	viper.SetDefault("NETWORK_FEE", 0.05)
	viper.SetDefault("PERCENT_DEVIATION_FROM_EXTERNAL_RATE", 0.01)
	viper.SetDefault("PERCENT_DEVIATION_FROM_MARKET_RATE", 0.1)
	viper.SetDefault("ACTIVE_AA_SERVICE", "STACKUP")

	return &OrderConfiguration{
		OrderFulfillmentValidity:         time.Duration(viper.GetInt("ORDER_FULFILLMENT_VALIDITY")) * time.Minute,
		ReceiveAddressValidity:           time.Duration(viper.GetInt("RECEIVE_ADDRESS_VALIDITY")) * time.Minute,
		OrderRequestValidity:             time.Duration(viper.GetInt("ORDER_REQUEST_VALIDITY")) * time.Second,
		GatewayContractAddress:           common.HexToAddress(viper.GetString("GATEWAY_CONTRACT_ADDRESS")),
		GatewayContractAddressTron:       viper.GetString("GATEWAY_CONTRACT_ADDRESS_TRON"),
		TronProApiKey:                    viper.GetString("TRON_PRO_API_KEY"),
		ActiveAAService:                  viper.GetString("ACTIVE_AA_SERVICE"),
		BundlerUrlEthereum:               viper.GetString("BUNDLER_URL_ETHEREUM"),
		PaymasterUrlEthereum:             viper.GetString("PAYMASTER_URL_ETHEREUM"),
		BundlerUrlPolygon:                viper.GetString("BUNDLER_URL_POLYGON"),
		PaymasterUrlPolygon:              viper.GetString("PAYMASTER_URL_POLYGON"),
		BundlerUrlBase:                   viper.GetString("BUNDLER_URL_BASE"),
		PaymasterUrlBase:                 viper.GetString("PAYMASTER_URL_BASE"),
		BundlerUrlBSC:                    viper.GetString("BUNDLER_URL_BSC"),
		PaymasterUrlBSC:                  viper.GetString("PAYMASTER_URL_BSC"),
		BundlerUrlArbitrum:               viper.GetString("BUNDLER_URL_ARBITRUM"),
		PaymasterUrlArbitrum:             viper.GetString("PAYMASTER_URL_ARBITRUM"),
		EntryPointContractAddress:        common.HexToAddress(viper.GetString("ENTRY_POINT_CONTRACT_ADDRESS")),
		BucketQueueRebuildInterval:       viper.GetInt("BUCKET_QUEUE_REBUILD_INTERVAL"),
		RefundCancellationCount:          viper.GetInt("REFUND_CANCELLATION_COUNT"),
		PercentDeviationFromExternalRate: decimal.NewFromFloat(viper.GetFloat64("PERCENT_DEVIATION_FROM_EXTERNAL_RATE")),
		PercentDeviationFromMarketRate:   decimal.NewFromFloat(viper.GetFloat64("PERCENT_DEVIATION_FROM_MARKET_RATE")),
	}
}

func init() {
	if err := SetupConfig(); err != nil {
		panic(fmt.Sprintf("config SetupConfig() error: %s", err))
	}
}
