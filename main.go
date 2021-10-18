package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

type account struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func deleteCurrentCredential() {

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
	}

}

func add() {
	newProfile := account{
		Username:    os.Args[2],
		AccessToken: os.Args[3],
	}

	jsonfile, err := os.Open("accounts.json")

	if err != nil {
		file, _ := os.OpenFile("accounts.json", os.O_CREATE, os.ModePerm)
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(newProfile)
		color.Green("added account")
		os.Exit(1)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	var datas []account

	json.Unmarshal(bytesValues, &datas)

	find := searchAccount()

	if (find == account{}) {
		datas = append(datas, newProfile)
		file, _ := os.OpenFile("accounts.json", os.O_CREATE, os.ModePerm)
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(datas)

	} else {
		color.Red("Account already exists")
	}
}

func searchAccount() account {
	jsonfile, err := os.Open("accounts.json")

	if err != nil {
		panic(err)
	}

	bytesValues, _ := ioutil.ReadAll(jsonfile)

	var accounts []account

	json.Unmarshal(bytesValues, &accounts)

	for i := 0; i < len(accounts); i++ {
		if os.Args[2] == accounts[i].Username {
			return accounts[i]
		}
	}

	return account{}
}

func printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("Add <username> <access_token> : to add a new account")
	fmt.Println("use <username> : to change account")
	fmt.Println("delete <username> : to delete account")
	os.Exit(1)
}

func validateArgs() {

	if len(os.Args) <= 2 {
		printUsage()
	}

}

// func delete() {
// 	jsonfile, err := os.Open("accounts.json")

// 	if err != nil {
// 		color.Red("does not have any account added")
// 		os.Exit(1)
// 	}

// 	bytesValues, _ := ioutil.ReadAll(jsonfile)

// 	var accounts []account

// 	json.Unmarshal(bytesValues, &accounts)

// 	index := 0

// 	for i := range accounts {
// 		if accounts[i].Username == os.Args[2] {
// 			accounts = append(accounts[:i], accounts[:i+1]...)
// 		}
// 	}
// 	accounts = accounts[:index]

// 	file, _ := json.MarshalIndent(accounts, "", " ")
// 	_ = ioutil.WriteFile("accounts.json", file, 0644)

// }

func switchAccount() {
	user := searchAccount()

	if (user == account{}) {
		panic("user not found")
	}

	deleteCurrentCredential()

	userParse := `/user:` + user.Username
	passParse := `/pass:` + user.AccessToken

	command := exec.Command("cmdkey", "/generic:git:https://github.com", userParse, passParse)

	output, err := command.Output()

	if err != nil {
		panic(string(output))
	}

	color.Green(`switch account: ` + user.Username)
}

func run() {
	validateArgs()

	switch os.Args[1] {

	case "add":
		if len(os.Args) > 3 {
			add()
		} else {
			printUsage()
		}

	case "use":
		switchAccount()

	// case "delete":
	// 	delete()

	default:
		printUsage()
	}
}

func main() {
	run()
}
