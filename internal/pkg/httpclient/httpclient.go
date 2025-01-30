package httpclient

import (
	"github.com/imroc/req/v3"
)

type HTTPClient struct {
	Client		*req.Client
	PostData	map[string]interface{}
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Client: req.C(),
		PostData: make(map[string]interface{}),
	}
}

func (hc HTTPClient) NewRequest() *req.Request {
	return hc.Client.NewRequest()
}

func (hc HTTPClient) SetBody(body interface{}) *req.Request {
	return hc.SetBody(body)
}

func (hc HTTPClient) Post(url string) (*req.Response, error) {
	return hc.NewRequest().Post(url)
}