package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	fmt.Println("\n=====Small Message=====\n")
	message := []byte("who are you? i am fine!")

	fmt.Println("MD5: %x\n\n", md5.Sum(message))
	fmt.Println("SHA1: %x\n\n", sha1.Sum(message))
	fmt.Println("SHA256: %x\n\n", sha256.Sum256(message))
	fmt.Println("SHA512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n-----Large Message-----\n")
	text := "先天下之忧而忧，后天下之乐而乐。落霞与孤鹜齐飞，秋水共长天一色。学而时习之，不亦说乎。欲穷千里目，更上一层楼。不睦君王天下事，此生逍遥吾自知。"
	message = []byte(text)

	fmt.Println("MD5: %x\n\n", md5.Sum(message))
	fmt.Println("SHA1: %x\n\n", sha1.Sum(message))
	fmt.Println("SHA256: %x\n\n", sha256.Sum256(message))
	fmt.Println("SHA512: %x\n\n", sha512.Sum512(message))
}
