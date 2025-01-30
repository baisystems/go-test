package httpclient

import (
	"github.com/imroc/req/v3"
)

type HTTPClient struct {
	Client		*req.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Client: req.C(),
	}
}

func (c HTTPClient) NewRequest() *req.Request {
	return c.Client.NewRequest()
}

func (c HTTPClient) SetBody(body interface{}) *req.Request {
	return c.SetBody(body)
}

func (c HTTPClient) Post(url string) (*req.Response, error) {
	return c.NewRequest().Post(url)
}

func (c HTTPClient) Get(url string) (*req.Response, error) {
	return c.NewRequest().Get(url)
}