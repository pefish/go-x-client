package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for sampling count.")
		os.Exit(1)
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	client, err := newOAuth2Client()
	if err != nil {
		panic(err)
	}

	samplingTweets(client, count)
}

func newOAuth2Client() (*go_x_client.Client, error) {
	in2 := &go_x_client.NewClientInput{
		AuthenticationMethod: go_x_client.AuthenMethodOAuth2BearerToken,
	}

	return go_x_client.NewClient(in2)
}
