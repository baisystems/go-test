package httpclient

import (
    "github.com/stretchr/testify/mock"
    "github.com/imroc/req/v3"
)

type MockClient struct {
    mock.Mock
}

func (m *MockClient) NewRequest() *req.Request {
    args := m.Called()
    return args.Get(0).(*req.Request)
}

func (m *MockClient) SetBody(body interface{}) *req.Request {
    args := m.Called(body)
    req, ok := args.Get(0).(*req.Request)
    if !ok {
        panic("type assertion to *req.Request failed")
    }
    return req
}

func (m *MockClient) Post(url string) (*req.Response, error) {
    args := m.Called(url)
    response, ok := args.Get(0).(*req.Response)
    if !ok {
        panic("type assertion to *req.Response failed")
    }
    return response, args.Error(1)
}

func (m *MockClient) Get(url string) (*req.Response, error) {
    args := m.Called(url)
    response, ok := args.Get(0).(*req.Response)
    if !ok {
        panic("type assertion to *req.Response failed")
    }
    return response, args.Error(1)
}