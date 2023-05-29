package validation

import "fmt"

type number interface {
	int | uint | uint64 | float32 | float64
}

func (r *Result) ValueIsBetween(value float64, from float64, to float64) *Result {
	return IsValueBetween(r, value, from, to)
}

func IsValueBetween[N number](r *Result, value N, from N, to N) *Result {
	if value < from || value > to {
		r.IsValid = false
		message := fmt.Sprintf("%v value must be between %v and %v", r.currentField, from, to)
		r.Messages = append(r.Messages, message)
	}
	return r
}
