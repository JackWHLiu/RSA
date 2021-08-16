package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	var prikey, _ = rsa.GenerateKey(rand.Reader, 1024)
	var pubkey = &prikey.PublicKey
	fmt.Printf("私钥:%x\n", prikey.D)
	fmt.Printf("公钥:%x\n", pubkey.N)
	var msg = "Go语言公链开发"
	fmt.Printf("原消息:%s\n", msg)
	var cipherText, _ = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubkey, []byte(msg), nil)
	fmt.Printf("加密后:%x\n", cipherText)
	var decryptText, _ = rsa.DecryptOAEP(sha256.New(), rand.Reader, prikey, cipherText, nil)
	fmt.Printf("解密后:%s\n", decryptText)
}
