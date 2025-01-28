package httpclient

import (
	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/imroc/req/v3"
)

type HTTPClient interface {
    NewRequest() *req.Request
	SetBody() error
	Post(url string) (resp model.PostResponse, err error)
	Get() (resp model.PostResponse, err error)
}