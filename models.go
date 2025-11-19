package enablebankinggo

import (
	"net/http"
	"time"
)

// Access represents the access rights requested for an authorized session.
type Access struct {
	// Accounts indicates whether access to accounts is requested.
	Accounts []*AccountIdentification `json:"accounts,omitempty"`

	// Balances indicates whether to request consent with balances access.
	Balances bool `json:"balances,omitempty"`

	// Transactions indicates whether to request consent with transactions access.
	Transactions bool `json:"transactions,omitempty"`

	// ValidUntil specifies the date and time until which the authorised session
	// remains valid. The value must be in the RFC3339 date-time format with a timezone
	// offset, e.g. 2025-12-01T12:00:00.000000+00:00. The provided value cannot exceed
	// the date and time, calculated as "now" + maximum_consent_validity (provided in
	// seconds for each ASPSP in response to the GET /aspsps request). The provided value
	// is subject to adjustment to comply with the ASPSP's requirements. Specifically, if
	// the provided value is less than the minimum consent validity allowed by the ASPSP
	// (e.g., some ASPSPs require a minimum of 1 hour or 1 day), the consent validity will
	// be adjusted to meet these requirements. However, the session validity will remain
	// exactly as specified. This means that even if the consent remains valid on the
	// ASPSP's side, the session will expire based on the initially provided value.
	ValidUntil string `json:"valid_until"`
}

// AccountIdentification represents account identification used to identify an account.
type AccountIdentification struct {
	// IBAN is the International Bank Account Number (IBAN) - identification used internationally by financial
	// institutions to uniquely identify the account of a customer. Further specifications of the format and
	// content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services -
	// International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN string `json:"iban,omitempty"`

	// Other is other identification if IBAN is not provided.
	Other *GenericIdentification `json:"other,omitempty"`
}

// AccountResource represents an authorized account.
type AccountResource struct {
	// AccountID is the primary account identifier.
	AccountID *AccountIdentification `json:"account_id,omitempty"`

	// AllAccountIDs is the list of all account identifiers provided by ASPSPs (including
	// primary identifier available in the AccountID field).
	AllAccountIDs []*GenericIdentification `json:"all_account_ids,omitempty"`

	// AccountServicer represents information about the financial institution servicing the account.
	AccountServicer *FinancialInstitutionIdentification `json:"account_servicer,omitempty"`

	// Name is the account holder(s) name.
	Name string `json:"name,omitempty"`

	// Details is the account description set by PSU or provided by ASPSP.
	Details string `json:"details,omitempty"`

	// Usage specifies the usage of the account.
	Usage Usage `json:"usage,omitempty"`

	// CashAccountType specifies the type of the account.
	CashAccountType CashAccountType `json:"cash_account_type"`

	// Product is the Product Name of the Bank for this account, proprietary definition.
	Product string `json:"product,omitempty"`

	// Currency specifies the currency of the account.
	Currency string `json:"currency"`

	// PSUStatus is the relationship between the PSU and the account - Account Holder - Co-account Holder - Attorney.
	PSUStatus string `json:"psu_status,omitempty"`

	// CreditLimit specifies the credit limit of the account.
	CreditLimit *AmountType `json:"credit_limit,omitempty"`

	// LegalAge Specifies whether Enable Banking is confident that the account holder is of legal age or is a minor.
	// The field takes the following values:
	// true if the account holder is of legal age;
	// false if the account holder is a minor;
	// null (or the field is not set) if it is not possible to determine whether the account holder is of legal age
	// or a minor or if the legal age check is not applicable (in cases such as if the account holder is a legal entity
	// or there are multiple account co-holders).
	LegalAge *bool `json:"legal_age,omitempty"`

	// PostalAddress is the postal address of the account holder.
	PostalAddress *PostalAddress `json:"postal_address,omitempty"`

	// UID is the Unique account identificator used for fetching account balances and transactions. It is valid only until
	// the session to which the account belongs is in the AUTHORIZED status. It can be not set in case it is know that it
	// is not possible to fetch balances and transactions for the account (for example, in case the account is blocked or
	// closed at the ASPSP side).
	UID string `json:"uid,omitempty"`

	// IdentificationHash is the primary account identification hash. It can be used for matching accounts between multiple
	// sessions (even in case the sessions are authorized by different PSUs).
	IdentificationHash string `json:"identification_hash"`

	// IdentificationHashes list of possible account identification hashes. Identification hash is based on the account number.
	// Some accounts may have multiple account numbers (e.g. IBAN and BBAN). This field contains all possible hashes. Not all
	// of these hashes can be used to uniquely identify an account and that the primary goal of them is to be able to fuzzy
	// matching of accounts by certain properties. Primary hash is included in this list.
	IdentificationHashes []string `json:"identification_hashes"`
}

// AmountType represents an amount with currency.
type AmountType struct {
	// Amount is the numerical value or monetary figure associated with a particular transaction, representing balance on an
	// account, a fee or similar. Represented as a decimal number, using . (dot) as a decimal separator. Allowed precision
	// (number of digits after the decimal separator) varies depending on the currency and is validated differently depending
	// on the context.
	Amount string `json:"amount"`

	// Currency is the currency code in ISO 4217 format.
	Currency string `json:"currency"`
}

// ASPSP represents an ASPSP.
type ASPSP struct {
	// Name is the name of the ASPSP (i.e. a bank or a similar financial institution).
	Name string `json:"name"`

	// Country is the two-letter ISO 3166 code of the country, in which ASPSP operates.
	Country string `json:"country"`
}

// ASPSPData represents detailed information about an ASPSP.
type ASPSPData struct {
	// Name is the name of the ASPSP (i.e. a bank or a similar financial institution).
	Name string `json:"name"`

	// Country is the two-letter ISO 3166 code of the country, in which ASPSP operates.
	Country string `json:"country"`

	// Logo is the ASPSP logo URL. It is possible to transform (e.g. resize) the logo by
	// adding special suffixes at the end of the URL.
	// For example, -/resize/500x/. For full list of possible transformations, please refer
	// to https://uploadcare.com/docs/transformations/image/.
	Logo string `json:"logo"`

	// PSUTypes is the list of PSU types supported by the ASPSP.
	PSUTypes []PSUType `json:"psu_types"`

	// AuthMethods is the list of available authentication methods. Provided in case multiple
	// methods are available or it is possible to supply authentication credentials while
	// initiating authorization.
	AuthMethods []*AuthMethod `json:"auth_methods"`

	// MaximumConsentValidity is the maximum consent validity which bank supports in
	// seconds.
	MaximumConsentValidity int64 `json:"maximum_consent_validity"`

	// Beta flag indicates whether implementation is in beta mode.
	Beta bool `json:"beta"`

	// BIC is the Bank Identifier Code (BIC) of the ASPSP, if available.
	BIC string `json:"bic,omitempty"`

	// RequiredPSUHeaders is the list of the headers required to indicate to data retrieval
	// endpoints that PSU is online. Either all required PSU headers or none of them are to
	// be provided, otherwise `PSU_HEADER_NOT_PROVIDED` error will be returned.
	RequiredPSUHeaders []string `json:"required_psu_headers,omitempty"`

	// Group is the group which the ASPSP belongs to, if available.
	Group *ASPSPGroup `json:"group,omitempty"`
}

// ASPSPGroup represents group which the ASPSP belongs to.
type ASPSPGroup struct {
	// Name is the name of the group, which the ASPSP belongs to.
	Name string `json:"name"`

	// Logo is the URL of the logo for the group to which the ASPSP belongs.
	// This URL supports the same transformation postfixes as ASPSP logo URLs.
	Logo string `json:"logo"`
}

// AuthMethod represents an authentication method supported by ASPSP.
type AuthMethod struct {
	// Name is the internal name of the authentication method.
	Name string `json:"name,omitempty"`

	// Title is the human-readable title of the authentication method.
	Title string `json:"title,omitempty"`

	// PSUType is the type to which the authentication method is applicable.
	PSUType PSUType `json:"psu_type"`

	// Credentials is the list of credentials which are possible to supply while initiating
	// authorization.
	Credentials []*Credential `json:"credentials,omitempty"`

	// Approach is the authentication approach used in the current authentication method.
	Approach AuthenticationApproach `json:"approach"`

	// HiddenMethod flag indicates whether the current authentication method is hidden from the user.
	// If true then the user will not be able to select this authentication method. It is only
	// possible to select this authentication method via API.
	HiddenMethod bool `json:"hidden_method"`
}

type BalanceResource struct {
	// Name is the name of the balance.
	Name string `json:"name"`

	// BalanceAmount represents the structure aiming to embed the amount and the currency to be used.
	BalanceAmmount *AmountType `json:"balance_amount"`

	// BalanceType specifies the type of balance.
	BalanceType BalanceType `json:"balance_type"`

	// LastChangeDateTime is the date and time when the balance was last changed.
	LastChangeDateTime *time.Time `json:"last_change_date_time,omitempty"`

	// ReferenceDate is the reference date for the balance.
	ReferenceDate string `json:"reference_date,omitempty"`

	// LastCommittedTransaction is the entry reference of the last transaction contributing to the balance value.
	LastCommittedTransaction string `json:"last_committed_transaction,omitempty"`
}

// BankTransactionCode allows the account servicer to correctly report a transaction,
// which in its turn will help account holders to perform their cash management and
// reconciliation operations.
type BankTransactionCode struct {
	// Description is arbitrary transaction categorization description.
	Description string `json:"description,omitempty"`

	// Code specifies the family of a transaction within the domain.
	Code string `json:"code,omitempty"`

	// SubCode specifies the sub-product family of a transaction within a specific family.
	SubCode string `json:"sub_code,omitempty"`
}

// ClearingSystemMemberIdentification represents information used to identify a member within a clearing system.
type ClearingSystemMemberIdentification struct {
	// ClearingSystemID is the specification of a pre-agreed offering between clearing agents or the
	// channel through which the payment instruction is processed.
	ClearingSystemID string `json:"clearing_system_id,omitempty"`

	// MemberID is the member identification within the clearing system.
	MemberID string `json:"member_id,omitempty"`
}

// ContactDetails specifies the contact details associated with a person or an organisation.
type ContactDetails struct {
	// EmailAddress is the email address of the contact.
	EmailAddress string `json:"email_address,omitempty"`

	// PhoneNumber is the phone number of the contact.
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Credential struct {
	// Name is the internal name of the credential. The name is to be used when passing
	// credentials to the "start user authorization" request.
	Name string `json:"name"`

	// Title is the title for the credential to be displayed to PSU.
	Title string `json:"title"`

	// Required indicates whether the credential is required.
	Required bool `json:"required"`

	// Description is the description of the credential to be displayed to PSU.
	Description string `json:"description,omitempty"`

	// Template is the Perl compatible regular expression used for check of the credential
	// format.
	Template string `json:"template,omitempty"`
}

// ExchangeRate provides details on the currency exchange rate and contract.
type ExchangeRate struct {
	// UnitCurrency is the ISO 4217 code of the currency, in which the rate of exchange is expressed
	// in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency string `json:"unit_currency,omitempty"`

	// ExchangeRate is the factor used for conversion of an amount from one currency to another.
	// This reflects the price at which one currency was bought with another currency.
	ExchangeRate string `json:"exchange_rate,omitempty"`

	// RateType specifies the type of exchange rate applied to the transaction
	RateType RateType `json:"rate_type,omitempty"`

	// ContractIdentification is the unique and unambiguous reference to the foreign exchange contract
	// agreed between the initiating party/creditor and the debtor agent.
	ContractIdentification string `json:"contract_identification,omitempty"`

	// InstructedAmount is the original amount, in which transaction was initiated. In particular,
	// for cross-currency card transactions, the value represents original value of a purchase or a withdrawal
	// in a currency different from the card's native or default currency.
	InstructedAmount *AmountType `json:"instructed_amount,omitempty"`
}

// FinancialInstitutionIdentification represents information used to identify a financial institution.
type FinancialInstitutionIdentification struct {
	// BICFI is the code allocated to a financial institution by the ISO 9362 Registration Authority
	// as described in ISO 9362 "Banking - Banking telecommunication messages - Business identification code (BIC)".
	BICFI string `json:"bic_fi,omitempty"`

	// ClearingSystemMemberID represents information used to identify a member within a clearing system.
	ClearingSystemMemberID *ClearingSystemMemberIdentification `json:"clearing_system_member_id,omitempty"`

	// Name is the name of the financial institution.
	Name string `json:"name,omitempty"`
}

// GenericIdentification represents generic identification scheme used to identify an account.
type GenericIdentification struct {
	// Identification is the identification of the account using other scheme than IBAN.
	Identification string `json:"identification"`

	// SchemeName is the name of the identification scheme.
	SchemeName string `json:"scheme_name"`

	// Issuer is the name of the identification issuer.
	Issuer string `json:"issuer,omitempty"`
}

// Header represents additional headers to include in the request.
type Header map[HeaderKey]string

func NewHeaders() Header {
	return make(Header)
}

// Set sets the header key-value pair to the Header map.
func (h Header) Set(key HeaderKey, value string) {
	h[key] = value
}

func (h Header) FillHTTPHeader(httpHeader http.Header) {
	for key, value := range h {
		httpHeader.Set(string(key), value)
	}
}

type PartyIdentification struct {
	// Name name by which a party is known and which is usually used to identify that party..
	Name string `json:"name,omitempty"`

	// PostalAddress information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress `json:"postal_address,omitempty"`

	// OrganizationID unique identification of an account, a person or an organisation, as assigned by an issuer.
	OrganizationID *GenericIdentification `json:"organization_id,omitempty"`

	// PrivateID unique identification of an account, a person or an organisation, as assigned by an issuer.
	PrivateID *GenericIdentification `json:"private_id,omitempty"`

	// ContactDetails specifies the contact details associated with a person or an organisation.
	ContactDetails *ContactDetails `json:"contact_details,omitempty"`
}

// PostalAddress represents a postal address.
type PostalAddress struct {
	// AddressType is the type of address.
	AddressType AddressType `json:"address_type,omitempty"`

	// Department is the identification of a division of a large organisation or building.
	Department string `json:"department,omitempty"`

	// SubDepartment is the identification of a sub-division of a large organisation or building.
	SubDepartment string `json:"sub_department,omitempty"`

	// StreetName is the name of the street.
	StreetName string `json:"street_name,omitempty"`

	// BuildingNumber is the building number.
	BuildingNumber string `json:"building_number,omitempty"`

	// PostCode is the identifier consisting of a group of letters and/or numbers that is
	// added to a postal address to assist the sorting of mail.
	PostCode string `json:"post_code,omitempty"`

	// TownName is the name of a built-up area, with defined boundaries, and a local government.
	TownName string `json:"town_name,omitempty"`

	// CountrySubDivision identifies a subdivision of a country such as state, region, county.
	CountrySubDivision string `json:"country_sub_division,omitempty"`

	// Country is the two-letter ISO 3166 code of the country in which a person resides (the place
	// of a person's home). In the case of a company, it is the country from which the affairs of
	// that company are directed..
	Country string `json:"country,omitempty"`

	// AddressLines is the unstructured address. The two lines must embed zip code and town name.
	AddressLines []string `json:"address_lines,omitempty"`
}

// SessionAccount represents account data stored in the user session.
type SessionAccount struct {
	// UID is the account identificator within the session.
	UID string `json:"uid"`

	// IdentificationHash is the global account identification hash.
	IdentificationHash string `json:"identification_hash"`

	// IdentificationHashes List of possible account identification hashes. Identification hash is
	// based on the account number. Some accounts may have multiple account numbers (e.g. IBAN and
	// BBAN). This field contains all possible hashes.
	IdentificationHashes []string `json:"identification_hashes"`
}

// Transaction represents an account transaction resource.
type Transaction struct {
	// EntryReference is the unique transaction identifier provided by ASPSP. This identifier is both unique
	// and immutable for accounts with the same identification hashes and can be used for matching
	// transactions across multiple PSU authentication sessions. Usually the same identifier is
	// available for transactions in ASPSP's online/mobile interface and is called archive ID or
	// similarly. Please note that this identifier is not globally unique and same entry references
	// are likely to occur for transactions belonging to different accounts.
	EntryReference string `json:"entry_reference,omitempty"`

	// MerchantCategoryCode is the category code conform to ISO 18245, related to the type of services
	// or goods the merchant provides for the transaction
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`

	// TransactionAmount is the monetary sum of the transaction
	TransactionAmount *AmountType `json:"transaction_amount"`

	// Creditor is the identification of the party receiving funds in the transaction
	Creditor *PartyIdentification `json:"creditor,omitempty"`

	// CreditorAccount is the identification of the account on which the transaction is credited
	CreditorAccount *AccountIdentification `json:"creditor_account,omitempty"`

	// CreditorAgent is the identification of the creditor agent
	CreditorAgent *FinancialInstitutionIdentification `json:"creditor_agent,omitempty"`

	// Debtor is the identification of the party sending funds in the transaction
	Debtor *PartyIdentification `json:"debtor,omitempty"`

	// DebtorAccount is the identification of the account on which the transaction is debited
	DebtorAccount *AccountIdentification `json:"debtor_account,omitempty"`

	// DebtorAgent is the identification of the debtor agent
	DebtorAgent *FinancialInstitutionIdentification `json:"debtor_agent,omitempty"`

	// BankTransactionCode allows the account servicer to correctly report a transaction,
	// which in its turn will help account holders to perform their cash management and
	// reconciliation operations.
	BankTransactionCode *BankTransactionCode `json:"bank_transaction_code,omitempty"`

	// CreditDebitIndicator is the accounting flow of the transaction
	CreditDebitIndicator CreditDebitIndicator `json:"credit_debit_indicator"`

	// Status is the available transaction status values
	Status TransactionStatus `json:"status"`

	// BookingDate is the booking date of the transaction on the account, i.e. the date at which
	// the transaction has been recorded on books
	BookingDate string `json:"booking_date,omitempty"`

	// ValueDate is the value date of the transaction on the account, i.e. the date at which funds
	// become available to the account holder (in case of a credit transaction), or cease to be
	// available to the account holder (in case of a debit transaction)
	ValueDate string `json:"value_date,omitempty"`

	// TransactionDate is the date used for specific purposes:
	// - for card transaction: date of the transaction
	// - for credit transfer: acquiring date of the transaction
	// - for direct debit: receiving date of the transaction
	TransactionDate string `json:"transaction_date,omitempty"`

	// BalanceAfterTransaction is the funds on the account after execution of the transaction
	BalanceAfterTransaction *AmountType `json:"balance_after_transaction,omitempty"`

	// ReferenceNumber is the credit transfer reference number (also known as the creditor reference or the structured
	// creditor reference). The value is set when it is known that the transaction data contains a reference number
	// (in either ISO 11649 or another format). If the format is known it is provided in the reference_number_schema field.
	ReferenceNumber string `json:"reference_number,omitempty"`

	// ReferenceNumberSchema indicates what kind of reference number is used.
	ReferenceNumberSchema ReferenceNumberScheme `json:"reference_number_schema,omitempty"`

	// RemittanceInformation is the payment details. For credit transfers may contain free text,
	// reference number or both at the same time (in case Extended Remittance Information is supported).
	// When it is known that remittance information contains a reference number (either based on
	// ISO 11649 or a local scheme), the reference number is also available via the reference_number field.
	RemittanceInformation []string `json:"remittance_information,omitempty"`

	// DebtorAccountAdditionalIdentification All other debtor account identifiers provided by ASPSPs
	DebtorAccountAdditionalIdentification []*GenericIdentification `json:"debtor_account_additional_identification,omitempty"`

	// CreditorAccountAdditionalIdentification All other creditor account identifiers provided by ASPSPs
	CreditorAccountAdditionalIdentification []*GenericIdentification `json:"creditor_account_additional_identification,omitempty"`

	// ExchangeRate provides details on the currency exchange rate and contract.
	ExchangeRate *ExchangeRate `json:"exchange_rate,omitempty"`

	// Note is the internal note made by PSU
	Note string `json:"note,omitempty"`

	// TransactionID is the identification used for fetching transaction details.
	// This value can not be used to uniquely identify transactions and may change
	// if the list of transactions is retrieved again. Null if fetching transaction
	// details is not supported.
	TransactionID string `json:"transaction_id,omitempty"`
}
