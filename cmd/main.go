package main

import (
	"fmt"

	requesthandler "github.com/lucastomic/syn-auth/internlas/request_handler"
)

func main() {
	fmt.Println("API server raised. Listening at port 8080")
	requesthandler.ListenAndServe()
}
