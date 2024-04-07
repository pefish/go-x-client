package main

import (
	"context"

	"github.com/pefish/go-x-client/tweet/managetweet"
	"github.com/pefish/go-x-client/tweet/managetweet/types"
)

// SimpleTweet posts a tweet with only text, and return posted tweet ID.
func SimpleTweet(c *go_x_client.Client, text string) (string, error) {
	p := &types.CreateInput{
		Text: go_x_client.String(text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return go_x_client.StringValue(res.Data.ID), nil
}

// DeleteTweet deletes a tweet specified by tweet ID.
func DeleteTweet(c *go_x_client.Client, id string) (bool, error) {
	p := &types.DeleteInput{
		ID: id,
	}

	res, err := managetweet.Delete(context.Background(), c, p)
	if err != nil {
		return false, err
	}

	return go_x_client.BoolValue(res.Data.Deleted), nil
}
