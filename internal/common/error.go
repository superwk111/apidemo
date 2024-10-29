package common

import (
	"fmt"
	"errors"
)

// CodeError is a custom error type that includes an error code and message
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

// NewCodeError creates a new CodeError
func NewCodeError(code int, msg string) error {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

// Common error codes
const (
	SUCCESS = 200
	ERROR   = 500

	// Auth errors (1000-1999)
	InvalidParams      = 1001
	UserNotExist      = 1002
	UserAlreadyExists = 1003
	PasswordError     = 1004
	GenerateTokenError = 1005
	TokenExpired      = 1006
	InvalidToken      = 1007

	// Database errors (2000-2999)
	DatabaseError = 2001
	RecordNotFound = 2002
	
	// Business errors (3000-3999)
	BusinessError = 3001
)

// Error messages map
var errorMap = map[int]string{
	SUCCESS:           "success",
	ERROR:            "internal server error",
	InvalidParams:    "invalid parameters",
	UserNotExist:     "user does not exist",
	UserAlreadyExists: "user already exists",
	PasswordError:    "incorrect password",
	GenerateTokenError: "failed to generate token",
	TokenExpired:     "token has expired",
	InvalidToken:     "invalid token",
	DatabaseError:    "database error",
	RecordNotFound:   "record not found",
	BusinessError:    "business error",
}

// GetMsg gets the error message based on code
func GetMsg(code int) string {
	msg, ok := errorMap[code]
	if ok {
		return msg
	}
	return errorMap[ERROR]
}

// Example usage functions
func InvalidParamsError() error {
	return NewCodeError(InvalidParams, GetMsg(InvalidParams))
}

func UserNotExistError() error {
	return NewCodeError(UserNotExist, GetMsg(UserNotExist))
}

func UserAlreadyExistsError() error {
	return NewCodeError(UserAlreadyExists, GetMsg(UserAlreadyExists))
}

func PasswordErrorError() error {
	return NewCodeError(PasswordError, GetMsg(PasswordError))
}

func TokenError() error {
	return NewCodeError(InvalidToken, GetMsg(InvalidToken))
}

// IsCodeError checks if an error is a CodeError
func IsCodeError(err error) (*CodeError, bool) {
	if err == nil {
		return nil, false
	}
	codeErr, ok := err.(*CodeError)
	return codeErr, ok
}

// HandleError converts any error to an appropriate CodeError
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	
	// If it's already a CodeError, return it directly
	if codeErr, ok := IsCodeError(err); ok {
		return codeErr
	}
	
	// Convert other errors to internal server error
	return NewCodeError(ERROR, GetMsg(ERROR))
}

// Success creates a success response
func Success(data interface{}) *CodeError {
	return &CodeError{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	}
}

// DatabaseErrorf creates a formatted database error
func DatabaseErrorf(format string, args ...interface{}) error {
	return NewCodeError(DatabaseError, fmt.Sprintf(format, args...))
}

func NewGenerateTokenError() error {
    return errors.New("failed to generate token")
}