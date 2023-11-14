// Error Inrastructure
package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"` //Json tags are hints for the encoder, ",omitempty" is the tag for the case if we are not setting this field, it will be ommitted from the response.
	Message string // Without Json tag, the key name will be as it is (retaining the  capital letters), but we generally want in small letters for the key.
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// Helper function
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound, // This is 404
	}
}
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError, // This is 500, e.g., database down (disconnected).
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity, // This is 422. This is different from bad-request(400) which
		// means not able to decode to DTO, but if it gets decoded and then
		// breaches business rules over the data and we can not process the entity it comes 422, i.e., StatusUnprocessableEntity.
	}
}

/*
200 is the code for success.
201 code is the payload created.
500 is for database down (disconnected).
422 code is StatusUnprocessableEntity. This is different from bad-request(400) which
		 means not able to decode to DTO, but if it gets decoded and then
		 breaches business rules over the data and we can not process the entity it comes 422, i.e., StatusUnprocessableEntity.
404 is for status not found.
301?
*/

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
