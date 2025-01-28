package httpclient

import (
	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/imroc/req/v3"
)

type MockClient struct {
	Client		*req.Client
	PostData	map[string]interface{}
}

func NewClient() *MockClient {
	return &MockClient{
		Client: req.C(),
		PostData: make(map[string]interface{}),
	}
}

func (mc MockClient) NewRequest() *req.Request {
	return mc.Client.NewRequest()
}

func (mc MockClient) Post() (resp model.PostResponse, err error) {
	
}