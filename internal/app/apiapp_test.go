package app

import (
	"fmt"
	"testing"

	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/baisystems/go-test/internal/pkg/httpclient"
	"github.com/h2non/gock"
	"github.com/imroc/req/v3"
)

// MockClient is a mock implementation of the HTTPClient interface
type MockClient struct{}

func (m *MockClient) R() *req.Request {
    return req.NewRequest()
}

func CreatePost(client httpclient.RealClient, newPost model.Post) (*model.PostResponse, error) {
    resp, err := client.
        SetBody(newPost).
        Post("https://jsonplaceholder.typicode.com/posts")
    if err != nil {
        return nil, err
    }

    if resp.IsSuccess() {
        var response model.PostResponse
        err := resp.UnmarshalJson(&response)
        if err != nil {
            return nil, err
        }
        return &response, nil
    }
    return nil, fmt.Errorf("request failed with status: %s", resp.Status)
}

func TestCreatePost(t *testing.T) {
    defer gock.Off()

    // Mock the HTTP request and response
    gock.New("https://jsonplaceholder.typicode.com").
        Post("/posts").
        Reply(201).
        JSON(map[string]interface{}{
            "id":     101,
            "title":  "foo",
            "body":   "bar",
            "userId": 1,
        })

    // Create a new mock client
    mockClient := &MockClient{}

    // Call the CreatePost function with the mock client
    response, err := CreatePost(mockClient, Post{
        Title:  "foo",
        Body:   "bar",
        UserID: 1,
    })

    // Assert the response and error
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if response == nil {
        t.Fatalf("Expected response, got nil")
    }
    if response.ID != 101 {
        t.Errorf("Expected ID 101, got %d", response.ID)
    }
    if response.Title != "foo" {
        t.Errorf("Expected title 'foo', got %s", response.Title)
    }
    if response.Body != "bar" {
        t.Errorf("Expected body 'bar', got %s", response.Body)
    }
    if response.UserID != 1 {
        t.Errorf("Expected userID 1, got %d", response.UserID)
    }

    // Verify that all mocks were called
    if !gock.IsDone() {
        t.Errorf("Expected all mocks to be called")
    }
}
