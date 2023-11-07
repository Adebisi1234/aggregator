// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/apikey"
	"github.com/paycrest/protocol/ent/fiatcurrency"
	"github.com/paycrest/protocol/ent/lockorderfulfillment"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/network"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/providerordertoken"
	"github.com/paycrest/protocol/ent/providerprofile"
	"github.com/paycrest/protocol/ent/providerrating"
	"github.com/paycrest/protocol/ent/provisionbucket"
	"github.com/paycrest/protocol/ent/receiveaddress"
	"github.com/paycrest/protocol/ent/schema"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/paycrest/protocol/ent/token"
	"github.com/paycrest/protocol/ent/user"
	"github.com/paycrest/protocol/ent/verificationtoken"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apikeyFields := schema.APIKey{}.Fields()
	_ = apikeyFields
	// apikeyDescSecret is the schema descriptor for secret field.
	apikeyDescSecret := apikeyFields[1].Descriptor()
	// apikey.SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	apikey.SecretValidator = apikeyDescSecret.Validators[0].(func(string) error)
	// apikeyDescID is the schema descriptor for id field.
	apikeyDescID := apikeyFields[0].Descriptor()
	// apikey.DefaultID holds the default value on creation for the id field.
	apikey.DefaultID = apikeyDescID.Default.(func() uuid.UUID)
	fiatcurrencyMixin := schema.FiatCurrency{}.Mixin()
	fiatcurrencyMixinFields0 := fiatcurrencyMixin[0].Fields()
	_ = fiatcurrencyMixinFields0
	fiatcurrencyFields := schema.FiatCurrency{}.Fields()
	_ = fiatcurrencyFields
	// fiatcurrencyDescCreatedAt is the schema descriptor for created_at field.
	fiatcurrencyDescCreatedAt := fiatcurrencyMixinFields0[0].Descriptor()
	// fiatcurrency.DefaultCreatedAt holds the default value on creation for the created_at field.
	fiatcurrency.DefaultCreatedAt = fiatcurrencyDescCreatedAt.Default.(func() time.Time)
	// fiatcurrencyDescUpdatedAt is the schema descriptor for updated_at field.
	fiatcurrencyDescUpdatedAt := fiatcurrencyMixinFields0[1].Descriptor()
	// fiatcurrency.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fiatcurrency.DefaultUpdatedAt = fiatcurrencyDescUpdatedAt.Default.(func() time.Time)
	// fiatcurrency.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fiatcurrency.UpdateDefaultUpdatedAt = fiatcurrencyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fiatcurrencyDescDecimals is the schema descriptor for decimals field.
	fiatcurrencyDescDecimals := fiatcurrencyFields[3].Descriptor()
	// fiatcurrency.DefaultDecimals holds the default value on creation for the decimals field.
	fiatcurrency.DefaultDecimals = fiatcurrencyDescDecimals.Default.(int)
	// fiatcurrencyDescIsEnabled is the schema descriptor for is_enabled field.
	fiatcurrencyDescIsEnabled := fiatcurrencyFields[7].Descriptor()
	// fiatcurrency.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	fiatcurrency.DefaultIsEnabled = fiatcurrencyDescIsEnabled.Default.(bool)
	// fiatcurrencyDescID is the schema descriptor for id field.
	fiatcurrencyDescID := fiatcurrencyFields[0].Descriptor()
	// fiatcurrency.DefaultID holds the default value on creation for the id field.
	fiatcurrency.DefaultID = fiatcurrencyDescID.Default.(func() uuid.UUID)
	lockorderfulfillmentMixin := schema.LockOrderFulfillment{}.Mixin()
	lockorderfulfillmentMixinFields0 := lockorderfulfillmentMixin[0].Fields()
	_ = lockorderfulfillmentMixinFields0
	lockorderfulfillmentFields := schema.LockOrderFulfillment{}.Fields()
	_ = lockorderfulfillmentFields
	// lockorderfulfillmentDescCreatedAt is the schema descriptor for created_at field.
	lockorderfulfillmentDescCreatedAt := lockorderfulfillmentMixinFields0[0].Descriptor()
	// lockorderfulfillment.DefaultCreatedAt holds the default value on creation for the created_at field.
	lockorderfulfillment.DefaultCreatedAt = lockorderfulfillmentDescCreatedAt.Default.(func() time.Time)
	// lockorderfulfillmentDescUpdatedAt is the schema descriptor for updated_at field.
	lockorderfulfillmentDescUpdatedAt := lockorderfulfillmentMixinFields0[1].Descriptor()
	// lockorderfulfillment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	lockorderfulfillment.DefaultUpdatedAt = lockorderfulfillmentDescUpdatedAt.Default.(func() time.Time)
	// lockorderfulfillment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	lockorderfulfillment.UpdateDefaultUpdatedAt = lockorderfulfillmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// lockorderfulfillmentDescID is the schema descriptor for id field.
	lockorderfulfillmentDescID := lockorderfulfillmentFields[0].Descriptor()
	// lockorderfulfillment.DefaultID holds the default value on creation for the id field.
	lockorderfulfillment.DefaultID = lockorderfulfillmentDescID.Default.(func() uuid.UUID)
	lockpaymentorderMixin := schema.LockPaymentOrder{}.Mixin()
	lockpaymentorderMixinFields0 := lockpaymentorderMixin[0].Fields()
	_ = lockpaymentorderMixinFields0
	lockpaymentorderFields := schema.LockPaymentOrder{}.Fields()
	_ = lockpaymentorderFields
	// lockpaymentorderDescCreatedAt is the schema descriptor for created_at field.
	lockpaymentorderDescCreatedAt := lockpaymentorderMixinFields0[0].Descriptor()
	// lockpaymentorder.DefaultCreatedAt holds the default value on creation for the created_at field.
	lockpaymentorder.DefaultCreatedAt = lockpaymentorderDescCreatedAt.Default.(func() time.Time)
	// lockpaymentorderDescUpdatedAt is the schema descriptor for updated_at field.
	lockpaymentorderDescUpdatedAt := lockpaymentorderMixinFields0[1].Descriptor()
	// lockpaymentorder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	lockpaymentorder.DefaultUpdatedAt = lockpaymentorderDescUpdatedAt.Default.(func() time.Time)
	// lockpaymentorder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	lockpaymentorder.UpdateDefaultUpdatedAt = lockpaymentorderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// lockpaymentorderDescTxHash is the schema descriptor for tx_hash field.
	lockpaymentorderDescTxHash := lockpaymentorderFields[5].Descriptor()
	// lockpaymentorder.TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	lockpaymentorder.TxHashValidator = lockpaymentorderDescTxHash.Validators[0].(func(string) error)
	// lockpaymentorderDescCancellationCount is the schema descriptor for cancellation_count field.
	lockpaymentorderDescCancellationCount := lockpaymentorderFields[13].Descriptor()
	// lockpaymentorder.DefaultCancellationCount holds the default value on creation for the cancellation_count field.
	lockpaymentorder.DefaultCancellationCount = lockpaymentorderDescCancellationCount.Default.(int)
	// lockpaymentorderDescCancellationReasons is the schema descriptor for cancellation_reasons field.
	lockpaymentorderDescCancellationReasons := lockpaymentorderFields[14].Descriptor()
	// lockpaymentorder.DefaultCancellationReasons holds the default value on creation for the cancellation_reasons field.
	lockpaymentorder.DefaultCancellationReasons = lockpaymentorderDescCancellationReasons.Default.([]string)
	// lockpaymentorderDescID is the schema descriptor for id field.
	lockpaymentorderDescID := lockpaymentorderFields[0].Descriptor()
	// lockpaymentorder.DefaultID holds the default value on creation for the id field.
	lockpaymentorder.DefaultID = lockpaymentorderDescID.Default.(func() uuid.UUID)
	networkMixin := schema.Network{}.Mixin()
	networkMixinFields0 := networkMixin[0].Fields()
	_ = networkMixinFields0
	networkFields := schema.Network{}.Fields()
	_ = networkFields
	// networkDescCreatedAt is the schema descriptor for created_at field.
	networkDescCreatedAt := networkMixinFields0[0].Descriptor()
	// network.DefaultCreatedAt holds the default value on creation for the created_at field.
	network.DefaultCreatedAt = networkDescCreatedAt.Default.(func() time.Time)
	// networkDescUpdatedAt is the schema descriptor for updated_at field.
	networkDescUpdatedAt := networkMixinFields0[1].Descriptor()
	// network.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	network.DefaultUpdatedAt = networkDescUpdatedAt.Default.(func() time.Time)
	// network.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	network.UpdateDefaultUpdatedAt = networkDescUpdatedAt.UpdateDefault.(func() time.Time)
	paymentorderMixin := schema.PaymentOrder{}.Mixin()
	paymentorderMixinFields0 := paymentorderMixin[0].Fields()
	_ = paymentorderMixinFields0
	paymentorderFields := schema.PaymentOrder{}.Fields()
	_ = paymentorderFields
	// paymentorderDescCreatedAt is the schema descriptor for created_at field.
	paymentorderDescCreatedAt := paymentorderMixinFields0[0].Descriptor()
	// paymentorder.DefaultCreatedAt holds the default value on creation for the created_at field.
	paymentorder.DefaultCreatedAt = paymentorderDescCreatedAt.Default.(func() time.Time)
	// paymentorderDescUpdatedAt is the schema descriptor for updated_at field.
	paymentorderDescUpdatedAt := paymentorderMixinFields0[1].Descriptor()
	// paymentorder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	paymentorder.DefaultUpdatedAt = paymentorderDescUpdatedAt.Default.(func() time.Time)
	// paymentorder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	paymentorder.UpdateDefaultUpdatedAt = paymentorderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// paymentorderDescTxHash is the schema descriptor for tx_hash field.
	paymentorderDescTxHash := paymentorderFields[3].Descriptor()
	// paymentorder.TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	paymentorder.TxHashValidator = paymentorderDescTxHash.Validators[0].(func(string) error)
	// paymentorderDescReceiveAddressText is the schema descriptor for receive_address_text field.
	paymentorderDescReceiveAddressText := paymentorderFields[4].Descriptor()
	// paymentorder.ReceiveAddressTextValidator is a validator for the "receive_address_text" field. It is called by the builders before save.
	paymentorder.ReceiveAddressTextValidator = paymentorderDescReceiveAddressText.Validators[0].(func(string) error)
	// paymentorderDescID is the schema descriptor for id field.
	paymentorderDescID := paymentorderFields[0].Descriptor()
	// paymentorder.DefaultID holds the default value on creation for the id field.
	paymentorder.DefaultID = paymentorderDescID.Default.(func() uuid.UUID)
	providerordertokenMixin := schema.ProviderOrderToken{}.Mixin()
	providerordertokenMixinFields0 := providerordertokenMixin[0].Fields()
	_ = providerordertokenMixinFields0
	providerordertokenFields := schema.ProviderOrderToken{}.Fields()
	_ = providerordertokenFields
	// providerordertokenDescCreatedAt is the schema descriptor for created_at field.
	providerordertokenDescCreatedAt := providerordertokenMixinFields0[0].Descriptor()
	// providerordertoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	providerordertoken.DefaultCreatedAt = providerordertokenDescCreatedAt.Default.(func() time.Time)
	// providerordertokenDescUpdatedAt is the schema descriptor for updated_at field.
	providerordertokenDescUpdatedAt := providerordertokenMixinFields0[1].Descriptor()
	// providerordertoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	providerordertoken.DefaultUpdatedAt = providerordertokenDescUpdatedAt.Default.(func() time.Time)
	// providerordertoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	providerordertoken.UpdateDefaultUpdatedAt = providerordertokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	providerprofileFields := schema.ProviderProfile{}.Fields()
	_ = providerprofileFields
	// providerprofileDescTradingName is the schema descriptor for trading_name field.
	providerprofileDescTradingName := providerprofileFields[1].Descriptor()
	// providerprofile.TradingNameValidator is a validator for the "trading_name" field. It is called by the builders before save.
	providerprofile.TradingNameValidator = providerprofileDescTradingName.Validators[0].(func(string) error)
	// providerprofileDescIsPartner is the schema descriptor for is_partner field.
	providerprofileDescIsPartner := providerprofileFields[4].Descriptor()
	// providerprofile.DefaultIsPartner holds the default value on creation for the is_partner field.
	providerprofile.DefaultIsPartner = providerprofileDescIsPartner.Default.(bool)
	// providerprofileDescUpdatedAt is the schema descriptor for updated_at field.
	providerprofileDescUpdatedAt := providerprofileFields[5].Descriptor()
	// providerprofile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	providerprofile.DefaultUpdatedAt = providerprofileDescUpdatedAt.Default.(func() time.Time)
	// providerprofile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	providerprofile.UpdateDefaultUpdatedAt = providerprofileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providerprofileDescID is the schema descriptor for id field.
	providerprofileDescID := providerprofileFields[0].Descriptor()
	// providerprofile.DefaultID holds the default value on creation for the id field.
	providerprofile.DefaultID = providerprofileDescID.Default.(func() string)
	providerratingMixin := schema.ProviderRating{}.Mixin()
	providerratingMixinFields0 := providerratingMixin[0].Fields()
	_ = providerratingMixinFields0
	providerratingFields := schema.ProviderRating{}.Fields()
	_ = providerratingFields
	// providerratingDescCreatedAt is the schema descriptor for created_at field.
	providerratingDescCreatedAt := providerratingMixinFields0[0].Descriptor()
	// providerrating.DefaultCreatedAt holds the default value on creation for the created_at field.
	providerrating.DefaultCreatedAt = providerratingDescCreatedAt.Default.(func() time.Time)
	// providerratingDescUpdatedAt is the schema descriptor for updated_at field.
	providerratingDescUpdatedAt := providerratingMixinFields0[1].Descriptor()
	// providerrating.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	providerrating.DefaultUpdatedAt = providerratingDescUpdatedAt.Default.(func() time.Time)
	// providerrating.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	providerrating.UpdateDefaultUpdatedAt = providerratingDescUpdatedAt.UpdateDefault.(func() time.Time)
	provisionbucketFields := schema.ProvisionBucket{}.Fields()
	_ = provisionbucketFields
	// provisionbucketDescCreatedAt is the schema descriptor for created_at field.
	provisionbucketDescCreatedAt := provisionbucketFields[2].Descriptor()
	// provisionbucket.DefaultCreatedAt holds the default value on creation for the created_at field.
	provisionbucket.DefaultCreatedAt = provisionbucketDescCreatedAt.Default.(func() time.Time)
	receiveaddressMixin := schema.ReceiveAddress{}.Mixin()
	receiveaddressMixinFields0 := receiveaddressMixin[0].Fields()
	_ = receiveaddressMixinFields0
	receiveaddressFields := schema.ReceiveAddress{}.Fields()
	_ = receiveaddressFields
	// receiveaddressDescCreatedAt is the schema descriptor for created_at field.
	receiveaddressDescCreatedAt := receiveaddressMixinFields0[0].Descriptor()
	// receiveaddress.DefaultCreatedAt holds the default value on creation for the created_at field.
	receiveaddress.DefaultCreatedAt = receiveaddressDescCreatedAt.Default.(func() time.Time)
	// receiveaddressDescUpdatedAt is the schema descriptor for updated_at field.
	receiveaddressDescUpdatedAt := receiveaddressMixinFields0[1].Descriptor()
	// receiveaddress.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	receiveaddress.DefaultUpdatedAt = receiveaddressDescUpdatedAt.Default.(func() time.Time)
	// receiveaddress.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	receiveaddress.UpdateDefaultUpdatedAt = receiveaddressDescUpdatedAt.UpdateDefault.(func() time.Time)
	senderprofileFields := schema.SenderProfile{}.Fields()
	_ = senderprofileFields
	// senderprofileDescDomainWhitelist is the schema descriptor for domain_whitelist field.
	senderprofileDescDomainWhitelist := senderprofileFields[5].Descriptor()
	// senderprofile.DefaultDomainWhitelist holds the default value on creation for the domain_whitelist field.
	senderprofile.DefaultDomainWhitelist = senderprofileDescDomainWhitelist.Default.([]string)
	// senderprofileDescUpdatedAt is the schema descriptor for updated_at field.
	senderprofileDescUpdatedAt := senderprofileFields[6].Descriptor()
	// senderprofile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	senderprofile.DefaultUpdatedAt = senderprofileDescUpdatedAt.Default.(func() time.Time)
	// senderprofile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	senderprofile.UpdateDefaultUpdatedAt = senderprofileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// senderprofileDescID is the schema descriptor for id field.
	senderprofileDescID := senderprofileFields[0].Descriptor()
	// senderprofile.DefaultID holds the default value on creation for the id field.
	senderprofile.DefaultID = senderprofileDescID.Default.(func() uuid.UUID)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenMixinFields0[0].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenMixinFields0[1].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescSymbol is the schema descriptor for symbol field.
	tokenDescSymbol := tokenFields[0].Descriptor()
	// token.SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	token.SymbolValidator = tokenDescSymbol.Validators[0].(func(string) error)
	// tokenDescContractAddress is the schema descriptor for contract_address field.
	tokenDescContractAddress := tokenFields[1].Descriptor()
	// token.ContractAddressValidator is a validator for the "contract_address" field. It is called by the builders before save.
	token.ContractAddressValidator = tokenDescContractAddress.Validators[0].(func(string) error)
	// tokenDescIsEnabled is the schema descriptor for is_enabled field.
	tokenDescIsEnabled := tokenFields[3].Descriptor()
	// token.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	token.DefaultIsEnabled = tokenDescIsEnabled.Default.(bool)
	userMixin := schema.User{}.Mixin()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[1].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[2].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescIsVerified is the schema descriptor for is_verified field.
	userDescIsVerified := userFields[6].Descriptor()
	// user.DefaultIsVerified holds the default value on creation for the is_verified field.
	user.DefaultIsVerified = userDescIsVerified.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	verificationtokenMixin := schema.VerificationToken{}.Mixin()
	verificationtokenHooks := schema.VerificationToken{}.Hooks()
	verificationtoken.Hooks[0] = verificationtokenHooks[0]
	verificationtokenMixinFields0 := verificationtokenMixin[0].Fields()
	_ = verificationtokenMixinFields0
	verificationtokenFields := schema.VerificationToken{}.Fields()
	_ = verificationtokenFields
	// verificationtokenDescCreatedAt is the schema descriptor for created_at field.
	verificationtokenDescCreatedAt := verificationtokenMixinFields0[0].Descriptor()
	// verificationtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	verificationtoken.DefaultCreatedAt = verificationtokenDescCreatedAt.Default.(func() time.Time)
	// verificationtokenDescUpdatedAt is the schema descriptor for updated_at field.
	verificationtokenDescUpdatedAt := verificationtokenMixinFields0[1].Descriptor()
	// verificationtoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	verificationtoken.DefaultUpdatedAt = verificationtokenDescUpdatedAt.Default.(func() time.Time)
	// verificationtoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	verificationtoken.UpdateDefaultUpdatedAt = verificationtokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// verificationtokenDescID is the schema descriptor for id field.
	verificationtokenDescID := verificationtokenFields[0].Descriptor()
	// verificationtoken.DefaultID holds the default value on creation for the id field.
	verificationtoken.DefaultID = verificationtokenDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
