package paymenturl

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Bob-Pay/Payment-URL-Generator-and-Verification/paymentdetails"
	"net/url"
	"strings"
)

// GeneratePayURL generates the payment URL.
func GeneratePayURL(config paymentdetails.PaymentConfig, kvPairs []paymentdetails.KeyValuePair) string {
	signature := GenerateSignature(kvPairs, config.Passphrase)
	params := []string{}

	for _, kv := range kvPairs {
		encodedValue := url.QueryEscape(strings.ReplaceAll(kv.Value, " ", "+"))
		params = append(params, kv.Key+"="+encodedValue)
	}

	queryString := strings.Join(params, "&")
	return config.BobPayWebsiteURL + "/pay?" + queryString + "&signature=" + signature
}

// GenerateSignature generates the MD5 signature.
func GenerateSignature(kvPairs []paymentdetails.KeyValuePair, passphrase string) string {
	params := []string{}

	for _, kv := range kvPairs {
		encodedValue := url.QueryEscape(strings.ReplaceAll(kv.Value, " ", "+"))
		params = append(params, kv.Key+"="+encodedValue)
	}

	stringToHash := strings.Join(params, "&") + "&passphrase=" + passphrase
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}
