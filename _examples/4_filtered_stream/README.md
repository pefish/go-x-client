filtered stream example
===

This is sample code that uses the `filteredstream` package.

# Steps to use Filtered Stream API

## 1. Create search stream rules.

Suppose you want to retrieve streams filtered by the keyword "Twitter API v2", create a search stream rule with the following code.

```go
func CreateSearchStreanmRules(c *go_x_client.Client) {
	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: go_x_client.String("Twitter API v2"), Tag: go_x_client.String("example rule")},
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
```

## 2. Search stream

To retrieve the streams to which you have applied the rules (filters) you have created, implement as follows.

```go
func SearchStream(c *go_x_client.Client) {
	p := &types.SearchStreamInput{}
	s, err := filteredstream.SearchStream(context.Background(), c, p)
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

		if cnt > 10 {
			s.Stop()
			break
		}
	}
}
```

# Tips

If you want to adjust the time to run the stream, pass `http.Client` with an arbitrary timeout as input when generating `go_x_client.Client` as shown below. (default is 30 seconds)

For example 120 sec.

```go
func main() {
	in := &go_x_client.NewClientInput{
		AuthenticationMethod: go_x_client.AuthenMethodOAuth2BearerToken,
		HTTPClient: &http.Client{
			Timeout: time.Duration(120) * time.Second,
		},
	}

	gotwiClient, err := go_x_client.NewClient(in)
}
```

If you do not want a timeout, set `Timeout: 0`. See the `http` package implementation below.

[http package - net/http - Go Packages](https://pkg.go.dev/net/http#Client.Timeout)

# Run example code

1. Set environment variables.

    ```bash
    export X_API_KEY=your-api-key
    export X_API_KEY_SECRET=your-api-key-secret
    ```
    
2. Create a search stream rule.

    ```bash
    go run . create 'Twitter'
    ```
    
3. Call filtered stream API.

    ```bash
    go run . stream
    ```
    
4. List search stream rules.

    ```bash
    go run . list
    ```
    
5. Delete a search stream rule.

    ```bash
    go run . delete rule-id
    ```