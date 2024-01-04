package main

import "fmt"

func main() {
	var length int

	fmt.Print("Password length: ")
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Println("Password length MUST be a number!")
		return
	}

	fmt.Println(length)
}