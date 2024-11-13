package errors

import "errors"

var (
    ErrUserNotFound      = errors.New("user not found")
    ErrUserAlreadyExists = errors.New("user with this email already exists")
    ErrInsufficientBalance = errors.New("insufficient balance")
    ErrInvalidInput     = errors.New("invalid input data")
    ErrLotNotFound = errors.New("lot not found")
    ErrInvalidPage = errors.New("invalid page parameters")
)

type ErrorType string

const (
    ErrorTypeNotFound      ErrorType = "NOT_FOUND"
    ErrorTypeValidation    ErrorType = "VALIDATION"
    ErrorTypeConflict      ErrorType = "CONFLICT"
    ErrorTypeInternal      ErrorType = "INTERNAL"
    ErrorTypeUnauthorized  ErrorType = "UNAUTHORIZED"
)

type AppError struct {
    Type    ErrorType `json:"type"`
    Message string    `json:"message"`
    Err     error    `json:"-"`
}

func (e *AppError) Error() string {
    return e.Message
}

func New(errType ErrorType, message string, err error) *AppError {
    return &AppError{
        Type:    errType,
        Message: message,
        Err:     err,
    }
}

func NewValidationError(message string) *AppError {
    return &AppError{
        Type:    ErrorTypeValidation,
        Message: message,
    }
}

func NewNotFoundError(message string) *AppError {
    return &AppError{
        Type:    ErrorTypeNotFound,
        Message: message,
    }
}

func NewConflictError(message string) *AppError {
    return &AppError{
        Type:    ErrorTypeConflict,
        Message: message,
    }
}
