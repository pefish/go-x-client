package quotetweet

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/tweet/quotetweet/types"
)

const (
	listEndpoint = "https://api.twitter.com/2/tweets/:id/quote_tweets"
)

// Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func List(ctx context.Context, c *go_x_client.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
