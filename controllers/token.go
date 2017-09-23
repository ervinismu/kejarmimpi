package controllers

import (
	"crypto/sha256"
	"encoding/base64"
)

//GenerateToken is for generate token using AES
func GenerateToken(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	h := sha256.New()
	h.Write([]byte(encoded))
	var sha = h.Sum(nil)
	var token = base64.StdEncoding.EncodeToString([]byte(sha))
	return token
}
