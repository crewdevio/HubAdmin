package cmd

import (
	"os"

	"github.com/crewdevio/HubAdmin/utils"
)

func Run() {
	utils.ValidateArgs()

	switch os.Args[1] {

	case "add":
		if len(os.Args) > 3 {
			Add()
		} else {
			utils.PrintUsage()
		}

	case "use":
		SwitchAccount()

	case "delete":
		Delete()

	default:
		utils.PrintUsage()
	}
}
