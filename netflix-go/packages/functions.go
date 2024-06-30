package packages

func GetInt64(data any) (res int64) {
	switch v := data.(type) {
	case int:
		res = int64(v)
	case int64:
		res = int64(v)
	case float32:
		res = int64(v)
	case float64:
		res = int64(v)
	}
	return res
}

func GetString(data any) (res string) {
	switch v := data.(type) {
	case string:
		res = string(v)
	}
	return res
}
