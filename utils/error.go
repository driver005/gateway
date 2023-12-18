package utils

import "time"

type ApplictaionErrorType string

const (
	DB_ERROR                    ApplictaionErrorType = "DB_ERROR"
	DUPLICATE_ERROR             ApplictaionErrorType = "DUPLICATE_ERROR"
	INVALID_ARGUMENT            ApplictaionErrorType = "INVALID_ARGUMENT"
	INVALID_DATA                ApplictaionErrorType = "INVALID_DATA"
	UNAUTHORIZED                ApplictaionErrorType = "UNAUTHORIZED"
	NOT_FOUND                   ApplictaionErrorType = "NOT_FOUND"
	NOT_ALLOWED                 ApplictaionErrorType = "NOT_ALLOWED"
	UNEXPECTED_STATE            ApplictaionErrorType = "UNEXPECTED_STATE"
	CONFLICT                    ApplictaionErrorType = "CONFLICT"
	PAYMENT_AUTHORIZATION_ERROR ApplictaionErrorType = "PAYMENT_AUTHORIZATION_ERROR"
	INSUFFICIENT_INVENTORY      ApplictaionErrorType = "INSUFFICIENT_INVENTORY"
	CART_INCOMPATIBLE_STATE     ApplictaionErrorType = "CART_INCOMPATIBLE_STATE"
)

var ApplictaionErrorTypes = struct {
	DB_ERROR                    ApplictaionErrorType
	DUPLICATE_ERROR             ApplictaionErrorType
	INVALID_ARGUMENT            ApplictaionErrorType
	INVALID_DATA                ApplictaionErrorType
	UNAUTHORIZED                ApplictaionErrorType
	NOT_FOUND                   ApplictaionErrorType
	NOT_ALLOWED                 ApplictaionErrorType
	UNEXPECTED_STATE            ApplictaionErrorType
	CONFLICT                    ApplictaionErrorType
	PAYMENT_AUTHORIZATION_ERROR ApplictaionErrorType
}{
	DB_ERROR:                    "DB_ERROR",
	DUPLICATE_ERROR:             "DUPLICATE_ERROR",
	INVALID_ARGUMENT:            "INVALID_ARGUMENT",
	INVALID_DATA:                "INVALID_DATA",
	UNAUTHORIZED:                "UNAUTHORIZED",
	NOT_FOUND:                   "NOT_FOUND",
	NOT_ALLOWED:                 "NOT_ALLOWED",
	UNEXPECTED_STATE:            "UNEXPECTED_STATE",
	CONFLICT:                    "CONFLICT",
	PAYMENT_AUTHORIZATION_ERROR: "PAYMENT_AUTHORIZATION_ERROR",
}

var ApplictaionErrorCodes = struct {
	INSUFFICIENT_INVENTORY  ApplictaionErrorType
	CART_INCOMPATIBLE_STATE ApplictaionErrorType
}{
	INSUFFICIENT_INVENTORY:  "INSUFFICIENT_INVENTORY",
	CART_INCOMPATIBLE_STATE: "CART_INCOMPATIBLE_STATE",
}

type ApplictaionError struct {
	Type    ApplictaionErrorType
	Message string
	Code    string
	Date    time.Time
}

func NewApplictaionError(typeStr ApplictaionErrorType, message string, code string, params ...interface{}) *ApplictaionError {
	return &ApplictaionError{
		Type:    typeStr,
		Message: message,
		Code:    code,
		Date:    time.Now(),
	}
}

func (e *ApplictaionError) Error() string {
	return e.Message
}
