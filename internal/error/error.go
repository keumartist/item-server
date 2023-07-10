package customerror

import "fmt"

type CustomError struct {
	Code    string
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	ErrUserNotFound       = &CustomError{Code: "USER_NOT_FOUND", Message: "User not found"}
	ErrItemNotFound       = &CustomError{Code: "ITEM_NOT_FOUND", Message: "Item not found"}
	ErrInternal           = &CustomError{Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"}
	ErrBadRequest         = &CustomError{Code: "BAD_REQUEST", Message: "Request is invalid"}
	ErrUnauthorized       = &CustomError{Code: "UNAUTHORIZED", Message: "Token is unauthorized"}
	ErrDuplicatedUserItem = &CustomError{Code: "ITEM_DUPLICATED", Message: "User already had an item"}
)
