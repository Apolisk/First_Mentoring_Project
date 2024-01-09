package main

import (
	"os"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var length int
	input(&length, "Password length:", "Password length MUST be a number!")

	var password string
	for i := 0; i < length; i++ {
		password += strconv.Itoa(rand.Intn(10))
	}

	fmt.Println(password)
}

func input(v any, m, e string) {
	fmt.Print(m, " ")
	_, err := fmt.Scanln(v)
	if err != nil {
		fmt.Printf("\n%s (%s)\n", e, err)
		os.Exit(1)
	}
}