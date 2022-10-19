package entities

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func ResponseSuccessJSON(message string, data interface{}) ResponseSuccess {
	jsonResponse := ResponseSuccess{
		Message: message,
		Data:    data,
	}
	return jsonResponse
}

func ResponseErrorJSON(code int, message string, error string) ResponseError {
	jsonResponse := ResponseError{
		Code:    code,
		Message: message,
		Error:   error,
	}
	return jsonResponse
}
