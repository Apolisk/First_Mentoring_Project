package passgen

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	letters  = "ABCDEFGHIKLMNOPQRSTVXYZabcdefghijklmnopqrstuvwxyz"
	specials = "!@#$%^&*()/?{}[]|"
)

// Config defines password generation settings.
type Config struct {
	Letters  bool
	Specials bool
}

// New generates a new password.
func New(n int, c Config) (Password, error) {
	if n == 0 {
		return "", errors.New("length cannot be zero")
	}
	var password Password
	switch {
	case c.Letters && c.Specials:
		for i := 0; i < n; i++ {
			if chance(33) {
				password += Password(strconv.Itoa(rand.Intn(10)))
			} else if chance(66) {
				password += Password(specials[rand.Intn(len(specials))])
			} else {
				password += Password(letters[rand.Intn(len(letters))])
			}
		}
	case c.Letters:
		for i := 0; i < n; i++ {
			if chance(50) {
				password += Password(strconv.Itoa(rand.Intn(10)))
			} else {
				password += Password(letters[rand.Intn(len(letters))])
			}
		}
	case c.Specials:
		for i := 0; i < n; i++ {
			if chance(50) {
				password += Password(strconv.Itoa(rand.Intn(10)))
			} else {
				password += Password(specials[rand.Intn(len(specials))])
			}
		}
	default:
		for i := 0; i < n; i++ {
			password += Password(strconv.Itoa(rand.Intn(10)))
		}
	}

	return password, nil
}

// genPases
func genPases(count, length int, letters, specials bool) (passwords []string) {
	passwords = make([]string, count)
	for i := 0; i < count; i++ {

		password, err := New(length, Config{Letters: letters, Specials: specials})
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		passwords[i] = string(password)
	}
	return passwords
}

func chance(x int) bool {
	return rand.Intn(100) < x-1
}

// Password is a password representation.
type Password string

// String implements fmt.Stringer.

func (p Password) String() string {
	return string(p)
}

// WriteFile writes the list of passwords to the file.
func WriteFile(path string, pws []string) error {
	data := strings.Join(pws, "\n")
	return os.WriteFile(path, []byte(data), 0644)
}
