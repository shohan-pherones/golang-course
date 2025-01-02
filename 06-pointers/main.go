package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	myNumber := 23

	var ptr *int = &myNumber // reference

	fmt.Println("Value of actual pointer is:", ptr)  // memory address
	fmt.Println("Value of actual pointer is:", *ptr) // value

	*ptr = *ptr * 2
	fmt.Println("New value is:", myNumber)
}
