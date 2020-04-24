package utils

import (
	"encoding/base64"
	"encoding/hex"
)

// ToBase64 ...
func ToBase64(s string) string {

	return base64.StdEncoding.EncodeToString([]byte(s))
}

// FromBase64 ...
func FromBase64(s string) string {

	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(str)
}

// ToHex ...
func ToHex(s string) string {

	return hex.EncodeToString([]byte(s))
}

// FromHex decodes a string from hex
func FromHex(s string) string {

	str, err := hex.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(str)
}
