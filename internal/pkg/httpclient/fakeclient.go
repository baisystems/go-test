package httpclient

import (
	// "github.com/baisystems/go-test/internal/pkg/model"
	"github.com/imroc/req/v3"
	// "github.com/h2non/gock"
)

type MockClient struct {
	Client		*req.Client
	PostData	map[string]interface{}
}

func NewMockClient() *MockClient {
	return &MockClient{
		Client: req.C(),
		PostData: make(map[string]interface{}),
	}
}

func (mc MockClient) NewRequest() *req.Request {
	return mc.Client.NewRequest()
}

func (mc MockClient) SetBody(body interface{}) *req.Request {
	return mc.Client.NewRequest()
}

func (mc MockClient) Post(url string) (interface{}, error) {
	return mc.NewRequest().Post(url)
}