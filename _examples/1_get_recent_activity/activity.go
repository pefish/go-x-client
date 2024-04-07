package main

import (
	"context"
	"fmt"

	"github.com/pefish/go-x-client/fields"
	"github.com/pefish/go-x-client/tweet/searchtweet"
	sttypes "github.com/pefish/go-x-client/tweet/searchtweet/types"
	"github.com/pefish/go-x-client/user/follow"
	ftypes "github.com/pefish/go-x-client/user/follow/types"
)

type twitterUser struct {
	ID       string
	Name     string
	Username string
}

func (f twitterUser) displayName() string {
	return fmt.Sprintf("%s@%s", f.Name, f.Username)
}

// onlyFollowsRecentActivity will output the accounts that are unilaterally following
// the specified user ID, along with up to three most recent tweets.
func onlyFollowsRecentActivity(c *go_x_client.Client, userID string) {
	// list follows
	followings := map[string]twitterUser{}

	paginationToken := "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowingsInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowings(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followings[go_x_client.StringValue(u.ID)] = twitterUser{
				ID:       go_x_client.StringValue(u.ID),
				Name:     go_x_client.StringValue(u.Name),
				Username: go_x_client.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = go_x_client.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// list followers
	followers := map[string]twitterUser{}

	paginationToken = "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowersInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowers(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followers[go_x_client.StringValue(u.ID)] = twitterUser{
				ID:       go_x_client.StringValue(u.ID),
				Name:     go_x_client.StringValue(u.Name),
				Username: go_x_client.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = go_x_client.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// only following
	onlyFollowings := map[string]twitterUser{}
	for fid, u := range followings {
		if _, ok := followers[fid]; ok {
			continue
		}

		onlyFollowings[fid] = u
	}

	// get recent tweets
	for _, onlyFollow := range onlyFollowings {
		p := &sttypes.ListRecentInput{
			MaxResults:  10,
			Query:       "from:" + onlyFollow.Username + " -is:retweet -is:reply",
			TweetFields: fields.TweetFieldList{fields.TweetFieldCreatedAt},
		}
		res, err := searchtweet.ListRecent(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("----- %s's recent Tweets -----\n", onlyFollow.displayName())
		c := 0
		for _, t := range res.Data {
			if c > 3 {
				break
			}
			fmt.Printf("[%s] %s\n", t.CreatedAt, go_x_client.StringValue(t.Text))
			c++
		}

		fmt.Println()
	}
}
