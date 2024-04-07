package volumestream

import (
	"context"

	go_x_client "github.com/pefish/go-x-client"
	"github.com/pefish/go-x-client/tweet/volumestream/types"
)

const (
	sampleStreamEndpoint = "https://api.twitter.com/2/tweets/sample/stream"
)

// Streams about 1% of all Tweets in real-time.
// If you have Academic Research access, you can connect up to two redundant connections to maximize your streaming up-time.
// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/api-reference/get-tweets-sample-stream
func SampleStream(ctx context.Context, c *go_x_client.Client, p *types.SampleStreamInput) (*go_x_client.StreamClient[*types.SampleStreamOutput], error) {
	tc := go_x_client.NewTypedClient[*types.SampleStreamOutput](c)
	s, err := tc.CallStreamAPI(ctx, sampleStreamEndpoint, "GET", p)
	if err != nil {
		return nil, err
	}

	return s, nil
}
