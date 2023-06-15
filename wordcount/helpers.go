package wordcount

func toAnySlice(v any) []any {
	if v == nil {
		return nil
	}
	switch v1 := v.(type) {
	case []any:
		return v1
	case any:
		return []any{v1}
	default:
		return nil
	}
}
