package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

/*
Encrypter Encrypts text based on AES Algorythim and return the encrypted text and his nonce
*/
func Encrypter(key string, text string) (string, string, error) {

	nonce := make([]byte, 12)
	if _, erro := io.ReadFull(rand.Reader, nonce); erro != nil {
		fmt.Printf("Erro ao gerar o nonce: [%s]\n", erro.Error())
		return "", "", erro
	}

	aesgcm, erro := GetAEAD(key)
	if erro != nil {
		fmt.Printf("[Encrypter] Erro ao gerar o aesgcm: [%s]\n", erro.Error())
		return "", "", erro
	}

	return fmt.Sprintf("%x", aesgcm.Seal(nil, nonce, []byte(text), nil)), fmt.Sprintf("%x", nonce), nil
}
