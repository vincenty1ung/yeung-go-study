package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

var aesBlock cipher.Block

func init() {
	a := []byte("1111111111111111")
	aesBlock, _ = aes.NewCipher(a)
}

func TestAes(t *testing.T) {
	a := []byte("我爱你")
	b, err := ECBEncrypt(aesBlock, a)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", hex.EncodeToString(b))

}

func TestAesECBDecrypt(t *testing.T) {
	txt := "f8c62857c68ae76e192cf7bbceda489c"
	decodeString, err := hex.DecodeString(txt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	b, err := ECBDecrypt(aesBlock, decodeString)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", b)

	str := "2022-01-01 00:04:05"

	timeLayout := "2006-01-02 15:04:05"
	parse, _ := time.ParseInLocation(timeLayout, str, time.Local)
	now := time.Now()
	fmt.Println(now.Day())
	fmt.Println(parse.Day())
	fmt.Println(now.Day() - parse.Day())
	fmt.Println(now.Sub(parse).String())
	parseDuration, _ := time.ParseDuration("7288h52m37.108883s")
	fmt.Println(parseDuration.Hours())
}

/*
func BenchmarkAES(b *testing.B) {
	a := []byte("1111111111111111")
	o := make([]byte, 50)
	d, err := ECBEncrypt(a, o, a, padding.PKCS5)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	for i := 0; i < b.N; i++ {
		_, err := ECBDecrypt(d, o, a, padding.PKCS5)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}
*/
