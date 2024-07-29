package api

import "net/http"

type ErrorCode int

const (
	Unknown ErrorCode = iota
	GraphNotFound
	PathNotFound
	GraphCreationFailed
)

var statuCode = map[ErrorCode]int{
	GraphNotFound:       http.StatusNotFound,
	PathNotFound:        http.StatusNotFound,
	GraphCreationFailed: http.StatusBadRequest,
}

func (e ErrorCode) HTTPStatus() int {
	if code, ok := statuCode[e]; ok {
		return code
	}
	return http.StatusInternalServerError
}
