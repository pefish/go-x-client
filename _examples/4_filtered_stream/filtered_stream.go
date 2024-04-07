package main

import (
	"context"
	"fmt"

	"github.com/pefish/go-x-client/tweet/filteredstream"
	"github.com/pefish/go-x-client/tweet/filteredstream/types"
)

// createSearchStreamRules lists search stream rules.
func listSearchStreamRules() {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.ListRulesInput{}
	res, err := filteredstream.ListRules(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", go_x_client.StringValue(r.ID), go_x_client.StringValue(r.Value), go_x_client.StringValue(r.Tag))
	}
}

func deleteSearchStreamRules(ruleID string) {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.DeleteRulesInput{
		Delete: &types.DeletingRules{
			IDs: []string{
				ruleID,
			},
		},
	}

	res, err := filteredstream.DeleteRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", go_x_client.StringValue(r.ID), go_x_client.StringValue(r.Value), go_x_client.StringValue(r.Tag))
	}
}

// createSearchStreamRules creates a search stream rule.
func createSearchStreamRules(keyword string) {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: go_x_client.String(keyword), Tag: go_x_client.String(keyword)},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", go_x_client.StringValue(r.ID), go_x_client.StringValue(r.Value), go_x_client.StringValue(r.Tag))
	}
}

// execSearchStream call GET /2/tweets/search/stream API
// and outputs up to 10 results.
func execSearchStream() {
	c, err := newGotwiClientWithTimeout(120)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.SearchStreamInput{}
	s, err := filteredstream.SearchStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			if t != nil {
				cnt++
				fmt.Println(go_x_client.StringValue(t.Data.ID), go_x_client.StringValue(t.Data.Text))
			}
		}

		if cnt > 10 {
			s.Stop()
			break
		}
	}
}
