package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - shohan.vercel.app"
	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
