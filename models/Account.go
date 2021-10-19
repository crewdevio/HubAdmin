package models

type Accounts struct {
	Accounts []Account
}

type Account struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Active      bool   `json:"active"`
}
