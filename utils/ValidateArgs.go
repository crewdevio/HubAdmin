package utils

import (
	"os"
)

func ValidateArgs() {
	if len(os.Args) > 1 {
		switch os.Args[1] {

		case "add":
			if len(os.Args) < 3 {
				PrintUsage()
				os.Exit(1)
			}
		case "use":
			if len(os.Args) < 2 {
				PrintUsage()
				os.Exit(1)
			}
		case "delete":
			if len(os.Args) < 2 {
				PrintUsage()
				os.Exit(1)
			}
		case "list":

		case "help":
			PrintUsage()

		default:
			PrintUsage()
			os.Exit(1)
		}
	} else {
		PrintUsage()
	}

}
