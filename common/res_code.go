package common

func Json(code int, daa interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":   code,
		"result": daa,
	}
}
