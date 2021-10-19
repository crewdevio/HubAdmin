package utils

import "os"

func ValidateArgs() {

	if len(os.Args) <= 2 {
		PrintUsage()
	}

}
