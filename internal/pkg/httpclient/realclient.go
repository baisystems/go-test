package httpclient

import (
	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/imroc/req/v3"
)

type RealClient struct {
	Client		*req.Client
	PostData	map[string]interface{}
}

func NewClient() *RealClient {
	return &RealClient{
		Client: req.C(),
		PostData: make(map[string]interface{}),
	}
}

func (rc RealClient) NewRequest() *req.Request {
	return rc.Client.NewRequest()
}

func (rc RealClient) Post() (resp model.PostResponse, err error) {
	
}