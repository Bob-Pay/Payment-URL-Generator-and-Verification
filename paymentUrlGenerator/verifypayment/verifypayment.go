package verifypayment

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Allowed IPs for Bob Pay	dev				sandbox
var allowedIPs = []string{"13.246.101.163", "13.245.58.93"}

// Notification represents the structure of the notification body.
type Notification struct {
	ID                   int     `json:"id"`
	UUID                 string  `json:"uuid"`
	ShortReference       string  `json:"short_reference"`
	FromBank             string  `json:"from_bank"`
	CustomPaymentID      string  `json:"custom_payment_id"`
	NotifyURL            string  `json:"notify_url"`
	SuccessURL           string  `json:"success_url"`
	PendingURL           string  `json:"pending_url"`
	CancelURL            string  `json:"cancel_url"`
	ItemName             string  `json:"item_name"`
	ItemDescription      string  `json:"item_description"`
	Amount               float64 `json:"amount"`
	Signature            string  `json:"signature"`
	TimeCreated          string  `json:"time_created"`
	AccountID            int     `json:"account_id"`
	AccountCode          string  `json:"account_code"`
	TransactingAsEmail   string  `json:"transacting_as_email"`
	TransactingAsMobile  string  `json:"transacting_as_mobile_number"`
	Status               string  `json:"status"`
	RecipientAccountCode string  `json:"recipient_account_code"`
	RecipientAccountID   int     `json:"recipient_account_id"`
	MobileNumber         string  `json:"mobile_number"`
	Email                string  `json:"email"`
	IsTest               bool    `json:"is_test"`
	PaymentMethod        string  `json:"payment_method"`
}

// VerifySourceIP checks if the source IP is allowed.
func VerifySourceIP(ip string) bool {
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}

// VerifySignature verifies the signature of the notification.
func VerifySignature(notification Notification, passphrase string) bool {
	params := []string{
		"recipient_account_code=" + url.QueryEscape(strings.ReplaceAll(notification.RecipientAccountCode, " ", "+")),
		"custom_payment_id=" + url.QueryEscape(strings.ReplaceAll(notification.CustomPaymentID, " ", "+")),
		"email=" + url.QueryEscape(strings.ReplaceAll(notification.Email, " ", "+")),
		"mobile_number=" + url.QueryEscape(strings.ReplaceAll(notification.MobileNumber, " ", "+")),
		"amount=" + url.QueryEscape(strings.ReplaceAll(fmt.Sprintf("%.2f", notification.Amount), " ", "+")),
		"item_name=" + url.QueryEscape(strings.ReplaceAll(notification.ItemName, " ", "+")),
		"item_description=" + url.QueryEscape(strings.ReplaceAll(notification.ItemDescription, " ", "+")),
		"notify_url=" + url.QueryEscape(strings.ReplaceAll(notification.NotifyURL, " ", "+")),
		"success_url=" + url.QueryEscape(strings.ReplaceAll(notification.SuccessURL, " ", "+")),
		"pending_url=" + url.QueryEscape(strings.ReplaceAll(notification.PendingURL, " ", "+")),
		"cancel_url=" + url.QueryEscape(strings.ReplaceAll(notification.CancelURL, " ", "+")),
	}
	stringToHash := strings.Join(params, "&") + "&passphrase=" + passphrase

	hash := md5.Sum([]byte(stringToHash))
	generatedHash := hex.EncodeToString(hash[:])

	// Normalize and compare
	return strings.ToLower(strings.TrimSpace(generatedHash)) == strings.ToLower(strings.TrimSpace(notification.Signature))
}

// ValidateAmount checks if the received amount matches the expected amount.
func ValidateAmount(receivedAmount, expectedAmount float64) bool {
	return receivedAmount == expectedAmount
}

// ValidateWithBobPay validates the payment with Bob Pay.
func ValidateWithBobPay(notificationBody []byte, sandbox bool) error {
	bobPayUrl := "https://api.bobpay.co.za/payments/intents/validate"
	if sandbox {
		bobPayUrl = "https://api.sandbox.bobpay.co.za/payments/intents/validate"
	}

	resp, err := http.Post(bobPayUrl, "application/json", bytes.NewBuffer(notificationBody))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("validation with Bob Pay failed")
	}
	return nil
}

// ProcessNotification processes the payment notification.
func ProcessNotification(notificationBody []byte, sourceIP string, passphrase string, expectedAmount float64, sandbox bool) error {
	// Verify source IP
	if !VerifySourceIP(sourceIP) {
		return errors.New("invalid source IP")
	}

	// Parse notification body
	var notification Notification
	if err := json.Unmarshal(notificationBody, &notification); err != nil {
		return err
	}

	// Verify signature
	if !VerifySignature(notification, passphrase) {
		return errors.New("invalid signature")
	}

	// Verify amount
	if !ValidateAmount(notification.Amount, expectedAmount) {
		return errors.New("amount mismatch")
	}

	// Validate with Bob Pay
	if err := ValidateWithBobPay(notificationBody, sandbox); err != nil {
		return err
	}

	// All validations passed
	return nil
}
