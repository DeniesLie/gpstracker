package mapper

func MapSlice[TFrom any, TInto any](a []TFrom, f func(TFrom) TInto) []TInto {
	r := make([]TInto, len(a))
	for i, e := range a {
		r[i] = f(e)
	}
	return r
}
