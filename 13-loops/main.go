package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("%v -> %v\n", index, day)
	// }

	val := 1

	for val < 10 {

		if val == 2 {
			goto don
		}

		if val == 5 {
			val++
			continue
		}

		fmt.Println(val)
		val++
	}

don:
	fmt.Println("don")
}
