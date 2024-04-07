package block

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/user/block/types"
)

const (
	listEndpoint   = "https://api.twitter.com/2/users/:id/blocking"
	createEndpoint = "https://api.twitter.com/2/users/:id/blocking"
	deleteEndpoint = "https://api.twitter.com/2/users/:source_user_id/blocking/:target_user_id"
)

// Returns a list of users who are blocked by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/get-users-blocking
func List(ctx context.Context, c *go_x_client.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user (in the path) to block the target user. The user (in the path) must match the user context authorizing the request.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/post-users-user_id-blocking
func Create(ctx context.Context, c *go_x_client.Client, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to unblock another user.
// The request succeeds with no action when the user sends a request to a user they're not blocking or have already unblocked.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/delete-users-user_id-blocking
func Delete(ctx context.Context, c *go_x_client.Client, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
