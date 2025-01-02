package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// loops
	for key, value := range languages {
		fmt.Printf("%v -> %v\n", key, value)
	}
}
