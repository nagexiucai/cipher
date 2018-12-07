package main

import (
	"unicode"
	"fmt"
)

// 加密、解密字符串的接口
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}
// 密钥
type cipher []int

// 密码算法
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'Z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

// 加密
func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a,b int) int {return a + b})
}

// 解密
func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a,b int) int {return a - b})
}

// 凯撒移动密钥
func NewCaesar(key int) Cipher {
	return NewShift(key)
}
// 新建移动密钥
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipher([]int{shift})
	return &c
}

func main() {
	c := NewCaesar(1)
	fmt.Println("用密钥 01 加密 abcd：", c.Encryption("abcd"))
	fmt.Println("用密钥 01 解密 bcde：", c.Decryption("bcde"))

	c = NewCaesar(10)
	fmt.Println("Encrypt Key(10) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(10) klmn =>", c.Decryption("klmn"))
	fmt.Println()

	c = NewCaesar(15)
	fmt.Println("Encrypt Key(15) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(15) pqrs =>", c.Decryption("pqrs"))
}