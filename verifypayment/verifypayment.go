package verifypayment

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bob-Pay/Payment-URL-Generator-and-Verification/paymentdetails"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VerifySourceIP checks if the source IP is allowed.
func VerifySourceIP(ip string, config paymentdetails.ValidationConfig) bool {
	for _, allowedIP := range config.AllowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}

// VerifySignature verifies the signature of the notification.
func VerifySignature(notification paymentdetails.PaymentNotification, passphrase string) bool {
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
func ValidateWithBobPay(notificationBody []byte, config paymentdetails.ValidationConfig, sandbox bool) error {
	bobPayUrl := config.BobPayValidationURL
	if sandbox {
		bobPayUrl = strings.Replace(bobPayUrl, "api.bobpay.co.za", "api.sandbox.bobpay.co.za", 1)
	}

	resp, err := http.Post(bobPayUrl, "application/json", bytes.NewBuffer(notificationBody))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("validation with Bob Pay failed")
	}
	return nil
}

func ProcessNotification(notificationBody []byte, sourceIP string, config paymentdetails.ValidationConfig, sandbox bool) error {
	// Verify source IP
	if !VerifySourceIP(sourceIP, config) {
		return errors.New("invalid source IP")
	}

	// Parse notification body
	var notification paymentdetails.PaymentNotification
	if err := json.Unmarshal(notificationBody, &notification); err != nil {
		return err
	}

	// Verify signature
	if !VerifySignature(notification, config.Passphrase) {
		return errors.New("invalid signature")
	}

	// Verify amount
	if config.ExpectedAmount != nil && !ValidateAmount(notification.Amount, *config.ExpectedAmount) {
		return errors.New("amount mismatch")
	}

	// Validate with Bob Pay
	if err := ValidateWithBobPay(notificationBody, config, sandbox); err != nil {
		return err
	}

	// All validations passed
	return nil
}
