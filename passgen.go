package passgen

// Config defines password generation settings.
type Config struct {
	Letters  bool
	Specials bool
}

// New generates a new password.
func New(n int, c Config) (Password, error) {
	return "", nil // TODO: implement
}

// Password is a password representation.
type Password string

// String implements fmt.Stringer.
func (p Password) String() string {
	return string(p)
}

// WriteFile writes the list of passwords to the file.
func WriteFile(path string, pws []Password) error {
	return nil
}
