package internal

func SuccessResponseMessage(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": data,
	}
}

func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": true,
		"data":   data,
	}
}

func ErrorResponse(err string) map[string]interface{} {
	return map[string]interface{}{
		"status": false,
		"error":  err,
	}
}
