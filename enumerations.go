package enablebankinggo

// BalanceType represents the type of balance.
type BalanceType string

const (
	// ClosingAvailableBalanceType (ISO20022 Closing Available) Closing available balance.
	ClosingAvailableBalanceType BalanceType = "CLAV"

	// ClosingBookedBalanceType (ISO20022 Closing Booked) Closing booked balance.
	ClosingBookedBalanceType BalanceType = "CLBD"

	// ForwardBookedBalanceType (ISO20022 ForwardAvailable) Balance that is at the disposal of account holders on the date specified.
	ForwardAvailableBalanceType BalanceType = "FWAV"

	// InformationBalanceType (ISO20022 Information) Balance for informational purposes.
	InformationBalanceType BalanceType = "INFO"

	// InterimAvailableBalanceType (ISO20022 Interim Available) Available balance calculated in the course of the day.
	InterimAvailableBalanceType BalanceType = "ITAV"

	// InterimBookedBalanceType (ISO20022 Interim Booked) Booked balance calculated in the course of the day.
	InterimBookedBalanceType BalanceType = "ITBD"

	// OpeningAvailableBalanceType (ISO20022 Opening Available) Opening balance that is at the disposal of account holders at the beginning of the date specified.
	OpeningAvailableBalanceType BalanceType = "OPAV"

	// OpeningBookedBalanceType (ISO20022 Opening Booked) Book balance of the account at the beginning of the account reporting period. It always equals the closing book balance from the previous report.
	OpeningBookedBalanceType BalanceType = "OPBD"

	// OtherBalanceType Other Balance.
	OtherBalanceType BalanceType = "OTHR"

	// PreviouslyClosedBookedBalanceType (ISO20022 Previously Closed Booked) Balance of the account at the end of the previous reporting period.
	PreviouslyClosedBookedBalanceType BalanceType = "PRCD"

	// ValueDateBalanceType Value-date balance.
	ValueDateBalanceType BalanceType = "VALU"

	// ExpectedBalanceType (ISO20022 Expected) Instant Balance.
	ExpectedBalanceType BalanceType = "XPCD"
)

// IsEmpty checks if the BalanceType is empty.
func (bt BalanceType) IsEmpty() bool {
	return bt == ""
}

// IsValid checks if the BalanceType is valid.
func (bt BalanceType) IsValid() bool {
	_, ok := balanceTypeDescriptions[bt]
	return ok
}

// Description returns the description of the BalanceType.
func (bt BalanceType) Description() string {
	if desc, ok := balanceTypeDescriptions[bt]; ok {
		return desc
	}

	return ""
}

var balanceTypeDescriptions = map[BalanceType]string{
	ClosingAvailableBalanceType:       "Closing available balance",
	ClosingBookedBalanceType:          "Closing booked balance",
	ForwardAvailableBalanceType:       "Forward available balance",
	InformationBalanceType:            "Information balance",
	InterimAvailableBalanceType:       "Interim available balance",
	InterimBookedBalanceType:          "Interim booked balance",
	OpeningAvailableBalanceType:       "Opening available balance",
	OpeningBookedBalanceType:          "Opening booked balance",
	OtherBalanceType:                  "Other balance",
	PreviouslyClosedBookedBalanceType: "Previously closed booked balance",
	ValueDateBalanceType:              "Value-date balance",
	ExpectedBalanceType:               "Expected (instant) balance",
}

// BalanceTypeDescriptions returns a map of BalanceType to their descriptions.
func BalanceTypeDescriptions() map[BalanceType]string {
	return balanceTypeDescriptions
}

// CreditDebitIndicator represents whether the amount is a credit or a debit.
type CreditDebitIndicator string

const (
	// CreditCreditDebitIndicator Credit type transaction.
	CreditCreditDebitIndicator CreditDebitIndicator = "CRDT"

	// DebitCreditDebitIndicator Debit type transaction.
	DebitCreditDebitIndicator CreditDebitIndicator = "DBIT"
)

var creditDebitIndicatorDescriptions = map[CreditDebitIndicator]string{
	CreditCreditDebitIndicator: "Credit",
	DebitCreditDebitIndicator:  "Debit",
}

// IsEmpty checks if the CreditDebitIndicator is empty.
func (cdi CreditDebitIndicator) IsEmpty() bool {
	return cdi == ""
}

// IsValid checks if the CreditDebitIndicator is valid.
func (cdi CreditDebitIndicator) IsValid() bool {
	_, ok := creditDebitIndicatorDescriptions[cdi]
	return ok
}

// Description returns the description of the CreditDebitIndicator.
func (cdi CreditDebitIndicator) Description() string {
	if desc, ok := creditDebitIndicatorDescriptions[cdi]; ok {
		return desc
	}

	return ""
}

// CreditDebitIndicatorDescriptions returns a map of CreditDebitIndicator to their descriptions.
func CreditDebitIndicatorDescriptions() map[CreditDebitIndicator]string {
	return creditDebitIndicatorDescriptions
}

// PSUType represents type supported by ASPSP.
type PSUType string

const (
	// BusinessPSUType represents business PSU type.
	BusinessPSUType PSUType = "business"
	// PersonalPSUType represents personal PSU type.
	PersonalPSUType PSUType = "personal"
)

var psuTypeDescriptions = map[PSUType]string{
	BusinessPSUType: "Business",
	PersonalPSUType: "Personal",
}

// IsEmpty checks if the PSUType is empty.
func (pt PSUType) IsEmpty() bool {
	return pt == ""
}

// IsValid checks if the PSUType is valid.
func (pt PSUType) IsValid() bool {
	_, ok := psuTypeDescriptions[pt]
	return ok
}

// Description returns the description of the PSUType.
func (pt PSUType) Description() string {
	if desc, ok := psuTypeDescriptions[pt]; ok {
		return desc
	}

	return ""
}

// PSUTypeKeys returns a slice of PSUType as strings.
func PSUTypeKeys() []string {
	keys := make([]string, 0, len(psuTypeDescriptions))
	for k := range psuTypeDescriptions {
		keys = append(keys, string(k))
	}
	return keys
}

// PSUTypeDescriptions returns a map of PSUType to their descriptions.
func PSUTypeDescriptions() map[PSUType]string {
	return psuTypeDescriptions
}

// RateType represents the type of exchange rate.
type RateType string

const (
	// AGRDRateType Exchange rate applied is the rate agreed between the parties.
	AGRDRateType RateType = "AGRD"

	// SALERateType Exchange rate applied is the market rate at the time of the sale.
	SALERateType RateType = "SALE"

	// SPOTRateType Exchange rate applied is the spot rate.
	SPOTRateType RateType = "SPOT"
)

var rateTypeDescriptions = map[RateType]string{
	AGRDRateType: "Agreed rate",
	SALERateType: "Sale rate",
	SPOTRateType: "Spot rate",
}

// IsEmpty checks if the RateType is empty.
func (rt RateType) IsEmpty() bool {
	return rt == ""
}

// IsValid checks if the RateType is valid.
func (rt RateType) IsValid() bool {
	_, ok := rateTypeDescriptions[rt]
	return ok
}

// Description returns the description of the RateType.
func (rt RateType) Description() string {
	if desc, ok := rateTypeDescriptions[rt]; ok {
		return desc
	}

	return ""
}

// RateTypeDescriptions returns a map of RateType to their descriptions.
func RateTypeDescriptions() map[RateType]string {
	return rateTypeDescriptions
}

// HeaderKey represents a header key.
type HeaderKey string

const (
	// PSUIPAddressHeaderKey is the header key for passing PSU IP address.
	PSUIPAddressHeaderKey HeaderKey = "Psu-Ip-Address"

	// PSUUserAgentHeaderKey is the header key for passing PSU browser User Agent.
	PSUUserAgentHeaderKey HeaderKey = "Psu-User-Agent"

	// PSURefererHeaderKey is the header key for passing PSU Referer.
	PSURefererHeaderKey HeaderKey = "Psu-Referer"

	// PSUAcceptHeaderKey is the header key for passing PSU accept header.
	PSUAcceptHeaderKey HeaderKey = "Psu-Accept"

	// PSUAcceptCharsetHeaderKey is the header key for passing PSU charset.
	PSUAcceptCharsetHeaderKey HeaderKey = "Psu-Accept-Charset"

	// PSUAcceptEncodingHeaderKey is the header key for passing PSU accept encoding.
	PSUAcceptEncodingHeaderKey HeaderKey = "Psu-Accept-Encoding"

	// PSUAcceptLanguageHeaderKey is the header key for passing PSU accept language.
	PSUAcceptLanguageHeaderKey HeaderKey = "Psu-Accept-language"

	// PSUGeoLocationHeaderKey is the header key for passing PSU geo location.
	PSUGeoLocationHeaderKey HeaderKey = "Psu-Geo-Location"
)

var headerKeyDescriptions = map[HeaderKey]string{
	PSUIPAddressHeaderKey:      "PSU IP Address",
	PSUUserAgentHeaderKey:      "PSU User Agent",
	PSURefererHeaderKey:        "PSU Referer",
	PSUAcceptHeaderKey:         "PSU Accept",
	PSUAcceptCharsetHeaderKey:  "PSU Accept Charset",
	PSUAcceptEncodingHeaderKey: "PSU Accept Encoding",
	PSUAcceptLanguageHeaderKey: "PSU Accept Language",
	PSUGeoLocationHeaderKey:    "PSU Geo Location",
}

// IsEmpty checks if the HeaderKey is empty.
func (hk HeaderKey) IsEmpty() bool {
	return hk == ""
}

// IsValid checks if the HeaderKey is valid.
func (hk HeaderKey) IsValid() bool {
	_, ok := headerKeyDescriptions[hk]
	return ok
}

// Description returns the description of the HeaderKey.
func (hk HeaderKey) Description() string {
	if desc, ok := headerKeyDescriptions[hk]; ok {
		return desc
	}

	return ""
}

// HeaderKeyDescriptions returns a map of HeaderKey to their descriptions.
func HeaderKeyDescriptions() map[HeaderKey]string {
	return headerKeyDescriptions
}

// AuthenticationApproach represents authentication approach supported by ASPSP
// authentication method.
type AuthenticationApproach string

const (
	// DecoupledAuthenticationApproach indicates that the TPP identifies the PSU and forwards the
	// identification to the ASPSP which processes the authentication through a decoupled device.
	DecoupledAuthenticationApproach AuthenticationApproach = "DECOUPLED"

	// EmbeddedAuthenticationApproach indicates that the TPP identifies the PSU and forwards the
	// identification to the ASPSP which starts the authentication. The TPP forwards one
	// authentication factor of the PSU (e.g. OTP or response to a challenge).
	EmbeddedAuthenticationApproach AuthenticationApproach = "EMBEDDED"

	// RedirectAuthenticationApproach indicates that the PSU is redirected by the TPP to the ASPSP
	// which processes identification and authentication.
	RedirectAuthenticationApproach AuthenticationApproach = "REDIRECT"
)

// Service represents services supported by ASPSP.
type Service string

const (
	// AccountInformationService indicates that the ASPSP supports account information service (AIS).
	AccountInformationService Service = "AIS"

	// PaymentInitiationService indicates that the ASPSP supports payment initiation service (PIS).
	PaymentInitiationService Service = "PIS"
)

var serviceDescriptions = map[Service]string{
	AccountInformationService: "Account Information Service",
	PaymentInitiationService:  "Payment Initiation Service",
}

// IsEmpty checks if the Service is empty.
func (s Service) IsEmpty() bool {
	return s == ""
}

// IsValid checks if the Service is valid.
func (s Service) IsValid() bool {
	_, ok := serviceDescriptions[s]
	return ok
}

// Description returns the description of the Service.
func (s Service) Description() string {
	if desc, ok := serviceDescriptions[s]; ok {
		return desc
	}

	return ""
}

// ServiceKeys returns a slice of Service as strings.
func ServiceKeys() []string {
	keys := make([]string, 0, len(serviceDescriptions))
	for k := range serviceDescriptions {
		keys = append(keys, string(k))
	}
	return keys
}

// ServiceDescriptions returns a map of Service to their descriptions.
func ServiceDescriptions() map[Service]string {
	return serviceDescriptions
}

// PaymentType represents payment types supported by ASPSP.
type PaymentType string

const (
	// BulkDomesticPaymentType indicates domestic bulk credit transfers.
	BulkDomesticPaymentType PaymentType = "BULK_DOMESTIC"

	// BulkSepaPaymentType indicates SEPA bulk credit transfers.
	BulkSepaPaymentType PaymentType = "BULK_SEPA"

	// CrossborderPaymentType indicates crossborder credit transfers.
	CrossborderPaymentType PaymentType = "CROSSBORDER"

	// DomesticPaymentType indicates domestic credit transfers.
	DomesticPaymentType PaymentType = "DOMESTIC"

	// DomesticSeGiroPaymentType indicates Swedish domestic Giro payments (BankGiro/PlusGiro).
	DomesticSeGiroPaymentType PaymentType = "DOMESTIC_SE_GIRO"

	// InstSepaPaymentType indicates instant SEPA credit transfers (without fallback to SEPA).
	InstSepaPaymentType PaymentType = "INST_SEPA"

	// InternalPaymentType indicates transfer made within an ASPSP.
	InternalPaymentType PaymentType = "INTERNAL"

	// SepaPaymentType indicates SEPA credit transfers.
	SepaPaymentType PaymentType = "SEPA"
)

// Environment represents application environment.
type Environment string

const (
	// ProductionEnvironment indicates production environment.
	ProductionEnvironment Environment = "PRODUCTION"

	// SandboxEnvironment indicates sandbox environment.
	SandboxEnvironment Environment = "SANDBOX"
)

// SchemeName represents identification scheme name.
type SchemeName string

const (
	// AlienRegistrationNumberScheme indicates AlienRegistrationNumber identification scheme.
	AlienRegistrationNumberScheme SchemeName = "ARNU"

	// BankPartyIdentificationScheme indicates BankPartyIdentification. Unique and unambiguous assignment
	// made by a specific bank or similar financial institution to identify a relationship as defined between
	// the bank and its client.
	BankPartyIdentificationScheme SchemeName = "BANK"

	// BasicBankAccountNumberScheme indicates Basic Bank Account Number. Represents a country-specific bank account number.
	BasicBankAccountNumberScheme SchemeName = "BBAN"

	// SwedishBankgiroNumberScheme indicates Swedish BankGiro account number. Used in domestic Swedish giro payments.
	SwedishBankgiroNumberScheme SchemeName = "BGNR"

	// PassportNumberScheme indicates passport number.
	PassportNumberScheme SchemeName = "CCPT"

	// ClearingIdentificationScheme indicates Clearing Identification Number.
	ClearingIdentificationScheme SchemeName = "CHID"

	// CountryIdentificationCodeScheme indicates Country Identification Code. Country authority given organisation
	// identification (e.g., corporate registration number).
	CountryIdentificationCodeScheme SchemeName = "COID"

	// CardPanScheme indicates Card PAN (masked or plain).
	CardPanScheme SchemeName = "CPAN"

	// CustomerIdentificationNumberIndividualScheme indicates Customer Identification Number for individuals.
	CustomerIdentificationNumberIndividualScheme SchemeName = "CUSI"

	// CorporateCustomerNumberScheme indicates Corporate Customer Number.
	CorporateCustomerNumberScheme SchemeName = "CUST"

	// DriversLicenseNumberScheme indicates Driver's License Number.
	DriversLicenseNumberScheme SchemeName = "DRLC"

	// DataUniversalNumberingSystemScheme indicates Data Universal Numbering System.
	DataUniversalNumberingSystemScheme SchemeName = "DUNS"

	// EmployerIdentificationNumberScheme indicates Employer Identification Number.
	EmployerIdentificationNumberScheme SchemeName = "EMPL"

	// GS1GLNIdentifierScheme indicates GS1 GLN Identifier.
	GS1GLNIdentifierScheme SchemeName = "GS1G"

	// InternationalBankAccountNumberScheme indicates International Bank Account Number (IBAN)  - identification
	// used internationally by financial institutions to uniquely identify the account of a customer.
	InternationalBankAccountNumberScheme SchemeName = "IBAN"

	// MaskedIBANScheme indicates Masked IBAN.
	MaskedIBANScheme SchemeName = "MIBN"

	// NationalIdentityNumberScheme indicates National Identity Number. Number assigned by an authority to
	// identify the national identity number of a person.
	NationalIdentityNumberScheme SchemeName = "NIDN"

	// OAUTH2AccessTokenScheme indicates OAUTH2 access token that is owned by the PISP being also an AISP and
	// that can be used in order to identify the PSU.
	OAUTH2AccessTokenScheme SchemeName = "OAUT"

	// OtherCorporateScheme indicates Other Corporate. Handelsbanken-specific code.
	OtherCorporateScheme SchemeName = "OTHC"

	// OtherIndividualScheme indicates Other Individual. Handelsbanken-specific code.
	OtherIndividualScheme SchemeName = "OTHI"

	// SwedishPlusGiroAccountNumberScheme indicates Swedish PlusGiro account number. Used in domestic Swedish
	// giro payments.
	SwedishPlusGiroAccountNumberScheme SchemeName = "PGNR"

	// SocialSecurityNumberScheme indicates Social Security Number.
	SocialSecurityNumberScheme SchemeName = "SOSE"

	// SIRENNumberScheme indicates The SIREN number is a 9 digit code assigned by INSEE, the French National
	// Institute for Statistics and Economic Studies, to identify an organisation in France.
	SIRENNumberScheme SchemeName = "SREN"

	// SIRETNumberScheme indicates The SIRET number is a 14 digit code assigned by INSEE, the French National
	// Institute for Statistics and Economic Studies, to identify an organisation unit in France. It consists
	// of the SIREN number, followed by a five digit classification number, to identify the local geographical
	// unit of that entity.
	SIRETNumberScheme SchemeName = "SRET"

	// TaxIdentificationNumberScheme indicates Tax Identification Number.
	TaxIdentificationNumberScheme SchemeName = "TXID"
)

// Usage represents account usage type.
type Usage string

const (
	// ProfessionalAccountUsage professional account usage.
	ProfessionalAccountUsage Usage = "ORGA"

	// PrivateAccountUsage private account usage.
	PrivateAccountUsage Usage = "PRIV"
)

// CashAccountType represents the type of account.
type CashAccountType string

const (
	// CurrentCashAccountType account used to post debits and credits when no specific account has been nominated.
	CurrentCashAccountType CashAccountType = "CACC"

	// CardPaymentCashAccountType account used for card payments only.
	CardPaymentCashAccountType CashAccountType = "CARD"

	// CashPaymentCashAccountType account used for the payment of cash.
	CashPaymentCashAccountType CashAccountType = "CASH"

	// LoanCashAccountType account used for loans.
	LoanCashAccountType CashAccountType = "LOAN"

	// OtherCashAccountType account not otherwise specified.
	OtherCashAccountType CashAccountType = "OTHR"

	// SavingsCashAccountType account used for savings.
	SavingsCashAccountType CashAccountType = "SVGS"
)

// AddressType represents available address types.
type AddressType string

const (
	// BusinessAddressType business address.
	BusinessAddressType AddressType = "Business"

	// CorrespondenceAddressType correspondence address.
	CorrespondenceAddressType AddressType = "Correspondence"

	// DeliveryToAddressType delivery address.
	DeliveryToAddressType AddressType = "DeliveryTo"

	// MailToAddressType mail to address.
	MailToAddressType AddressType = "MailTo"

	// POBoxAddressType PO Box address.
	POBoxAddressType AddressType = "POBox"

	// PostalAddressType postal address.
	PostalAddressType AddressType = "Postal"

	// ResidentialAddressType residential address.
	ResidentialAddressType AddressType = "Residential"

	// StatementAddressType statement address.
	StatementAddressType AddressType = "Statement"
)

// ReferenceNumberScheme represents reference number schemes.
type ReferenceNumberScheme string

const (
	// BelgianReferenceNumberScheme indicates Belgian reference number.
	BelgianReferenceNumberScheme ReferenceNumberScheme = "BERF"

	// FinnishReferenceNumberScheme indicates Finnish reference number.
	FinnishReferenceNumberScheme ReferenceNumberScheme = "FIRF"

	// InternationalReferenceNumberScheme indicates International reference number (starting with RF).
	InternationalReferenceNumberScheme ReferenceNumberScheme = "INTL"

	// NorwegianKIDScheme indicates Norwegian KID (OCR).
	NorwegianKIDScheme ReferenceNumberScheme = "NORF"

	// SEPADirectDebitMandateIDScheme indicates SEPA Direct Debit Mandate ID.
	SEPADirectDebitMandateIDScheme ReferenceNumberScheme = "SDDM"

	// SwedishBankgiroOCRScheme indicates Swedish Bankgiro OCR.
	SwedishBankgiroOCRScheme ReferenceNumberScheme = "SEBG"
)

// SessionStatus represents status of a user session.
type SessionStatus string

const (
	// AuthorizedSessionStatus indicates the session is authorised for access to account information.
	AuthorizedSessionStatus SessionStatus = "AUTHORIZED"

	// CancelledSessionStatus indicates the session authorisation has been cancelled by the end-user.
	CancelledSessionStatus SessionStatus = "CANCELLED"

	// ClosedSessionStatus indicates the session has been closed by the application.
	ClosedSessionStatus SessionStatus = "CLOSED"

	// ExpiredSessionStatus indicates the session has expired.
	ExpiredSessionStatus SessionStatus = "EXPIRED"

	// InvalidSessionStatus indicates the session authorisation has failed.
	InvalidSessionStatus SessionStatus = "INVALID"

	// PendingAuthorizationSessionStatus indicates the session authorisation by the end-user is pending.
	PendingAuthorizationSessionStatus SessionStatus = "PENDING_AUTHORIZATION"

	// ReturnedFromBankSessionStatus indicates the session authorisation has completed successfully by the end-user.
	ReturnedFromBankSessionStatus SessionStatus = "RETURNED_FROM_BANK"

	// RevokedSessionStatus indicates the session has been revoked by the end-user.
	RevokedSessionStatus SessionStatus = "REVOKED"
)

// TransactionsFetchStrategy represents strategy for fetching transactions.
type TransactionsFetchStrategy string

const (
	// DefaultTransactionsFetchStrategy fetches transactions as requested by the user by passing the date_from and
	// date_to parameters to an ASPSP. If not date_from or date_to is passed, then meaningful defaults are used.
	DefaultTransactionsFetchStrategy TransactionsFetchStrategy = "default"

	// LongestTransactionsFetchStrategy tries to find the longest possible period of transactions and fetches
	// transactions for that period. Passed date_from is also taken into account. This strategy may use extra ASPSP
	// calls. date_to is ignored in this strategy.
	LongestTransactionsFetchStrategy TransactionsFetchStrategy = "longest"
)

// TransactionStatus represents the status of a transaction.
type TransactionStatus string

const (
	// AccountedTransactionStatus accounted transaction (ISO20022 Closing Booked).
	AccountedTransactionStatus TransactionStatus = "BOOK"

	// CancelledTransactionStatus cancelled transaction.
	CancelledTransactionStatus TransactionStatus = "CNCL"

	// HoldTransactionStatus account hold.
	HoldTransactionStatus TransactionStatus = "HOLD"

	// OtherTransactionStatus transaction with unknown status or not fitting the other options.
	OtherTransactionStatus TransactionStatus = "OTHR"

	// InstantBalanceTransactionStatus instant Balance Transaction (ISO20022 Expected).
	InstantBalanceTransactionStatus TransactionStatus = "PDNG"

	// RejectedTransactionStatus rejected transaction.
	RejectedTransactionStatus TransactionStatus = "RJCT"

	// ScheduledTransactionStatus scheduled transaction.
	ScheduledTransactionStatus TransactionStatus = "SCHD"
)

var transactionStatusDescriptions = map[TransactionStatus]string{
	AccountedTransactionStatus:      "Accounted transaction",
	CancelledTransactionStatus:      "Cancelled transaction",
	HoldTransactionStatus:           "Account hold",
	OtherTransactionStatus:          "Other transaction",
	InstantBalanceTransactionStatus: "Instant balance transaction",
	RejectedTransactionStatus:       "Rejected transaction",
	ScheduledTransactionStatus:      "Scheduled transaction",
}

// IsEmpty checks if the TransactionStatus is empty.
func (ts TransactionStatus) IsEmpty() bool {
	return ts == ""
}

// IsValid checks if the TransactionStatus is valid.
func (ts TransactionStatus) IsValid() bool {
	_, ok := transactionStatusDescriptions[ts]
	return ok
}

// Description returns the description of the TransactionStatus.
func (ts TransactionStatus) Description() string {
	if desc, ok := transactionStatusDescriptions[ts]; ok {
		return desc
	}

	return ""
}

// TransactionStatusDescriptions returns a map of TransactionStatus to their descriptions.
func TransactionStatusDescriptions() map[TransactionStatus]string {
	return transactionStatusDescriptions
}

// TransactionStatusKeys returns a slice of TransactionStatus as strings.
func TransactionStatusKeys() []string {
	keys := make([]string, 0, len(transactionStatusDescriptions))
	for k := range transactionStatusDescriptions {
		keys = append(keys, string(k))
	}
	return keys
}
