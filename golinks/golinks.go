package golinks

const (
	baseURL          = "https://api.golinks.io"
	defaultUserAgent = "go-links"
)

// API Doc:
// https://docs.golinks.io/#19f87188-f2fb-4b75-bf69-83f0c0ca5029
type service struct {
	client *Client
}
