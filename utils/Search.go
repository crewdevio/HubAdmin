package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crewdevio/HubAdmin/models"
	"github.com/fatih/color"
)

func SearchAccount() models.Account {
	jsonfile, err := os.Open(AccountsPath())

	if err != nil {
		color.Red("file accouts.json dont exists")
		os.Exit(1)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	var accounts []models.Account

	json.Unmarshal(bytesValues, &accounts)

	for i := 0; i < len(accounts); i++ {
		if os.Args[2] == accounts[i].Username {
			return accounts[i]
		}
	}

	return models.Account{}
}
