package utils

import (
	"encoding/hex"
	"fmt"
)

/*
Decrypter decripts text based on nonce and key
*/
func Decrypter(key string, nonce string, text string) (string, error) {

	textocifrado, _ := hex.DecodeString(text)
	nonceInterno, _ := hex.DecodeString(nonce)

	//fmt.Printf("[Decrypter] texto: [%v] \n", texto)
	//fmt.Printf("[Decrypter] textocifrado: [%v] \n", textocifrado)
	//fmt.Printf("[Decrypter] Nonce: [%v] \n", nonce)
	//fmt.Printf("[Decrypter] nonceInterno: [%v] \n", nonceInterno)

	aesgcm, erro := GetAEAD(key)
	if erro != nil {
		fmt.Printf("[Decrypter] Erro ao gerar o aesgcm: [%s]\n", erro.Error())
		return "", erro
	}

	retorno, erro := aesgcm.Open(nil, nonceInterno, textocifrado, nil)
	if erro != nil {
		fmt.Printf("[Decrypter] Erro ao abrir o texto: [%s]\n", erro.Error())
		return "", erro
	}

	return fmt.Sprintf("%s", retorno), nil

}
