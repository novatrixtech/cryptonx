package cryptonx

import (
	"crypto/rand"
	"fmt"
	"io"
)

/*
Encrypter Encrypts text based on AES Algorythim and return the encrypted text and his nonce
*/
func Encrypter(key string, text string) (encryptedText string, nonce string, err error) {
	err = nil
	nonceTemp := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, nonceTemp); err != nil {
		fmt.Printf("[Encrypter] Error during nonce generation: [%s]\n", err.Error())
		return
	}

	aesgcm, err := GetAEAD(key)
	if err != nil {
		fmt.Printf("[Encrypter] Error during AEAD generation: [%s]\n", err.Error())
		return
	}
	encryptedText = fmt.Sprintf("%x", aesgcm.Seal(nil, nonceTemp, []byte(text), nil))
	nonce = fmt.Sprintf("%x", nonceTemp)
	return
}
