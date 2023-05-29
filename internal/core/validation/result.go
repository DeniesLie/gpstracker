package validation

type FieldResult struct {
	FieldName string   `json:"fieldName"`
	Messages  []string `json:"messages"`
}

type Result struct {
	IsValid      bool     `json:"isValid"`
	Messages     []string `json:"messages"`
	currentField string
}

func ValidResult() Result {
	return Result{IsValid: true}
}

func InvalidResult(message string) Result {
	return Result{IsValid: false, Messages: []string{message}}
}

func (r *Result) Field(fieldName string) *Result {
	r.IsValid = true
	r.currentField = fieldName
	return r
}

func AggregateResult(res1 Result, res2 Result) Result {
	result := Result{}
	result.IsValid = res1.IsValid && res2.IsValid
	result.Messages = append(res1.Messages, res2.Messages...)
	return result
}
