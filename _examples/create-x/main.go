package main

import (
	"context"
	"fmt"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/tweet/managetweet"
	"github.com/pefish/go-x-client/tweet/managetweet/types"
)

func main() {
	in := &go_x_client.NewClientInput{
		AuthenticationMethod: go_x_client.AuthenMethodOAuth1UserContext,
		ApiKey:               "xEZQ8Q9csXEfzT1MNK43JYTu2",
		ApiKeySecret:         "",
		AccessToken:          "",
		AccessTokenSecret:    "",
	}

	c, err := go_x_client.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := "ttest"
	u, err := managetweet.Create(context.Background(), c, &types.CreateInput{
		Text: &text,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:          ", go_x_client.StringValue(u.Data.ID))
	fmt.Println("Text:        ", go_x_client.StringValue(u.Data.Text))
}
