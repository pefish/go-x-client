package main

import (
	"fmt"
	"os"

	go_x_client "github.com/pefish/go-x-client"
)

const (
	OAuthTokenEnvKeyName       = "X_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "X_ACCESS_TOKEN_SECRET"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for your account id.")
		os.Exit(1)
	}

	accountID := args[1]

	oauth1Client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	onlyFollowsRecentActivity(oauth1Client, accountID)
}

func newOAuth1Client() (*go_x_client.Client, error) {
	in := &go_x_client.NewClientInput{
		AuthenticationMethod: go_x_client.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return go_x_client.NewClient(in)
}
