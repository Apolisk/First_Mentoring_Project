package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var length int

	fmt.Print("Password length : ")
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Println("Password length MUST be a number!")
		return
	}
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	password := ""
	for i := 0; i < length; i++ {
		randVal := numbers[rand.Intn(len(numbers))]
		password += strconv.Itoa(randVal)
	}
	fmt.Println("Yes")
	fmt.Println(password)
}
