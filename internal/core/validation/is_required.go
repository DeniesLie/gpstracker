package validation

import "time"

func (r *Result) IsRequired(value string) *Result {
	if value == "" {
		r.IsValid = false
		r.Messages = append(r.Messages, r.currentField+" is required")
	}
	return r
}

func (r *Result) IsRequiredTime(value time.Time) *Result {
	if value.IsZero() {
		r.IsValid = false
		r.Messages = append(r.Messages, r.currentField+" is required")
	}
	return r
}
