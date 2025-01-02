package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruits = []string{"Apple", "Banana", "Grapes"}
	fmt.Printf("Type of fruits is: %T\n", fruits)

	fruits = append(fruits, "Tomato", "Mango")
	fmt.Println(fruits)

	fruits = append(fruits[:3])
	fmt.Println(fruits)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 934
	highscores[2] = 434
	highscores[3] = 834
	// highscores[4] = 777

	highscores = append(highscores, 555, 666, 777)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)

	// how to remove a value from slices based on index
	var courses = []string{"react.js", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)
}
