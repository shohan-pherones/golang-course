package main

import "fmt"

const LoginToken string = "token" // public var

func main() {
	var username string = "shohan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float64 = 255.25444111215155551512111
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "shohan.vercel.app"
	fmt.Println(website)

	// no var style
	numberOfUser := 20.25
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
