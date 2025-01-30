package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/baisystems/go-test/internal/pkg/httpclient"
	"github.com/baisystems/go-test/internal/pkg/model"
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

func TestGetUser(t *testing.T) {
    mockClient := new(httpclient.MockClient)
    mockReq := new(req.Request)
    mockResp := &req.Response{
        Response: &http.Response{
            StatusCode: 200,
            Status:     "200 OK",
        },
        Request: mockReq,
    }

    mockResp.SetBody([]byte(`{"id": 1, "name": "John Doe", "email": "johndoe@example.com"}`))

    // Set up expectations
    mockClient.On("NewRequest").Return(mockReq)
    mockClient.On("SetBody", mock.Anything).Return(mockReq)
    mockClient.On("Get", "https://jsonplaceholder.typicode.com/users/1").Return(mockResp, nil)

    resp, err := GetUser(mockClient, 1)

    // Assert the response with no error
    require.NoError(t, err)
    require.NotNil(t, resp)

    bodyBytes := resp.Bytes()
    bodyString := string(bodyBytes)
    fmt.Printf("Response Body: %s\n", bodyString)

    var user model.User
    err = json.Unmarshal([]byte(bodyString), &user)
    if err != nil {
        fmt.Printf("Failed to unmarshal JSON: %v\n", err)
        return
    }

    fmt.Printf("User ID: %d\n", user.ID)
    fmt.Printf("User Name: %s\n", user.Name)
    fmt.Printf("User Email: %s\n", user.Email)

    require.Equal(t, int64(1), user.ID)
    require.Equal(t, "John Doe", user.Name)
    require.Equal(t, "johndoe@example.com", user.Email)
}