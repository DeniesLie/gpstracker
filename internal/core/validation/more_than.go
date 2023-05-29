package validation

import "fmt"

func (r *Result) IsMoreThan(value uint, lowBound uint) *Result {
	return IsMoreThan(r, value, lowBound)
}

func IsMoreThan[N number](r *Result, value N, lowBound N) *Result {
	if value <= lowBound {
		r.IsValid = false
		message := fmt.Sprintf("%v value must be more than %v", r.currentField, lowBound)
		r.Messages = append(r.Messages, message)
	}
	return r
}
