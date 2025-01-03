package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/users"

func main() {
	fmt.Println("Web Requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	defer res.Body.Close()

	databyte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
