package cryptonx

import (
	"encoding/hex"
	"fmt"
)

/*
Decrypter decripts text based on nonce and key
*/
func Decrypter(key string, nonce string, text string) (decryptedText string, err error) {
	err = nil
	textocifrado, _ := hex.DecodeString(text)
	nonceInterno, _ := hex.DecodeString(nonce)

	//fmt.Printf("[Decrypter] texto: [%v] \n", texto)
	//fmt.Printf("[Decrypter] textocifrado: [%v] \n", textocifrado)
	//fmt.Printf("[Decrypter] Nonce: [%v] \n", nonce)
	//fmt.Printf("[Decrypter] nonceInterno: [%v] \n", nonceInterno)

	aesgcm, err := GetAEAD(key)
	if err != nil {
		fmt.Printf("[Decrypter] Error during aesgcm generation: [%s]\n", err.Error())
		return
	}

	retorno, err := aesgcm.Open(nil, nonceInterno, textocifrado, nil)
	if err != nil {
		fmt.Printf("[Decrypter] Error during opening text operation: [%s]\n", err.Error())
		return
	}
	decryptedText = fmt.Sprintf("%s", retorno)
	return

}
