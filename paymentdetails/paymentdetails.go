package paymentdetails

// KeyValuePair represents a key-value pair.
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PaymentConfig represents the configuration for a payment.
type PaymentConfig struct {
	BobPayWebsiteURL string `json:"bobPayWebsiteURL"`
	Passphrase       string `json:"passphrase"`
	NotifyURL        string `json:"notifyUrl"`
	SuccessURL       string `json:"successUrl"`
	PendingURL       string `json:"pendingUrl"`
	CancelURL        string `json:"cancelUrl"`
}

// PaymentDetails represents the details of a payment.
type PaymentDetails struct {
	RecipientAccountCode string            `json:"recipient_account_code"`
	CustomPaymentID      string            `json:"custom_payment_id"`
	Email                *string           `json:"email,omitempty"`
	MobileNumber         *string           `json:"mobile_number,omitempty"`
	Amount               string            `json:"amount"`
	ItemName             string            `json:"item_name"`
	ItemDescription      string            `json:"item_description"`
	AdditionalFields     map[string]string `json:"-"` // For dynamic fields
}

// PaymentNotification represents a payment notification.
type PaymentNotification struct {
	ID                        int     `json:"id"`
	UUID                      string  `json:"uuid"`
	ShortReference            string  `json:"short_reference"`
	FromBank                  string  `json:"from_bank"`
	CustomPaymentID           string  `json:"custom_payment_id"`
	NotifyURL                 string  `json:"notify_url"`
	SuccessURL                string  `json:"success_url"`
	PendingURL                string  `json:"pending_url"`
	CancelURL                 string  `json:"cancel_url"`
	ItemName                  string  `json:"item_name"`
	ItemDescription           string  `json:"item_description"`
	Amount                    float64 `json:"amount"`
	Signature                 string  `json:"signature"`
	TimeCreated               string  `json:"time_created"`
	AccountID                 int     `json:"account_id"`
	AccountCode               string  `json:"account_code"`
	TransactingAsEmail        string  `json:"transacting_as_email"`
	TransactingAsMobileNumber string  `json:"transacting_as_mobile_number"`
	Status                    string  `json:"status"`
	RecipientAccountCode      string  `json:"recipient_account_code"`
	RecipientAccountID        int     `json:"recipient_account_id"`
	MobileNumber              string  `json:"mobile_number"`
	Email                     string  `json:"email"`
	IsTest                    bool    `json:"is_test"`
	PaymentMethod             string  `json:"payment_method"`
}

// ValidationConfig represents the configuration for validation.
type ValidationConfig struct {
	Passphrase          string   `json:"passphrase"`
	ExpectedAmount      *float64 `json:"expectedAmount,omitempty"`
	AllowedIPs          []string `json:"allowedIps"`
	BobPayValidationURL string   `json:"bobPayValidationUrl"`
}
