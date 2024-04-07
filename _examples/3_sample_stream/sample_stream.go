package main

import (
	"context"
	"fmt"

	"github.com/pefish/go-x-client/tweet/volumestream"
	"github.com/pefish/go-x-client/tweet/volumestream/types"
)

func samplingTweets(c *go_x_client.Client, count int) {
	p := &types.SampleStreamInput{}
	s, err := volumestream.SampleStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		cnt++
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(go_x_client.StringValue(t.Data.ID), go_x_client.StringValue(t.Data.Text))
		}

		if cnt > count {
			s.Stop()
			break
		}
	}
}
