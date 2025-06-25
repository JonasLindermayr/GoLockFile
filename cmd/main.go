package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JonasLindermayr/GoLockFile/utils"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: use 'encrypt', 'decrypt' or 'help'")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "encrypt":
		encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
		filePath := encryptCmd.String("file", "", "path to file")
		password := encryptCmd.String("p", "", "password")
		name := encryptCmd.String("n", "", "new filename")
		overwrite := encryptCmd.Bool("w", false, "delete original file after encryption")

		encryptCmd.Parse(os.Args[2:])

		if *filePath == "" || *password == "" {
			fmt.Println("Error: usage of encrypt:")
			encryptCmd.PrintDefaults()
			os.Exit(1)
		}

		err := utils.EncryptFile(*filePath, *name, *password, *overwrite)

		if err != nil {
			fmt.Println("Error while encrypting:", err)
			os.Exit(1)
		}
		fmt.Println("Successfull encrypted")

	case "decrypt":
		decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
		filePath := decryptCmd.String("file", "", "path to file")
		password := decryptCmd.String("p", "", "password")
		overwrite := decryptCmd.Bool("w", false, "delete encrypted file after decryption")

		decryptCmd.Parse(os.Args[2:])

		if *filePath == "" || *password == "" {
			fmt.Println("Error: usage of decrypt:")
			decryptCmd.PrintDefaults()
			os.Exit(1)
		}

		err := utils.DecryptFile(*filePath, *password, *overwrite)
		if err != nil {
			fmt.Println("Error while decrypting:", err)
			os.Exit(1)
		}
		fmt.Println("Successfull decrypted")

	case "help":
		fmt.Println("Usage:")
		fmt.Println(">  golockfile encrypt -file <filepath> -p <password>")
		fmt.Println(">  golockfile encrypt -file <filepath> -p <password> -n <filename>")
		fmt.Println(">  golockfile encrypt -file <filepath> -p <password> -w // deletes original file")
		fmt.Println(">  golockfile decrypt -file <filepath> -p <password>")
		fmt.Println(">  golockfile decrypt -file <filepath> -p <password> -w // deletes encrypted file")
		fmt.Println(">  golockfile help                             ")

	default:
		fmt.Println("Error: unknown subcommand:", os.Args[1])
		os.Exit(1)
	}

}
