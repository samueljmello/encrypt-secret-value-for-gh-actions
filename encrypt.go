package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/nacl/box"
)

func main() {

	if len(os.Args) != 3 {
		ExitOnError(errors.New("Expected 2 arguments: <value-to-encrypt> and <public-key>"))
	}

	// get values from arguments
	valueToEncrypt := os.Args[1:][0]
	publicKey := os.Args[2:][0]

	// print value
	fmt.Println(fmt.Sprintf("Encrypting: %s", valueToEncrypt))

	// encrypt with public key
	encryptedValue, err := PublicKeyEncode(valueToEncrypt, publicKey)
	if err != nil {
		ExitOnError(err)
	}

	fmt.Println(fmt.Sprintf("Encrypted: %s", encryptedValue))
}

func PublicKeyEncode(text string, publicKey string) (string, error) {

	bytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	var decodedKey [32]byte
	copy(decodedKey[:], bytes)

	encrypted, err := box.SealAnonymous(nil, []byte(text), (*[32]byte)(bytes), rand.Reader)
	if err != nil {
		return "", err
	}
	// Encode the encrypted value in base64
	encryptedValue := base64.StdEncoding.EncodeToString(encrypted)

	return encryptedValue, nil
}

func ExitOnError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
