package validation

type ValidationError struct {
	Res Result
}

func (e ValidationError) Error() string {
	return "data is invalid, see result for details"
}
