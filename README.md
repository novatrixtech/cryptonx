# cryptonx
Crypto is a lib to encrypt and decrypt messages using Go Language. It's very simple and straightforward.

## Example
```
package main

import (
	"fmt"

	"github.com/novatrixtech/cryptonx"
)

func main() {
	var key = "Abcdefghijklmnopqrstuvxz12345678"
	var text = "This is my text"
	textEncrypted, nonce, err := cryptonx.Encrypter(key, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted text is: ", textEncrypted)
	fmt.Println("The nonce is: ", nonce)
	textDecrypted, err := cryptonx.Decrypter(key, nonce, textEncrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted text is: ", textDecrypted)
}
```

