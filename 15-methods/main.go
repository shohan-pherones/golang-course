package main

import "fmt"

func main() {
	shohan := User{"shohan", "shohan@go.dev", true, 20}
	shohan.getStatus()
	shohan.NewMail()
	fmt.Println(shohan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user:", u.Email)
}
