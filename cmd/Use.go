package cmd

import (
	"os/exec"

	"github.com/crewdevio/HubAdmin/models"
	"github.com/crewdevio/HubAdmin/utils"
	"github.com/fatih/color"
)

func SwitchAccount() {
	user := utils.SearchAccount()

	if (user == models.Account{}) {
		panic("user not found")
	}

	utils.DeleteCurrentCredential()

	userParse := `/user:` + user.Username
	passParse := `/pass:` + user.AccessToken

	command := exec.Command("cmdkey", "/generic:git:https://github.com", userParse, passParse)

	output, err := command.Output()

	if err != nil {
		panic(string(output))
	}

	color.Green(`switch account: ` + user.Username)
}
