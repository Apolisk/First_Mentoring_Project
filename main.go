package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// Generation config
	var (
		length    uint
		uppercase bool // TODO: Make flag "letters" instead
	)

	input(&length, "Password length:", "Password length must be a positive number!")
	input(&uppercase, "Uppercase? (y/n):", "")

	password := generate(length, uppercase)
	fmt.Println(password)
}

func input(v any, msg, emsg string) {
	fmt.Print(msg, " ")

	if b, ok := v.(*bool); ok {
		var s string
		fmt.Scanln(&s)
		*b = s == "y"
		return
	}

	_, err := fmt.Scanln(v)
	if err != nil {
		fmt.Printf("\n%s (%s)\n", emsg, err)
		os.Exit(1)
	}
}

func generate(length uint, uppercase bool) (password string) {
	if length == 0 {
		return "Password length cannot be zero"
	}

	if uppercase {
		const upperLetters = "ABCDEFGHIKLMNOPQRSTVXYZ"

		for i := 0; i < int(length); i++ {
			if chance(50) {
				password += strconv.Itoa(rand.Intn(10))
			} else {
				password += string(upperLetters[rand.Intn(len(upperLetters))])
			}
		}

		return password
	} 

	for i := 0; i < int(length); i++ {
		password += strconv.Itoa(rand.Intn(10))
	}

	return password
}

func chance(x int) bool {
	return rand.Intn(100) < x - 1
}
