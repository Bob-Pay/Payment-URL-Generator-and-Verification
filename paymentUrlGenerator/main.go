package main

import (
	"awesomeProject/paymenturl"
	"fmt"
)

func main() {
	// Define the key-value pairs for the payment parameters
	kvPairs := []paymenturl.KVPair{
		{Key: "recipient_account_code", Value: "SAN001"},
		{Key: "custom_payment_id", Value: "478"},
		{Key: "email", Value: "testemail@bob.co.za"},
		{Key: "mobile_number", Value: ""},
		{Key: "amount", Value: "499.99"},
		{Key: "item_name", Value: "Order 1452"},
		{Key: "item_description", Value: "Lego Star Wars New Test"},
		{Key: "notify_url", Value: "https://api-sandbox.bobpay.co.za/payment/bobpay"},
		{Key: "success_url", Value: "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478"},
		{Key: "pending_url", Value: "https://sandbox.bobpay.co.za/accounts/payment-confirmation?provider=bobpay&id=478"},
		{Key: "cancel_url", Value: "https://sandbox.bobpay.co.za/accounts/payment-cancel?provider=bobpay&id=478"},
	}

	// Define the Bob Pay website URL and passphrase
	bobPayWebsiteURL := "https://sandbox.bobpay.co.za"
	passphrase := "0W5LORYafx"

	// Generate the payment URL
	paymentURL := paymenturl.GeneratePayURL(bobPayWebsiteURL, kvPairs, passphrase)
	fmt.Println("Generated Payment URL:", paymentURL)

}
