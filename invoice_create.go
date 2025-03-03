// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type InvoiceCreate struct {

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// An automatic invoice means a corresponding transaction is run using the account's billing information at the same time the invoice is created. Manual invoices are created without a corresponding transaction. The merchant must enter a manual payment transaction or have the customer pay the invoice with an automatic method, like credit card, PayPal, Amazon, or ACH bank payment.
	CollectionMethod *string `json:"collection_method,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings for charge invoices. Specify custom notes to add or override Customer Notes on charge invoices.
	ChargeCustomerNotes *string `json:"charge_customer_notes,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings for credit invoices. Specify customer notes to add or override Customer Notes on credit invoices.
	CreditCustomerNotes *string `json:"credit_customer_notes,omitempty"`

	// Integer paired with `Net Terms Type` and representing the number
	// of days past the current date (for `net` Net Terms Type) or days after
	// the last day of the current month (for `eom` Net Terms Type) that the
	// invoice will become past due. For any value, an additional 24 hours is
	// added to ensure the customer has the entire last day to make payment before
	// becoming past due.  For example:
	// If an invoice is due `net 0`, it is due 'On Receipt' and will become past due 24 hours after it's created.
	// If an invoice is due `net 30`, it will become past due at 31 days exactly.
	// If an invoice is due `eom 30`, it will become past due 31 days from the last day of the current month.
	// When `eom` Net Terms Type is passed, the value for `Net Terms` is restricted to `0, 15, 30, 45, 60, or 90`.
	// For more information please visit our docs page (https://docs.recurly.com/docs/manual-payments#section-collection-terms)
	NetTerms *int `json:"net_terms,omitempty"`

	// Optionally supplied string that may be either `net` or `eom` (end-of-month).
	// When `net`, an invoice becomes past due the specified number of `Net Terms` days from the current date.
	// When `eom` an invoice becomes past due the specified number of `Net Terms` days from the last day of the current month.
	// This field is only available when the EOM Net Terms feature is enabled.
	NetTermsType *string `json:"net_terms_type,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber *string `json:"po_number,omitempty"`

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions.
	TermsAndConditions *string `json:"terms_and_conditions,omitempty"`

	// VAT Reverse Charge Notes only appear if you have EU VAT enabled or are using your own Avalara AvaTax account and the customer is in the EU, has a VAT number, and is in a different country than your own. This will default to the VAT Reverse Charge Notes text specified on the Tax Settings page in your Recurly admin, unless custom notes were created with the original subscription.
	VatReverseChargeNotes *string `json:"vat_reverse_charge_notes,omitempty"`
}
