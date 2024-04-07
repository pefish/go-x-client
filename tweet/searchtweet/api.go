package searchtweet

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/tweet/searchtweet/types"
)

const (
	listRecentEndpoint = "https://api.twitter.com/2/tweets/search/recent"
	listAllEndpoint    = "https://api.twitter.com/2/tweets/search/all"
)

// The recent search endpoint returns Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func ListRecent(ctx context.Context, c *go_x_client.Client, p *types.ListRecentInput) (*types.ListRecentOutput, error) {
	res := &types.ListRecentOutput{}
	if err := c.CallAPI(ctx, listRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func ListAll(ctx context.Context, c *go_x_client.Client, p *types.ListAllInput) (*types.ListAllOutput, error) {
	res := &types.ListAllOutput{}
	if err := c.CallAPI(ctx, listAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
