package api

import "fmt"

type GraphError struct {
	Code        ErrorCode `json:"code,omitempty"`
	Message     string    `json:"message,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (m GraphError) Error() string {
	return fmt.Sprintf("code: %d msg: %s", m.Code, m.Description)
}
