package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ApplictaionErrorType string

const (
	DB_ERROR                        ApplictaionErrorType = "DB_ERROR"
	DUPLICATE_ERROR                 ApplictaionErrorType = "DUPLICATE_ERROR"
	INVALID_ARGUMENT                ApplictaionErrorType = "INVALID_ARGUMENT"
	INVALID_DATA                    ApplictaionErrorType = "INVALID_DATA"
	UNAUTHORIZED                    ApplictaionErrorType = "UNAUTHORIZED"
	NOT_FOUND                       ApplictaionErrorType = "NOT_FOUND"
	NOT_ALLOWED                     ApplictaionErrorType = "NOT_ALLOWED"
	UNEXPECTED_STATE                ApplictaionErrorType = "UNEXPECTED_STATE"
	CONFLICT                        ApplictaionErrorType = "CONFLICT"
	PAYMENT_AUTHORIZATION_ERROR     ApplictaionErrorType = "PAYMENT_AUTHORIZATION_ERROR"
	INSUFFICIENT_INVENTORY          ApplictaionErrorType = "INSUFFICIENT_INVENTORY"
	CART_INCOMPATIBLE_STATE         ApplictaionErrorType = "CART_INCOMPATIBLE_STATE"
	QueryRunnerAlreadyReleasedError ApplictaionErrorType = "QueryRunnerAlreadyReleasedError"
	TransactionAlreadyStartedError  ApplictaionErrorType = "TransactionAlreadyStartedError"
	TransactionNotStartedError      ApplictaionErrorType = "TransactionNotStartedError"
)

type ApplictaionError struct {
	Type    ApplictaionErrorType
	Message string
	Code    string
	Date    time.Time
}

func NewApplictaionError(typeStr ApplictaionErrorType, message string, params ...interface{}) *ApplictaionError {
	return &ApplictaionError{
		Type:    typeStr,
		Message: message,
		Date:    time.Now(),
	}
}

func (e *ApplictaionError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		e.Type = UNEXPECTED_STATE
		return err.Error()
	}

	return string(out)
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	// Status code defaults to 500
	statusCode := 500

	// Retrieve the custom status code if it's a *fiber.Error
	var e *ApplictaionError
	if errors.As(err, &e) {
		// Handle the error and set appropriate values
		switch e.Type {
		case QueryRunnerAlreadyReleasedError, TransactionAlreadyStartedError, TransactionNotStartedError, CONFLICT:
			e.Code = "invalid_state_error"
			e.Message = "The request conflicted with another request. You may retry the request with the provided Idempotency-Key."
			statusCode = 409
		case UNAUTHORIZED:
			statusCode = 401
		case PAYMENT_AUTHORIZATION_ERROR:
			statusCode = 422
		case DUPLICATE_ERROR:
			statusCode = 422
			e.Code = "invalid_request_error"
		case NOT_ALLOWED, INVALID_DATA:
			statusCode = 400
		case NOT_FOUND:
			statusCode = 404
		case DB_ERROR:
			statusCode = 500
			e.Code = "api_error"
		case UNEXPECTED_STATE, INVALID_ARGUMENT:
			statusCode = 500
		default:
			e.Code = "unknown_error"
			e.Message = "An unknown error occurred."
			e.Type = "unknown_error"
		}
	}

	// Set Content-Type: application/json; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Return status code with error message
	return ctx.Status(statusCode).JSON(e)
}
