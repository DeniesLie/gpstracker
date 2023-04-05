package validation

import "fmt"

func (r *Result) HasValues(value string, allowedValues []string) *Result {
	if !contains(allowedValues, value) {
		r.IsValid = false
		message := fmt.Sprintf("%s value is not allowed. Allowed values: %v", r.currentField, allowedValues)
		r.Messages = append(r.Messages, message)
	}
	return r
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
