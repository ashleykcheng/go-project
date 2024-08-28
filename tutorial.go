package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("Please input your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!", name)

	fmt.Printf("Please enter your age: ")
	var age uint
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You cannot play")
		return
	}

}
