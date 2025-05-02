# BobPay Integration Package

This project provides a Go package for integrating with the BobPay payment system. It includes functionality for generating payment URLs, verifying payment notifications, and validating payments with BobPay.

## Features

- **Payment URL Generation**: Create secure payment URLs with signatures.
- **Notification Processing**: Verify and process payment notifications from BobPay.
- **Signature Verification**: Ensure the integrity of incoming notifications using MD5 signatures.
- **Amount Validation**: Validate the payment amount against the expected value.
- **BobPay Validation**: Validate payments directly with BobPay's API.

## Project Structure

```
paymentUrlGenerator/
├── go.mod
├── paymenturl/
│   ├── paymenturl.go
├── verifypayment/
│   ├── verifypayment.go
│   ├── verifypayment_test.go
├── main.go (optional for examples)
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/awesomeProject.git
   ```
2. Navigate to the project directory:
   ```bash
   cd awesomeProject
   ```
3. Initialize the Go module (if not already done):
   ```bash
   go mod tidy
   ```

## Usage

### 1. Generate a Payment URL

Use the `paymenturl.GeneratePayURL` function to create a payment URL.

```go
import "awesomeProject/paymenturl"

kvPairs := []paymenturl.KVPair{
    {Key: "amount", Value: "499.99"},
    {Key: "item_name", Value: "Order 1452"},
    {Key: "email", Value: "testemail@bob.co.za"},
}
passphrase := "your-passphrase"
url := paymenturl.GeneratePayURL("https://sandbox.bobpay.co.za", kvPairs, passphrase)
fmt.Println("Payment URL:", url)
```

### 2. Process a Payment Notification

Use the `verifypayment.ProcessNotification` function to verify and process notifications.

```go
import "awesomeProject/verifypayment"

notificationBody := []byte(`{...}`) // JSON notification from BobPay
sourceIP := "13.245.58.93"
passphrase := "your-passphrase"
expectedAmount := 499.99
sandbox := true

err := verifypayment.ProcessNotification(notificationBody, sourceIP, passphrase, expectedAmount, sandbox)
if err != nil {
    fmt.Println("Error processing notification:", err)
} else {
    fmt.Println("Notification processed successfully!")
}
```

## Testing

Run the tests using the following command:

```bash
go test ./...
```

## Configuration

- **Allowed IPs**: Update the `allowedIPs` variable in `verifypayment.go` to include the IPs allowed to send notifications.
- **Sandbox Mode**: Use the `sandbox` parameter in `ProcessNotification` and `ValidateWithBobPay` to toggle between sandbox and production environments.
