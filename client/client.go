package client

// Client struct for bitbucket http client
type Client struct {
	username string
	password string
}

// New creates a new bitbucket http client
func New(username string, password string) *Client {
	return &Client{username, password}
}

func (c *Client) call(url string) {

}
