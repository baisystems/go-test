package app

import (
	"fmt"
	"log"

    "github.com/imroc/req/v3"
	"github.com/baisystems/go-test/internal/pkg/httpclient"
	"github.com/baisystems/go-test/internal/pkg/model"
)

func RunAPI() error {

    if err := DoGet(); err != nil {
        fmt.Println(err.Error())
    }

	if err := DoPost(); err != nil {
        fmt.Println(err.Error())
    }

	fmt.Printf("\n\n\n")
	return nil
}

func DoPost() error {
	fmt.Printf("\n\n==========> Calling DoPost() ...\n")

    client := httpclient.NewHTTPClient()

    newPost := model.Post{
        Title:  "My 1st post to an endpont",
        Body:   "This is the centent of the post. You can't retrieve the post afterwards because it's sent to a fake endpoint.",
        UserID: 1,
    }

    client.NewRequest().SetBody(newPost)
    if err := CreatePost(client); err != nil {
        log.Fatalf("Failed to create post: %v", err)
        return err
    }

	return nil
}

func CreatePost(client httpclient.HTTPClientInterface) error {

    resp, err := client.Post("https://jsonplaceholder.typicode.com/posts")

    if err != nil {
        return err
    }

    if resp.Response.StatusCode == 201 {
        fmt.Println("real or mock http post request succeed")
        return nil
    }
    return fmt.Errorf("request failed with status: %s", resp.Status)
}

func DoGet() error {
	fmt.Printf("\n\n==========> Calling DoGet() ...\n")

    client := httpclient.NewHTTPClient()
    client.NewRequest()
    
    resp, err := GetUser(client, 1); 
    if err != nil {
        log.Fatalf("Failed calling GetUser func: %v", err)
		return err
    }

    var user model.User
    err = resp.UnmarshalJson(&user)
    if err != nil {
        log.Fatalf("Failed to unmarshal JSON: %v", err)
        return err
    }

    fmt.Printf("User ID: %d\n", user.ID)
    fmt.Printf("User Name: %s\n", user.Name)
    fmt.Printf("User Email: %s\n", user.Email)
	return nil
}

func GetUser(client httpclient.HTTPClientInterface, userId int) (*req.Response, error) {
    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userId)
    return client.Get(url)
}
