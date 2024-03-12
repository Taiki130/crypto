package main

import (
	"crypto/aes"
	"log"
)

func main() {
	key := "1234567890123456" // 16bit
	cb, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	pt := []byte("hellohellohellohelloh") // 126bit
	et := make([]byte, aes.BlockSize)

	cb.Encrypt(et, pt)
	log.Printf("暗号化した値は%s", et)

	dt := make([]byte, aes.BlockSize)
	cb.Decrypt(dt, et)
	log.Printf("復号した値は%s", dt)
}
