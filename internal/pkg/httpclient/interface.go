package httpclient

import (
	"github.com/imroc/req/v3"
)

type HTTPClientInterface interface {
    NewRequest() *req.Request
	SetBody(body interface{}) *req.Request
	Post(url string) (*req.Response, error)
	// Get(url string) (*req.Response, error)
}