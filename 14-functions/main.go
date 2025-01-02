package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, proMessage := proAdder(3, 5, 5, 4, 3, 2, 7)
	fmt.Printf("Pro message: %v and Pro result is: %v\n", proMessage, proResult)
}

func greeter() {
	fmt.Println("Hello from golang")
}

func adder(valOne, valTwo int) int {
	return valOne + valTwo
}

func proAdder(vals ...int) (int, string) {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total, "Hi Pro result func"
}
