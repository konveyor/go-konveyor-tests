package auth

import (
	"github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
)

type LoginObject struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Login performs a login request to the hub and returns a token
func Login(client *addon.Client, username, password string) (string, error) {
	login := LoginObject{User: username, Password: password}
	err := client.Post(api.AuthLoginRoot, &login)
	if err != nil {
		return "", err
	}
	return login.Token, nil
}
