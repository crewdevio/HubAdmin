package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/crewdevio/HubAdmin/utils"
	"github.com/fatih/color"
)

func List() {
	jsonfile, err := os.Open(utils.AccountsPath())

	if err != nil {
		color.Red("file accouts.json dont exists")
		os.Exit(1)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	type account struct {
		Username string `json:"username"`
		Active   bool   `json:"active"`
	}

	var accounts []account

	json.Unmarshal(bytesValues, &accounts)

	fmt.Println()
	for _, i := range accounts {
		color.Yellow("Username: " + i.Username)
		if i.Active {
			color.Green("Active: " + strconv.FormatBool(i.Active))
		} else {
			color.Red("Active: " + strconv.FormatBool(i.Active))
		}

		fmt.Println()
	}
}
