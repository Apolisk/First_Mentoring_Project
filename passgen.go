package main

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"sync"
)

type (
	// Password is a password representation.
	Password string

	// Passwords is a list of passwords.
	Passwords []Password
)

type charset = []byte

// Set of rules.
var (
	digits   = charset("0123456789")
	letters  = charset("abcdefghijklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTVXYZ")
	specials = charset("!@#$%^&*()/?{}[]|")
)

// Config defines password generation settings.
type Config struct {
	Letters  bool
	Specials bool
}

// New generates a new password with given length n.
func New(n int, c Config) (Password, error) {
	if n <= 0 {
		return "", errors.New("length must be a positive number")
	}

	rules := []charset{digits}
	if c.Letters {
		rules = append(rules, letters)
	}
	if c.Specials {
		rules = append(rules, specials)
	}

	return generate(n, rules...), nil
}

// Many generates a list of passwords with given length n and count of them.
func Many(count, n int, c Config) (ps Passwords, err error) {
	ps = make(Passwords, count)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			ps[i], err = New(n, c)
			mu.Unlock()
			if err != nil {
				return
			}
		}(i)
	}
	wg.Wait()
	return ps, nil
}

// generate generates a random password with given rules.
func generate(n int, rules ...charset) Password {
	p := make([]byte, n)

	// The first character should always be a letter.
	p[0] = pick(letters)

	for i := 1; i < n; i++ {
		// Pick a random rule.
		rule := pick(rules)
		// Pick a random character from the rule.
		p[i] = pick(rule)
	}

	return Password(p)
}

// pick picks a random value from a given slice.
func pick[T any](a []T) T {
	return a[rand.Intn(len(a))]
}

// String implements fmt.Stringer.
func (p Password) String() string {
	return string(p)
}

// Strings converts the Passwords type to slice of string.
func (ps Passwords) Strings() []string {
	s := make([]string, len(ps))
	for i, p := range ps {
		s[i] = p.String()
	}
	return s
}

// WriteFile writes the content of passwords to the file.
func (ps Passwords) WriteFile(path string) error {
	data := strings.Join(ps.Strings(), "\n")
	return os.WriteFile(path, []byte(data), 0644)
}
