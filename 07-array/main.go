package main

import "fmt"

func main() {
	fmt.Println("My Arrays")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	fmt.Println("My fruits array:", fruits)
	fmt.Println("My fruits array:", len(fruits))

	var vegetables = [5]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegetables array:", vegetables)
	fmt.Println("Vegetables array:", len(vegetables))
}
