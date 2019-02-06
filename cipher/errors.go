package cipher

type cipherError struct {
	Message string
}

// Err ...
func Err(msg string) error {
	return &cipherError{
		Message: msg,
	}
}

// Error ...
func (e *cipherError) Error() string {
	return e.Message
}
