package verifypayment

import (
	"encoding/json"
	"testing"
)

func TestVerifySourceIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"13.246.115.225", true},
		{"13.246.100.25", true},
		{"192.168.1.1", false},
	}

	for _, test := range tests {
		result := VerifySourceIP(test.ip)
		if result != test.expected {
			t.Errorf("VerifySourceIP(%s) = %v; want %v", test.ip, result, test.expected)
		}
	}
}

func TestVerifySignature(t *testing.T) {
	notification := Notification{
		RecipientAccountCode: "SAN001",
		CustomPaymentID:      "478",
		Email:                "customer@bob.co.za",
		MobileNumber:         "",
		Amount:               499.99,
		ItemName:             "Order 1452",
		ItemDescription:      "Lego Star Wars",
		NotifyURL:            "https://api-sandbox.bobpay.co.za/payment/bobpay",
		SuccessURL:           "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478",
		PendingURL:           "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478",
		CancelURL:            "https://sandbox.bobpay.co.za/accounts/payment-cancel?provider=bobpay&id=478",
		Signature:            "b57d2e8b09e17c9977a82c5095ff8c8e",
	}

	passphrase := "0W5LORYafx"
	if !VerifySignature(notification, passphrase) {
		t.Errorf("VerifySignature failed for valid signature")
	}
}

func TestProcessNotification(t *testing.T) {
	notification := Notification{
		ID:                   55645,
		UUID:                 "1ff43e5d-9477-4a42-a6d8-7a444095b78c",
		ShortReference:       "3S7QD",
		FromBank:             "standard-bank",
		CustomPaymentID:      "478",
		NotifyURL:            "https://api-sandbox.bobpay.co.za/payment/bobpay",
		SuccessURL:           "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478",
		PendingURL:           "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478",
		CancelURL:            "https://sandbox.bobpay.co.za/accounts/payment-cancel?provider=bobpay&id=478",
		ItemName:             "Order+1452",
		ItemDescription:      "Lego+Star+Wars+New+Test",
		Amount:               499.99,
		Signature:            "9f3993aeaf6f3f189f8aa4f5feac2de8",
		TimeCreated:          "2025-04-30T16:01:32.557184+02:00",
		AccountID:            517,
		AccountCode:          "AUT162",
		TransactingAsEmail:   "leo+admin@bob.co.za",
		TransactingAsMobile:  "",
		Status:               "paid",
		RecipientAccountCode: "SAN001",
		RecipientAccountID:   236,
		MobileNumber:         "",
		Email:                "testemail@bob.co.za",
		IsTest:               false,
		PaymentMethod:        "credit-card",
	}

	notificationBody, _ := json.Marshal(notification)
	sourceIP := "13.245.58.93"
	passphrase := "0W5LORYafx"
	expectedAmount := 499.99
	sandbox := true

	err := ProcessNotification(notificationBody, sourceIP, passphrase, expectedAmount, sandbox)
	if err != nil {
		t.Errorf("ProcessNotification failed: %v", err)
	}
}
