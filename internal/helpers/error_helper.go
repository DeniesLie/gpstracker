package helpers

func Message(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func ErrorsEqual(err1 error, err2 error) bool {
	return Message(err1) == Message(err2)
}
