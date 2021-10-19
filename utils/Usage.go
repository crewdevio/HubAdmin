package utils

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Println("Usage: ")
	fmt.Println("Add <username> <access_token> : to add a new account")
	fmt.Println("use <username> : to change account")
	fmt.Println("delete <username> : to delete account")
	os.Exit(1)
}
