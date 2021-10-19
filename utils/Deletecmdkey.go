package utils

import (
	"os/exec"
	"strings"
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
	}

}
