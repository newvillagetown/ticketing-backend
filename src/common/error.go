package common

type ResError struct {
	ErrType string `json:"errType,omitempty"`
	Msg     string `json:"msg,omitempty"`
}
