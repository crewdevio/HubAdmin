package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crewdevio/HubAdmin/models"
	"github.com/crewdevio/HubAdmin/utils"
	"github.com/fatih/color"
)

func Delete() {
	jsonfile, err := os.Open(utils.AccountsPath())

	if err != nil {
		color.Red("does not have any account added")
		os.Exit(1)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	var accounts []models.Account

	json.Unmarshal(bytesValues, &accounts)

	var index int

	for i := range accounts {
		if accounts[i].Username == os.Args[2] {
			index = i
		}
	}

	accounts = utils.Substract(accounts, index)

	file, _ := json.MarshalIndent(accounts, "", " ")
	_ = ioutil.WriteFile(utils.AccountsPath(), file, 0644)

}
