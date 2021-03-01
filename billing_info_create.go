// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type BillingInfoCreate struct {

	// A token [generated by Recurly.js](https://developers.recurly.com/reference/recurly-js/#getting-a-token).
	TokenId *string `json:"token_id,omitempty"`

	// First name
	FirstName *string `json:"first_name,omitempty"`

	// Last name
	LastName *string `json:"last_name,omitempty"`

	// Company name
	Company *string `json:"company,omitempty"`

	Address *AddressCreate `json:"address,omitempty"`

	// Credit card number, spaces and dashes are accepted.
	Number *string `json:"number,omitempty"`

	// Expiration month
	Month *string `json:"month,omitempty"`

	// Expiration year
	Year *string `json:"year,omitempty"`

	// *STRONGLY RECOMMENDED*
	Cvv *string `json:"cvv,omitempty"`

	// VAT number
	VatNumber *string `json:"vat_number,omitempty"`

	// *STRONGLY RECOMMENDED* Customer's IP address when updating their billing information.
	IpAddress *string `json:"ip_address,omitempty"`

	// A token used in place of a credit card in order to perform transactions. Must be used in conjunction with `gateway_code`.
	GatewayToken *string `json:"gateway_token,omitempty"`

	// An identifier for a specific payment gateway. Must be used in conjunction with `gateway_token`.
	GatewayCode *string `json:"gateway_code,omitempty"`

	// Amazon billing agreement ID
	AmazonBillingAgreementId *string `json:"amazon_billing_agreement_id,omitempty"`

	// PayPal billing agreement ID
	PaypalBillingAgreementId *string `json:"paypal_billing_agreement_id,omitempty"`

	// Fraud Session ID
	FraudSessionId *string `json:"fraud_session_id,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType *string `json:"transaction_type,omitempty"`

	// A token generated by Recurly.js after completing a 3-D Secure device fingerprinting or authentication challenge.
	ThreeDSecureActionResultTokenId *string `json:"three_d_secure_action_result_token_id,omitempty"`

	// The International Bank Account Number, up to 34 alphanumeric characters comprising a country code; two check digits; and a number that includes the domestic bank account number, branch identifier, and potential routing information
	Iban *string `json:"iban,omitempty"`

	// Tax identifier is required if adding a billing info that is a consumer card in Brazil. This would be the customer's CPF, CPF is a Brazilian tax identifier for all tax paying residents.
	TaxIdentifier *string `json:"tax_identifier,omitempty"`

	// this field and a value of 'cpf' are required if adding a billing info that is an elo or hipercard type in Brazil.
	TaxIdentifierType *string `json:"tax_identifier_type,omitempty"`

	// The `primary_payment_method` indicator is used to designate the primary billing info on the account. The first billing info created on an account will always become primary. Adding additional billing infos provides the flexibility to mark another billing info as primary, or adding additional non-primary billing infos. This can be accomplished by passing the `primary_payment_method` indicator. When adding billing infos via the billing_info and /accounts endpoints, this value is not permitted, and will return an error if provided.
	PrimaryPaymentMethod *bool `json:"primary_payment_method,omitempty"`
}