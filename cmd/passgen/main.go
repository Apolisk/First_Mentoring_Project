package main

import (
	"fmt"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Generation config
	var (
		length    uint
		letters   bool
		count     int
	)

	input(&length, "Password length:", "Password length must be a positive number!")
	input(&letters, "Add letters? (y/n):", "")
	input(&count, "How much passwords you want?", "Password quantity must be a positive number!")

	passwords := make([]string, count)
	for i := 0; i < count; i++ {
		password, err := generate(length, letters)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		passwords[i] = password
	}

	if err := write(passwords); err != nil {
		fmt.Println("Error dumping passwords to file:", err)
	}
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

func generate(length uint, letters bool) (password string, err error) {
	if length == 0 {
		return "", errors.New("length cannot be zero")
	}

	if letters {
		const set = "ABCDEFGHIKLMNOPQRSTVXYZabcdefghijklmnopqrstuvwxyz"

		for i := 0; i < int(length); i++ {
			if chance(50) {
				password += strconv.Itoa(rand.Intn(10))
			} else {
				password += string(set[rand.Intn(len(set))])
			}
		}

		return password, nil
	}

	for i := 0; i < int(length); i++ {
		password += strconv.Itoa(rand.Intn(10))
	}
	
	return password, nil
}

func chance(x int) bool {
	return rand.Intn(100) < x-1
}

func write(passwords []string) error {
	data := strings.Join(passwords, "\n")
	return os.WriteFile("output.txt", []byte(data), 0644)
}
