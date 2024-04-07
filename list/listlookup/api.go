package lists

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/list/listlookup/types"
)

const (
	getEndpoint       = "https://api.twitter.com/2/lists/:id"
	listOwnedEndpoint = "https://api.twitter.com/2/users/:id/owned_lists"
)

// Returns the details of a specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-lists-id
func Get(ctx context.Context, c *go_x_client.Client, p *types.GetInput) (*types.GetOutput, error) {
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists owned by the specified user.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-users-id-owned_lists
func ListOwned(ctx context.Context, c *go_x_client.Client, p *types.ListOwnedInput) (*types.ListOwnedOutput, error) {
	res := &types.ListOwnedOutput{}
	if err := c.CallAPI(ctx, listOwnedEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
