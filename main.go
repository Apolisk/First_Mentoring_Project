package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var length int
	var uppercase bool

	input(&length, "Password length:", "Password length MUST be a number!")
	input(&uppercase, "Type 1  if you want to use uppercase character, else type 0", "Must be bool type")

	fmt.Println(generate(length, uppercase))
}

func input(v any, m, e string) {
	fmt.Print(m, " ")
	_, err := fmt.Scanln(v)
	if err != nil {
		fmt.Printf("\n%s (%s)\n", e, err)
		os.Exit(1)
	}
}

func generate(l int, u bool) string {
	var password string
	const UpperLetters = "ABCDEFGHIKLMNOPQRSTVXYZ"

	if l <= 0 {
		return "Password length must  be greater then 0"
	} else if l > 0 && u == false {
		for i := 0; i < l; i++ {
			password += strconv.Itoa(rand.Intn(10))
		}
	} else {
		for i := 0; i < l; i++ {
			randnum := strconv.Itoa(rand.Intn(10))
			randchar := string(rune(UpperLetters[rand.Intn(len(UpperLetters))]))
			password += randnum
			password += randchar
		}
	}
	return password
}
