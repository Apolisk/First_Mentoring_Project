package passgen

import (
	"errors"
	"math/rand"
	"os"
	"strings"
)

type (
	// Password is a password representation.
	Password string

	// Passwords is a list of passwords.
	Passwords []Password
)

// Set of rules.
const (
	digits   = "0123456789"
	letters  = "ABCDEFGHIKLMNOPQRSTVXYZabcdefghijklmnopqrstuvwxyz"
	specials = "!@#$%^&*()/?{}[]|"
)

// Config defines password generation settings.
type Config struct {
	Letters   bool
	Specials  bool
	BeginChar bool
}

// New generates a new password with given length n.
func New(n int, c Config) (Password, error) {
	if n <= 0 {
		return "", errors.New("length must be a positive number")
	}

	rules := []string{digits}
	if c.BeginChar {
		rules = append(rules, string(letters[rand.Intn(len(letters))]))
	}
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
	for i := 0; i < count; i++ {
		ps[i], err = New(n, c)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

// generate generates a random password with given rules.
func generate(n int, rules ...string) Password {
	p := make([]byte, 0, n)

	//If we use BeginChar rule
	if len(rules[1]) == 1 {
		p = append(p, rules[1][0])
		n -= 1
	}

	for i := 0; i < n; i++ {
		// Pick a random rule.
		rule := rules[rand.Intn(len(rules))]
		// Pick a random character from the rule.\
		p = append(p, rule[rand.Intn(len(rule))])
	}
	return Password(p)
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
