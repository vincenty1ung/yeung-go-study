package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"

	aes2 "github.com/uncleyeung/yeung-go-study/crypto/aes"
)

var (
	key       string
	txt       string
	isEncrypt bool
)

func init() {
	flag.StringVar(&key, "k", "", "设置密钥，必须长度为16 48 64")
	flag.StringVar(&txt, "t", "", "设置文本")
	// -e 加密 不加-e 解密
	flag.BoolVar(&isEncrypt, "e", false, "默认加密")
	// flag.StringVar(&port, "p", "", "设置端口")
	flag.Parse()

}

func main() {
	flag.Args()
	if len(txt) == 0 && len(txt) == 0 {
		panic("设置关键参数")
	}
	if isEncrypt {
		Encrypt(key, txt)
	} else {
		Decrypt(key, txt)
	}
}

func InitNewCipher(key string) cipher.Block {
	a := []byte(key)
	aesBlock, _ := aes.NewCipher(a)
	return aesBlock
}

// 加密
func Encrypt(key, src string) {
	newCipher := InitNewCipher(key)
	encrypt, err := aes2.ECBEncrypt(newCipher, []byte(src))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(encrypt))
}

// 解密
func Decrypt(key, src string) {
	newCipher := InitNewCipher(key)
	decodeString, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	decrypt, err := aes2.ECBDecrypt(newCipher, decodeString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decrypt))
}
