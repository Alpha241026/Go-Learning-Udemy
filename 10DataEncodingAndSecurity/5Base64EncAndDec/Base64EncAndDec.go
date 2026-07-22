package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to the wonderful world of Go !"

	encoded := base64.StdEncoding.EncodeToString([]byte(data)) //encode a string into its Base64 representation
	fmt.Println(encoded)

	encodedStr := "V2VsY29tZSB0byB0aGUgd29uZGVyZnVsIHdvcmxkIG9mIEdvIQ=="

	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr) //decode a Base64 string back into its original bytes
	if err != nil {
		log.Fatal(err)
	}

	if string(decodedStr) != data {
		log.Fatal("decoded string does not match encoded data")
	}

	rawData := []byte{0xDE, 0xAD, 0xEF, 0xCA, 0xFE}
	binaryCodedToString := base64.StdEncoding.EncodeToString(rawData) //encode arbitrary binary data as printable text

	fmt.Println(string(binaryCodedToString))

	b64Str := "3q3vyv4="
	decodedStr, err = base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.Fatal(err)
	}

}
