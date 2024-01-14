package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Generation config
	var (
		length    uint
		uppercase bool     // TODO: Make flag "letters" instead
		quantity  int      // Quantity of passwords
		passwords []string // Slice to save all generated passwords
	)

	input(&length, "Password length:", "Password length must be a positive number!")
	input(&uppercase, "Uppercase? (y/n):", "")
	input(&quantity, "How much passwords you want?: ", "Password quantity must be a positive number!")

	for i := 0; i < quantity; i++ {
		passwords = append(passwords, generate(length, uppercase))
	}

	write(passwords)
}

func input(v any, msg, emsg string) {
	fmt.Print(msg, " ")

	if b, ok := v.(*bool); ok {
		var s string
		_, _ = fmt.Scanln(&s)
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
		const Letters = "ABCDEFGHIKLMNOPQRSTVXYZabcdefghijklmnopqrstuvwxyz"

		for i := 0; i < int(length); i++ {
			if chance(50) {
				password += strconv.Itoa(rand.Intn(10))
			} else {
				password += string(Letters[rand.Intn(len(Letters))])
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
	return rand.Intn(100) < x-1
}

func write(passwords []string) {
	data := strings.Join(passwords, "\n")
	err := os.WriteFile("output.txt", []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
