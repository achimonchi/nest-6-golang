package view

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
	Error   interface{} `json:"error"`
}

func SuccessCreated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Error:  err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}
func ErrNotFound() *Response {
	return &Response{
		Status: http.StatusNotFound,
		Error:  "NO_DATA",
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status: http.StatusUnauthorized,
		Error:  "UNAUTHORIZED",
	}
}
