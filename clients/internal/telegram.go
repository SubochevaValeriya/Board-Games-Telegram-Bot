package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

//func (c *Client) Updates(offset int, limit int) ([]Update, error) {
//	q := url.Values{}
//	q.Add("offset")
//}

func (c *Client) SendMessage() {

}
