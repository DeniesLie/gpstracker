package service

import "fmt"

type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("resource %s was not found", e.Resource)
}

type BusinessError struct {
	Message string
}

func (e BusinessError) Error() string {
	return e.Message
}
