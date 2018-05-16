package main

type assmapError struct {
	message string
}

func newAssmapError(message string) assmapError {
	return assmapError{message}
}

func (error assmapError) Error() string {
	return error.message
}
