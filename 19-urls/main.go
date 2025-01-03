package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://shohan.dev:3000/learn?coursename=reactjs&paymentid=ght5775"

func main() {
	fmt.Println("URLs in Golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	for key, val := range qparams {
		fmt.Printf("%v --> %v\n", key, val)
	}

	// construction
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "shohan.dev",
		Path:     "/learn",
		RawQuery: "user=shohan",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
