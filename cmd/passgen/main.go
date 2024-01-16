package main

import (
	"fmt"
	"os"

	"github.com/Apolisk/passgen"
)

func main() {
	var (
		count    int
		length   int
		letters  bool
		specials bool
	)

	input(&count, "How much passwords you want?", "Password count must be a positive number!")
	input(&length, "Password length:", "Password length must be a positive number!")
	input(&letters, "Add letters? (y/n):", "")
	input(&specials, "Add specials? (y/n):", "")

	config := passgen.Config{
		Letters:  letters,
		Specials: specials,
	}

	pwds, err := passgen.Many(count, length, config)
	if err != nil {
		fmt.Println("Error generating a password:", err)
		return
	}

	if err := pwds.WriteFile("output.txt"); err != nil {
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
