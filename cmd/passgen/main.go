package main

import (
	"fmt"
	"github.com/Apolisk/passgen"
	"os"
)

func main() {
	// Generation config
	var (
		length   uint
		letters  bool
		count    int
		specials bool
	)

	input(&length, "Password length:", "Password length must be a positive number!")
	input(&letters, "Add letters? (y/n):", "")
	input(&specials, "Add specials? (y/n):", "")
	input(&count, "How much passwords you want?", "Password quantity must be a positive number!")

	passwords := make([]string, count)
	for i := 0; i < count; i++ {

		password, err := passgen.New(int(length), passgen.Config{Letters: letters, Specials: specials})
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		passwords[i] = string(password)
	}

	if err := passgen.WriteFile("output.txt", passwords); err != nil {
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
