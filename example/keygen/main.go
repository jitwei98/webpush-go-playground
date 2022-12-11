package main

import (
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
)

func main() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("private key: %s\n", privateKey)
	fmt.Printf("public key: %s\n", publicKey)
}
