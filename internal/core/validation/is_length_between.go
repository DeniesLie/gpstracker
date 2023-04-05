package validation

import "fmt"

func (r *Result) LengthIsBetween(value string, from int, to int) *Result {
	length := len(value)
	if length < from || length > to {
		r.IsValid = false
		message := fmt.Sprintf("Length of '%s' value must be between %d and %d", r.currentField, from, to)
		r.Messages = append(r.Messages, message)
	}
	return r
}
