package models

type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func CustomError(message string, code int) *Error {
	return &Error{
		StatusCode: code,
		Message:    message,
	}
}
