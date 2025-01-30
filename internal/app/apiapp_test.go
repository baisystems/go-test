package app

import (
	"net/http"
	"testing"

	"github.com/baisystems/go-test/internal/pkg/httpclient"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
    mockClient := new(httpclient.MockClient)
    mockReq := new(req.Request)
    mockResp := &req.Response{
        Response: &http.Response{
            StatusCode: 201,
            Status:     "201 CREATED",
        },
    }

    // Set up expectations
    mockClient.On("NewRequest").Return(mockReq)
    mockClient.On("SetBody", mock.Anything).Return(mockReq)
    mockClient.On("Post", "https://jsonplaceholder.typicode.com/posts").Return(mockResp, nil)
    
    err := CreatePost(mockClient)

    // Assert the response with no error
    require.NoError(t, err)
}
