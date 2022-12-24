package errs

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization        Type = "AUTHORIZATION"
	BadRequest           Type = "BAD_REQUEST"
	Conflict             Type = "CONFLICT"
	Internal             Type = "INTERNAL"
	NotFound             Type = "NOT_FOUND"
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE"
)

type Error struct {
	Type        Type              `json:"type"`
	Code        int               `json:"code,omitempty"`
	Source      map[string]string `json:"source,omitempty"`
	Title       string            `json:"title,omitempty"`
	Detail      string            `json:"detail,omitempty"`
	InvalidArgs []InvalidArgument `json:"invalidArgs,omitempty"`
}

// InvalidArgument - used to help extract validation errors
type InvalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func (e *Error) Error() string {
	return e.Detail
}

func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:   Authorization,
		Code:   http.StatusUnauthorized,
		Detail: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:   BadRequest,
		Code:   http.StatusBadRequest,
		Detail: reason,
	}
}

func NewConflict(name, value string) *Error {
	return &Error{
		Type:   Conflict,
		Code:   http.StatusConflict,
		Detail: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

func NewInternal() *Error {
	return &Error{
		Type:   Internal,
		Code:   http.StatusInternalServerError,
		Detail: "Internal server error.",
	}
}

func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:   NotFound,
		Code:   http.StatusNotFound,
		Detail: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

func NewServiceUnavailable() *Error {
	return &Error{
		Type:   ServiceUnavailable,
		Code:   http.StatusServiceUnavailable,
		Detail: "Service unavailable or timed out",
	}
}

func NewUnsupportedMediaType(reason string) *Error {
	return &Error{
		Type:   UnsupportedMediaType,
		Code:   http.StatusUnsupportedMediaType,
		Detail: reason,
	}
}
