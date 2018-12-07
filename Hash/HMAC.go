package main

import (
	"encoding/base64"
	"fmt"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"crypto/sha512"
)

// 本例限制在个人应用

var secret = "nagexiucai.com"

// 产生16比特的密码随机数据的盐
func salt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func main() {
	message := "Today our civilians are so happy!" // 今儿个老百姓、真呀真高兴。
	egg := salt()
	fmt.Println("Mesage:", message)
	fmt.Println("Salt:", egg)

	hash := hmac.New(sha256.New, []byte(secret))
	io.WriteString(hash, message+egg)
	fmt.Printf("HMAC-SHA256: %x\n", hash.Sum(nil))

	hash = hmac.New(sha512.New, []byte(secret))
	io.WriteString(hash, message+egg)
	fmt.Printf("HMAC-SHA512: %x\n", hash.Sum(nil))
}
