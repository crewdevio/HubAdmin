package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/crewdevio/HubAdmin/models"
	"github.com/fatih/color"
)

func DeleteCurrentCredential() {

	list := exec.Command("cmdkey", "/list:git:https://github.com")

	out, err := list.Output()

	if err != nil {
		panic(err)
	}
	output := string(out)

	if !strings.Contains(output, "*") {
		delete := exec.Command("cmdkey", "/delete:git:https://github.com")

		_, err := delete.Output()

		if err != nil {
			panic(err)
		}

		jsonfile, err := os.Open(AccountsPath())

		if err != nil {
			color.Red("does not have any account added")
			os.Exit(1)
		}

		bytesValues, _ := ioutil.ReadAll(jsonfile)

		var accounts []models.Account

		json.Unmarshal(bytesValues, &accounts)

		var index int

		fmt.Println(index)
		for i := range accounts {
			if accounts[i].Active {
				index = i
			}
		}

		accounts[index].Active = false

		file, _ := json.MarshalIndent(accounts, "", " ")
		_ = ioutil.WriteFile(AccountsPath(), file, 0644)
	}

}
