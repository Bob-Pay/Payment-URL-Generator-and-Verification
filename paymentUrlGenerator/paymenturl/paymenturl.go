package paymenturl

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"
)

// KVPair represents a key-value pair.
type KVPair struct {
	Key   string
	Value string
}

// GeneratePayURL generates the payment URL.
func GeneratePayURL(bobPayWebsiteURL string, kvPairs []KVPair, passphrase string) string {
	signature := GenerateSignature(kvPairs, passphrase)
	params := []string{}

	for _, kv := range kvPairs {
		encodedValue := url.QueryEscape(strings.ReplaceAll(kv.Value, " ", "+"))
		params = append(params, kv.Key+"="+encodedValue)
	}

	queryString := strings.Join(params, "&")
	return bobPayWebsiteURL + "/pay?" + queryString + "&signature=" + signature
}

// GenerateSignature generates the MD5 signature.
func GenerateSignature(kvPairs []KVPair, passphrase string) string {
	params := []string{}

	for _, kv := range kvPairs {
		encodedValue := url.QueryEscape(strings.ReplaceAll(kv.Value, " ", "+"))
		params = append(params, kv.Key+"="+encodedValue)
	}

	stringToHash := strings.Join(params, "&") + "&passphrase=" + passphrase
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}
