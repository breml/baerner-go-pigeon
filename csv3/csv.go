package csv

type CSV []Line

type Line []Value

type Value string

func doc(lines any) (any, error) {
	lns := toAnySlice(lines)

	csv := make(CSV, 0, len(lns))
	for _, l := range lns {
		csv = append(csv, l.(Line))
	}
	return csv, nil
}

func line(values, lastvalue any) (any, error) {
	vals := toAnySlice(values)

	fields := make(Line, 0, len(vals)+1)
	for _, v := range vals {
		fields = append(fields, v.(Value))
	}
	fields = append(fields, lastvalue.(Value))

	return fields, nil
}

func lettersequence(chars []byte) (any, error) {
	return Value(chars), nil
}

func doublequotedvalue(chars any) (any, error) {
	chrs := toAnySlice(chars)

	var value string

	for _, chr := range chrs {
		value += chr.(string)
	}
	return Value(value), nil
}
