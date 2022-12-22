package model

import "fmt"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = fmt.Sprint("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = fmt.Sprint("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = fmt.Sprint("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = fmt.Sprint("given Param is not valid")
)
