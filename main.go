package main

import (
	"./encoder"
	"flag"
	"fmt"
	"os"
)

func main() {
	newEncoder := encoder.NewEncoder([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g',
		'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
		[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
			'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
			'V', 'W', 'X', 'Y', 'Z'})

	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptText := encryptCmd.String("t", "", "t")
	encryptKey := encryptCmd.Int("k", 0, "k")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptText := decryptCmd.String("t", "", "t")
	decryptKey := decryptCmd.Int("k", 0, "k")

	if len(os.Args) < 2 {
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "encrypt":
		err := encryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
		result, err := newEncoder.Encrypt(*encryptText, *encryptKey)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	case "decrypt":
		err := decryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
		result, err := newEncoder.Decrypt(*decryptText, *decryptKey)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	default:
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

}
