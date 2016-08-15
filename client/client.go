package main

// Client struct for bitbucket http client
type Client struct {
	username string
	password string
}

// NewClient creates a new bitbucket http client
func NewClient(username string, password string) *Client {
	return &Client{username, password}
}

func (c *Client) call(url string) {

}
