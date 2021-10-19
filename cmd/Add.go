package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crewdevio/HubAdmin/models"
	"github.com/crewdevio/HubAdmin/utils"
	"github.com/fatih/color"
)

func Add() {
	newProfile := models.Account{
		Username:    os.Args[2],
		AccessToken: os.Args[3],
	}

	jsonfile, err := os.Open(utils.AccountsPath())

	if err != nil {
		file, _ := os.OpenFile(utils.AccountsPath(), os.O_CREATE, os.ModePerm)
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(newProfile)
		color.Green("added account")
		os.Exit(1)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	var datas []models.Account

	json.Unmarshal(bytesValues, &datas)

	find := utils.SearchAccount()

	if (find == models.Account{}) {
		datas = append(datas, newProfile)
		file, _ := os.OpenFile(utils.AccountsPath(), os.O_CREATE, os.ModePerm)
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(datas)
		color.Green("added account: " + os.Args[2])
	} else {
		color.Red("Account already exists")
	}
}
