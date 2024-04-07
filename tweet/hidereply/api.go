package hidereply

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/tweet/hidereply/types"
)

const updateEndpoint = "https://api.twitter.com/2/tweets/:id/hidden"

// Hides or unhides a reply to a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func Update(ctx context.Context, c *go_x_client.Client, p *types.UpdateInput) (*types.UpdateOutput, error) {
	res := &types.UpdateOutput{}
	if err := c.CallAPI(ctx, updateEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
