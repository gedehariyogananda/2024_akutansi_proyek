package Utils

import (
	"errors"
	"net/http"
)

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func HandleStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	case func() bool {
		var notFoundErr *NotFoundError
		return errors.As(err, &notFoundErr)
	}():
		return http.StatusNotFound
	case func() bool {
		var validationErr *ValidationError
		return errors.As(err, &validationErr)
	}():
		return http.StatusBadRequest
	case func() bool {
		var unauthorizedErr *UnauthorizedError
		return errors.As(err, &unauthorizedErr)
	}():
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
