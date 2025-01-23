package main

import (
	"fmt"

	"github.com/baisystems/go-test/internal/app"
)

func main() {

    if err := app.RunUser(); err != nil {
        fmt.Println(err.Error())
    }

    if err := app.RunAPI(); err != nil {
        fmt.Println(err.Error())
    }

}
