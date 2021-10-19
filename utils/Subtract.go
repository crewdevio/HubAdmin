package utils

import (
	"github.com/crewdevio/HubAdmin/models"
)

func Substract(arr []models.Account, index int) []models.Account {

	if index == 0 && len(arr) > 1 {
		return arr[index+1:]

	} else if len(arr)-1 == index && len(arr) > 2 {
		return arr[:index]

	} else {
		return append(arr[:index], arr[index+1:]...)
	}

}
