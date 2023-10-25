// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APIKeysColumns holds the columns for the "api_keys" table.
	APIKeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "secret", Type: field.TypeString, Unique: true},
		{Name: "provider_profile_api_key", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "sender_profile_api_key", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "validator_profile_api_key", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// APIKeysTable holds the schema information for the "api_keys" table.
	APIKeysTable = &schema.Table{
		Name:       "api_keys",
		Columns:    APIKeysColumns,
		PrimaryKey: []*schema.Column{APIKeysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_keys_provider_profiles_api_key",
				Columns:    []*schema.Column{APIKeysColumns[2]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "api_keys_sender_profiles_api_key",
				Columns:    []*schema.Column{APIKeysColumns[3]},
				RefColumns: []*schema.Column{SenderProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "api_keys_validator_profiles_api_key",
				Columns:    []*schema.Column{APIKeysColumns[4]},
				RefColumns: []*schema.Column{ValidatorProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FiatCurrenciesColumns holds the columns for the "fiat_currencies" table.
	FiatCurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "short_name", Type: field.TypeString, Unique: true},
		{Name: "decimals", Type: field.TypeInt, Default: 2},
		{Name: "symbol", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "market_rate", Type: field.TypeFloat64},
	}
	// FiatCurrenciesTable holds the schema information for the "fiat_currencies" table.
	FiatCurrenciesTable = &schema.Table{
		Name:       "fiat_currencies",
		Columns:    FiatCurrenciesColumns,
		PrimaryKey: []*schema.Column{FiatCurrenciesColumns[0]},
	}
	// LockOrderFulfillmentsColumns holds the columns for the "lock_order_fulfillments" table.
	LockOrderFulfillmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "tx_id", Type: field.TypeString},
		{Name: "tx_receipt_image", Type: field.TypeString},
		{Name: "confirmations", Type: field.TypeInt, Default: 0},
		{Name: "validation_errors", Type: field.TypeJSON},
		{Name: "lock_payment_order_fulfillment", Type: field.TypeUUID, Unique: true},
	}
	// LockOrderFulfillmentsTable holds the schema information for the "lock_order_fulfillments" table.
	LockOrderFulfillmentsTable = &schema.Table{
		Name:       "lock_order_fulfillments",
		Columns:    LockOrderFulfillmentsColumns,
		PrimaryKey: []*schema.Column{LockOrderFulfillmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lock_order_fulfillments_lock_payment_orders_fulfillment",
				Columns:    []*schema.Column{LockOrderFulfillmentsColumns[7]},
				RefColumns: []*schema.Column{LockPaymentOrdersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LockPaymentOrdersColumns holds the columns for the "lock_payment_orders" table.
	LockPaymentOrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "order_id", Type: field.TypeString},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "rate", Type: field.TypeFloat64},
		{Name: "order_percent", Type: field.TypeFloat64, Nullable: true},
		{Name: "tx_hash", Type: field.TypeString, Nullable: true, Size: 70},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "processing", "cancelled", "fulfilled", "validated", "settled"}, Default: "pending"},
		{Name: "block_number", Type: field.TypeInt64},
		{Name: "institution", Type: field.TypeString},
		{Name: "account_identifier", Type: field.TypeString},
		{Name: "account_name", Type: field.TypeString},
		{Name: "cancellation_count", Type: field.TypeInt, Default: 0},
		{Name: "cancellation_reasons", Type: field.TypeJSON},
		{Name: "provider_profile_assigned_orders", Type: field.TypeString, Nullable: true},
		{Name: "provision_bucket_lock_payment_orders", Type: field.TypeInt, Nullable: true},
		{Name: "token_lock_payment_orders", Type: field.TypeInt},
	}
	// LockPaymentOrdersTable holds the schema information for the "lock_payment_orders" table.
	LockPaymentOrdersTable = &schema.Table{
		Name:       "lock_payment_orders",
		Columns:    LockPaymentOrdersColumns,
		PrimaryKey: []*schema.Column{LockPaymentOrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lock_payment_orders_provider_profiles_assigned_orders",
				Columns:    []*schema.Column{LockPaymentOrdersColumns[15]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "lock_payment_orders_provision_buckets_lock_payment_orders",
				Columns:    []*schema.Column{LockPaymentOrdersColumns[16]},
				RefColumns: []*schema.Column{ProvisionBucketsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "lock_payment_orders_tokens_lock_payment_orders",
				Columns:    []*schema.Column{LockPaymentOrdersColumns[17]},
				RefColumns: []*schema.Column{TokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NetworksColumns holds the columns for the "networks" table.
	NetworksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chain_id", Type: field.TypeInt64},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "rpc_endpoint", Type: field.TypeString},
		{Name: "is_testnet", Type: field.TypeBool},
	}
	// NetworksTable holds the schema information for the "networks" table.
	NetworksTable = &schema.Table{
		Name:       "networks",
		Columns:    NetworksColumns,
		PrimaryKey: []*schema.Column{NetworksColumns[0]},
	}
	// PaymentOrdersColumns holds the columns for the "payment_orders" table.
	PaymentOrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "amount_paid", Type: field.TypeFloat64},
		{Name: "tx_hash", Type: field.TypeString, Nullable: true, Size: 70},
		{Name: "receive_address_text", Type: field.TypeString, Size: 60},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"initiated", "pending", "settled", "cancelled", "failed", "refunded"}, Default: "initiated"},
		{Name: "last_used", Type: field.TypeTime, Nullable: true},
		{Name: "api_key_payment_orders", Type: field.TypeUUID, Nullable: true},
		{Name: "sender_profile_payment_orders", Type: field.TypeUUID},
		{Name: "token_payment_orders", Type: field.TypeInt},
	}
	// PaymentOrdersTable holds the schema information for the "payment_orders" table.
	PaymentOrdersTable = &schema.Table{
		Name:       "payment_orders",
		Columns:    PaymentOrdersColumns,
		PrimaryKey: []*schema.Column{PaymentOrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_orders_api_keys_payment_orders",
				Columns:    []*schema.Column{PaymentOrdersColumns[9]},
				RefColumns: []*schema.Column{APIKeysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "payment_orders_sender_profiles_payment_orders",
				Columns:    []*schema.Column{PaymentOrdersColumns[10]},
				RefColumns: []*schema.Column{SenderProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "payment_orders_tokens_payment_orders",
				Columns:    []*schema.Column{PaymentOrdersColumns[11]},
				RefColumns: []*schema.Column{TokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PaymentOrderRecipientsColumns holds the columns for the "payment_order_recipients" table.
	PaymentOrderRecipientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "institution", Type: field.TypeString},
		{Name: "account_identifier", Type: field.TypeString},
		{Name: "account_name", Type: field.TypeString},
		{Name: "provider_id", Type: field.TypeString, Nullable: true},
		{Name: "payment_order_recipient", Type: field.TypeUUID, Unique: true},
	}
	// PaymentOrderRecipientsTable holds the schema information for the "payment_order_recipients" table.
	PaymentOrderRecipientsTable = &schema.Table{
		Name:       "payment_order_recipients",
		Columns:    PaymentOrderRecipientsColumns,
		PrimaryKey: []*schema.Column{PaymentOrderRecipientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_order_recipients_payment_orders_recipient",
				Columns:    []*schema.Column{PaymentOrderRecipientsColumns[5]},
				RefColumns: []*schema.Column{PaymentOrdersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProviderAvailabilitiesColumns holds the columns for the "provider_availabilities" table.
	ProviderAvailabilitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cadence", Type: field.TypeEnum, Enums: []string{"always", "weekdays", "weekends"}},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "provider_profile_availability", Type: field.TypeString, Unique: true},
	}
	// ProviderAvailabilitiesTable holds the schema information for the "provider_availabilities" table.
	ProviderAvailabilitiesTable = &schema.Table{
		Name:       "provider_availabilities",
		Columns:    ProviderAvailabilitiesColumns,
		PrimaryKey: []*schema.Column{ProviderAvailabilitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provider_availabilities_provider_profiles_availability",
				Columns:    []*schema.Column{ProviderAvailabilitiesColumns[4]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProviderOrderTokensColumns holds the columns for the "provider_order_tokens" table.
	ProviderOrderTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "symbol", Type: field.TypeString, Unique: true},
		{Name: "fixed_conversion_rate", Type: field.TypeFloat64},
		{Name: "floating_conversion_rate", Type: field.TypeFloat64},
		{Name: "conversion_rate_type", Type: field.TypeEnum, Enums: []string{"fixed", "floating"}},
		{Name: "max_order_amount", Type: field.TypeFloat64},
		{Name: "min_order_amount", Type: field.TypeFloat64},
		{Name: "addresses", Type: field.TypeJSON},
		{Name: "provider_profile_order_tokens", Type: field.TypeString, Nullable: true},
	}
	// ProviderOrderTokensTable holds the schema information for the "provider_order_tokens" table.
	ProviderOrderTokensTable = &schema.Table{
		Name:       "provider_order_tokens",
		Columns:    ProviderOrderTokensColumns,
		PrimaryKey: []*schema.Column{ProviderOrderTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provider_order_tokens_provider_profiles_order_tokens",
				Columns:    []*schema.Column{ProviderOrderTokensColumns[10]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProviderProfilesColumns holds the columns for the "provider_profiles" table.
	ProviderProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "trading_name", Type: field.TypeString, Size: 80},
		{Name: "host_identifier", Type: field.TypeString, Nullable: true},
		{Name: "provision_mode", Type: field.TypeEnum, Enums: []string{"manual", "auto"}, Default: "auto"},
		{Name: "is_partner", Type: field.TypeBool, Default: false},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "fiat_currency_provider", Type: field.TypeUUID, Unique: true},
		{Name: "user_provider_profile", Type: field.TypeUUID, Unique: true},
	}
	// ProviderProfilesTable holds the schema information for the "provider_profiles" table.
	ProviderProfilesTable = &schema.Table{
		Name:       "provider_profiles",
		Columns:    ProviderProfilesColumns,
		PrimaryKey: []*schema.Column{ProviderProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provider_profiles_fiat_currencies_provider",
				Columns:    []*schema.Column{ProviderProfilesColumns[6]},
				RefColumns: []*schema.Column{FiatCurrenciesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "provider_profiles_users_provider_profile",
				Columns:    []*schema.Column{ProviderProfilesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProviderRatingsColumns holds the columns for the "provider_ratings" table.
	ProviderRatingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "trust_score", Type: field.TypeFloat64},
		{Name: "provider_profile_provider_rating", Type: field.TypeString, Unique: true},
	}
	// ProviderRatingsTable holds the schema information for the "provider_ratings" table.
	ProviderRatingsTable = &schema.Table{
		Name:       "provider_ratings",
		Columns:    ProviderRatingsColumns,
		PrimaryKey: []*schema.Column{ProviderRatingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provider_ratings_provider_profiles_provider_rating",
				Columns:    []*schema.Column{ProviderRatingsColumns[4]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProvisionBucketsColumns holds the columns for the "provision_buckets" table.
	ProvisionBucketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "min_amount", Type: field.TypeFloat64},
		{Name: "max_amount", Type: field.TypeFloat64},
		{Name: "currency", Type: field.TypeString, Size: 3},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ProvisionBucketsTable holds the schema information for the "provision_buckets" table.
	ProvisionBucketsTable = &schema.Table{
		Name:       "provision_buckets",
		Columns:    ProvisionBucketsColumns,
		PrimaryKey: []*schema.Column{ProvisionBucketsColumns[0]},
	}
	// ReceiveAddressesColumns holds the columns for the "receive_addresses" table.
	ReceiveAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "salt", Type: field.TypeBytes, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unused", "partial", "used", "expired"}, Default: "unused"},
		{Name: "last_indexed_block", Type: field.TypeInt64, Nullable: true},
		{Name: "last_used", Type: field.TypeTime, Nullable: true},
		{Name: "valid_until", Type: field.TypeTime, Nullable: true},
		{Name: "payment_order_receive_address", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// ReceiveAddressesTable holds the schema information for the "receive_addresses" table.
	ReceiveAddressesTable = &schema.Table{
		Name:       "receive_addresses",
		Columns:    ReceiveAddressesColumns,
		PrimaryKey: []*schema.Column{ReceiveAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "receive_addresses_payment_orders_receive_address",
				Columns:    []*schema.Column{ReceiveAddressesColumns[9]},
				RefColumns: []*schema.Column{PaymentOrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SenderProfilesColumns holds the columns for the "sender_profiles" table.
	SenderProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "webhook_url", Type: field.TypeString, Nullable: true},
		{Name: "domain_whitelist", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_sender_profile", Type: field.TypeUUID, Unique: true},
	}
	// SenderProfilesTable holds the schema information for the "sender_profiles" table.
	SenderProfilesTable = &schema.Table{
		Name:       "sender_profiles",
		Columns:    SenderProfilesColumns,
		PrimaryKey: []*schema.Column{SenderProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sender_profiles_users_sender_profile",
				Columns:    []*schema.Column{SenderProfilesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "symbol", Type: field.TypeString, Size: 10},
		{Name: "contract_address", Type: field.TypeString, Size: 60},
		{Name: "decimals", Type: field.TypeInt8},
		{Name: "is_enabled", Type: field.TypeBool, Default: false},
		{Name: "network_tokens", Type: field.TypeInt},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_networks_tokens",
				Columns:    []*schema.Column{TokensColumns[7]},
				RefColumns: []*schema.Column{NetworksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "first_name", Type: field.TypeString, Size: 80},
		{Name: "last_name", Type: field.TypeString, Size: 80},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "scope", Type: field.TypeEnum, Enums: []string{"sender", "provider", "tx_validator"}},
		{Name: "is_verified", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email_scope",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[5], UsersColumns[7]},
			},
		},
	}
	// ValidatorProfilesColumns holds the columns for the "validator_profiles" table.
	ValidatorProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "wallet_address", Type: field.TypeString, Nullable: true},
		{Name: "host_identifier", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_validator_profile", Type: field.TypeUUID, Unique: true},
	}
	// ValidatorProfilesTable holds the schema information for the "validator_profiles" table.
	ValidatorProfilesTable = &schema.Table{
		Name:       "validator_profiles",
		Columns:    ValidatorProfilesColumns,
		PrimaryKey: []*schema.Column{ValidatorProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "validator_profiles_users_validator_profile",
				Columns:    []*schema.Column{ValidatorProfilesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VerificationTokensColumns holds the columns for the "verification_tokens" table.
	VerificationTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeString},
		{Name: "scope", Type: field.TypeEnum, Enums: []string{"verification"}},
		{Name: "user_verification_token", Type: field.TypeUUID},
	}
	// VerificationTokensTable holds the schema information for the "verification_tokens" table.
	VerificationTokensTable = &schema.Table{
		Name:       "verification_tokens",
		Columns:    VerificationTokensColumns,
		PrimaryKey: []*schema.Column{VerificationTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "verification_tokens_users_verification_token",
				Columns:    []*schema.Column{VerificationTokensColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LockOrderFulfillmentValidatorsColumns holds the columns for the "lock_order_fulfillment_validators" table.
	LockOrderFulfillmentValidatorsColumns = []*schema.Column{
		{Name: "lock_order_fulfillment_id", Type: field.TypeUUID},
		{Name: "validator_profile_id", Type: field.TypeUUID},
	}
	// LockOrderFulfillmentValidatorsTable holds the schema information for the "lock_order_fulfillment_validators" table.
	LockOrderFulfillmentValidatorsTable = &schema.Table{
		Name:       "lock_order_fulfillment_validators",
		Columns:    LockOrderFulfillmentValidatorsColumns,
		PrimaryKey: []*schema.Column{LockOrderFulfillmentValidatorsColumns[0], LockOrderFulfillmentValidatorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lock_order_fulfillment_validators_lock_order_fulfillment_id",
				Columns:    []*schema.Column{LockOrderFulfillmentValidatorsColumns[0]},
				RefColumns: []*schema.Column{LockOrderFulfillmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "lock_order_fulfillment_validators_validator_profile_id",
				Columns:    []*schema.Column{LockOrderFulfillmentValidatorsColumns[1]},
				RefColumns: []*schema.Column{ValidatorProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProvisionBucketProviderProfilesColumns holds the columns for the "provision_bucket_provider_profiles" table.
	ProvisionBucketProviderProfilesColumns = []*schema.Column{
		{Name: "provision_bucket_id", Type: field.TypeInt},
		{Name: "provider_profile_id", Type: field.TypeString},
	}
	// ProvisionBucketProviderProfilesTable holds the schema information for the "provision_bucket_provider_profiles" table.
	ProvisionBucketProviderProfilesTable = &schema.Table{
		Name:       "provision_bucket_provider_profiles",
		Columns:    ProvisionBucketProviderProfilesColumns,
		PrimaryKey: []*schema.Column{ProvisionBucketProviderProfilesColumns[0], ProvisionBucketProviderProfilesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provision_bucket_provider_profiles_provision_bucket_id",
				Columns:    []*schema.Column{ProvisionBucketProviderProfilesColumns[0]},
				RefColumns: []*schema.Column{ProvisionBucketsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "provision_bucket_provider_profiles_provider_profile_id",
				Columns:    []*schema.Column{ProvisionBucketProviderProfilesColumns[1]},
				RefColumns: []*schema.Column{ProviderProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APIKeysTable,
		FiatCurrenciesTable,
		LockOrderFulfillmentsTable,
		LockPaymentOrdersTable,
		NetworksTable,
		PaymentOrdersTable,
		PaymentOrderRecipientsTable,
		ProviderAvailabilitiesTable,
		ProviderOrderTokensTable,
		ProviderProfilesTable,
		ProviderRatingsTable,
		ProvisionBucketsTable,
		ReceiveAddressesTable,
		SenderProfilesTable,
		TokensTable,
		UsersTable,
		ValidatorProfilesTable,
		VerificationTokensTable,
		LockOrderFulfillmentValidatorsTable,
		ProvisionBucketProviderProfilesTable,
	}
)

func init() {
	APIKeysTable.ForeignKeys[0].RefTable = ProviderProfilesTable
	APIKeysTable.ForeignKeys[1].RefTable = SenderProfilesTable
	APIKeysTable.ForeignKeys[2].RefTable = ValidatorProfilesTable
	LockOrderFulfillmentsTable.ForeignKeys[0].RefTable = LockPaymentOrdersTable
	LockPaymentOrdersTable.ForeignKeys[0].RefTable = ProviderProfilesTable
	LockPaymentOrdersTable.ForeignKeys[1].RefTable = ProvisionBucketsTable
	LockPaymentOrdersTable.ForeignKeys[2].RefTable = TokensTable
	PaymentOrdersTable.ForeignKeys[0].RefTable = APIKeysTable
	PaymentOrdersTable.ForeignKeys[1].RefTable = SenderProfilesTable
	PaymentOrdersTable.ForeignKeys[2].RefTable = TokensTable
	PaymentOrderRecipientsTable.ForeignKeys[0].RefTable = PaymentOrdersTable
	ProviderAvailabilitiesTable.ForeignKeys[0].RefTable = ProviderProfilesTable
	ProviderOrderTokensTable.ForeignKeys[0].RefTable = ProviderProfilesTable
	ProviderProfilesTable.ForeignKeys[0].RefTable = FiatCurrenciesTable
	ProviderProfilesTable.ForeignKeys[1].RefTable = UsersTable
	ProviderRatingsTable.ForeignKeys[0].RefTable = ProviderProfilesTable
	ReceiveAddressesTable.ForeignKeys[0].RefTable = PaymentOrdersTable
	SenderProfilesTable.ForeignKeys[0].RefTable = UsersTable
	TokensTable.ForeignKeys[0].RefTable = NetworksTable
	ValidatorProfilesTable.ForeignKeys[0].RefTable = UsersTable
	VerificationTokensTable.ForeignKeys[0].RefTable = UsersTable
	LockOrderFulfillmentValidatorsTable.ForeignKeys[0].RefTable = LockOrderFulfillmentsTable
	LockOrderFulfillmentValidatorsTable.ForeignKeys[1].RefTable = ValidatorProfilesTable
	ProvisionBucketProviderProfilesTable.ForeignKeys[0].RefTable = ProvisionBucketsTable
	ProvisionBucketProviderProfilesTable.ForeignKeys[1].RefTable = ProviderProfilesTable
}
