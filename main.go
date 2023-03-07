package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func main() {
	var passwordLength int

	fmt.Println("welcome to the password generator")
	time.Sleep(2 * time.Second)
	fmt.Println("enter a length for password (min :8 )")
	fmt.Scanln(&passwordLength)
	for passwordLength < 8 {
		fmt.Println("Invalid password length specified, minimum length is 8")
		fmt.Scanln(&passwordLength)

	}

	passwordBytes := make([]byte, passwordLength)
	_, err := rand.Read(passwordBytes)
	if err != nil {
		fmt.Println("Error generating random bytes")
		os.Exit(1)
	}

	password := base64.URLEncoding.EncodeToString(passwordBytes)
	fmt.Println(password[:passwordLength])
}
