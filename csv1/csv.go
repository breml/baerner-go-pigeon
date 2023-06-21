package csv

func line(values, lastvalue any) (any, error) {
	vals := toAnySlice(values)

	vals = append(vals, lastvalue)

	return vals, nil
}

func doublequotedvalue(chars any) (any, error) {
	chrs := toAnySlice(chars)

	var value string

	for _, chr := range chrs {
		value += chr.(string)
	}
	return value, nil
}
