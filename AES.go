package main

import (
	"io/ioutil"
	"encoding/pem"
	"crypto/rand"
	"log"
	"fmt"
	"crypto/cipher"
	"crypto/aes"
)

const (
	keyFile = "aes.key"
	encryptedFile = "aes.enc"
)

var IV = []byte("0123456789ABCDEF")

func readKey(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, _ := pem.Decode(key)
	return block.Bytes, nil
}

func createKey() []byte {
	genKey := make([]byte, 16)
	_, err := rand.Read(genKey)
	if err != nil {
		log.Fatalln("读取随机密钥失败：", err)
	}
	return genKey
}

func saveKey(filename string, key []byte) {
	block := &pem.Block{
		Type: "AES KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0644)
	if err != nil {
		log.Fatalln("保存密钥失败", filename, err)
	}
}

func aesKey() []byte {
	file := fmt.Sprintf(keyFile)
	key, err := readKey(file)
	if err != nil {
		log.Println("创建新的AES密钥")
		key = createKey()
		saveKey(file, key)
	}
	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		log.Fatalln("创建AES算法失败：", err)
	}
	return c
}

func encryption(plainText string) {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	err := ioutil.WriteFile(fmt.Sprintf(encryptedFile), bytes, 0644)
	if err != nil {
		log.Fatalln("写加密文件失败：", err)
	}else{
		fmt.Println("文件消息加密成功：", encryptedFile)
	}
}

func decryption() []byte {
	bytes, err := ioutil.ReadFile(fmt.Sprintf(encryptedFile))
	if err != nil {
		log.Fatalln("读取加密文件失败：", err)
	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}

func main() {
	var plainText = "i am bob nx!"
	encryption(plainText)
	fmt.Printf("%s", decryption())
}
