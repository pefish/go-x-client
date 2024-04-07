package main

import (
	"context"
	"fmt"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/fields"
	"github.com/pefish/go-x-client/user/userlookup"
	"github.com/pefish/go-x-client/user/userlookup/types"
)

func main() {
	in := &go_x_client.NewClientInput{
		AuthenticationMethod: go_x_client.AuthenMethodOAuth1UserContext,
		ApiKey:               "xEZQ8Q9csXEfzT1MNK43JYTu2",
		ApiKeySecret:         "",
		AccessToken:          "836769434437992449-xZP1rfQLNIjrh3BrKxw6QzeS12DKUm8",
		AccessTokenSecret:    "",
	}

	c, err := go_x_client.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.GetMeInput{
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetMe(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:          ", go_x_client.StringValue(u.Data.ID))
	fmt.Println("Name:        ", go_x_client.StringValue(u.Data.Name))
	fmt.Println("Username:    ", go_x_client.StringValue(u.Data.Username))
	fmt.Println("CreatedAt:   ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", go_x_client.StringValue(t.Text))
		}
	}
}
