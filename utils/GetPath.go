package utils

import (
	"log"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
)

func AccountsPath() string {
	usr, errorGetPath := user.Current()

	if errorGetPath != nil {
		color.Red("getting user home directory: ")
		log.Fatal(errorGetPath)
	}

	return filepath.Join(usr.HomeDir, "huba.json")
}
