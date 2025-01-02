package main

import "fmt"

func main() {
	shohan := User{"shohan", "shohan@go.dev", true, 20}
	fmt.Println(shohan)
	fmt.Printf("%+v\n", shohan)
	fmt.Println(shohan.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
