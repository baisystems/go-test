package app

import (
    "fmt"
    "log"

    "github.com/imroc/req/v3"
	"github.com/baisystems/go-test/internal/pkg/model"
)

func RunAPI() error {

	fmt.Printf("\n\n==========> Calling RunAPI() ...\n")

	// Create a new req client
    client := req.C()

    // Make a GET request to the API
    resp, err := client.R().Get("https://jsonplaceholder.typicode.com/users/1")
    if err != nil {
		log.Fatalf("Failed to make request: %v", err)
		return err
    }

    // Check if the request was successful
    if resp.IsSuccessState() {
        // Parse the JSON response into the User struct
        var user model.User
        err := resp.UnmarshalJson(&user)
        if err != nil {
            log.Fatalf("Failed to unmarshal JSON: %v", err)
			return err
        }

        // Print some of the data returned
        fmt.Printf("User ID: %d\n", user.ID)
        fmt.Printf("User Name: %s\n", user.Name)
        fmt.Printf("User Email: %s\n", user.Email)
    } else {
        log.Fatalf("Request failed with status: %s", resp.Status)
		return err
    }

	fmt.Printf("\n\n\n")
	return nil
}