package main

import (
	"crypto/des"
	"log"
)

func main() {
	key := []byte("secretsa")

	cb, err := des.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	pt := []byte("plaintxt")

	et := make([]byte, des.BlockSize)
	cb.Encrypt(et, pt)
	log.Printf("暗号化した値は%s", et)

	dt := make([]byte, des.BlockSize)
	cb.Decrypt(dt, et)
	log.Printf("復号した値は%s", dt)
}
