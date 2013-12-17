package security

import (
	"encoding/base64"
)

var coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func EncodeBase64(src string) string {
	return coder.EncodeToString([]byte(src))
}

func DecodeBase64(src string) string {
	result, _ := coder.DecodeString(src)
	return string(result)
}
