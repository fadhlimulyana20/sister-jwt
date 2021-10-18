package model

type Json struct {
	Status           string      `json:"status"`
	Message          string      `json:"message,omitempty"`
	Data             interface{} `json:"data,omitempty"`
	ErrorCode        int         `json:"errorCode,omitempty"`
	ErrorDescription string      `json:"errorDescription,omitempty"`
}
