package errors

// Error the constant error type,
type Error string

// New creates the new Error.
func New(text string) Error {
	return Error(text)
}

func (e Error) Error() string {
	return string(e)
}
