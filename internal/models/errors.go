package models

import "github.com/gofiber/fiber/v2"

const (
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeInvalidCredentials = "INVALID_CREDENTIALS"
	ErrCodeMissingAuth        = "MISSING_AUTH_HEADER"
	ErrCodeInvalidToken       = "INVALID_TOKEN"
	ErrCodeExpiredToken       = "EXPIRED_TOKEN"

	ErrCodeForbidden         = "FORBIDDEN"
	ErrCodeInsufficientPerms = "INSUFFICIENT_PERMISSIONS"

	ErrCodeValidationFailed = "VALIDATION_FAILED"
	ErrCodeInvalidInput     = "INVALID_INPUT"
	ErrCodeInvalidFormat    = "INVALID_FORMAT"


	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeAlreadyExists = "ALREADY_EXISTS"


	ErrCodeInternalError = "INTERNAL_ERROR"
	ErrCodeDatabaseError = "DATABASE_ERROR"
)

func NewErrorResponse(message, code, requestID string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Message:   message,
			Code:      code,
			RequestID: requestID,
		},
	}
}

func SendError(c *fiber.Ctx, status int, message, code, requestID string) error {
	return c.Status(status).JSON(NewErrorResponse(message, code, requestID))
}

func SendBadRequest(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusBadRequest, message, ErrCodeInvalidInput, requestID)
}

func SendUnauthorized(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusUnauthorized, message, ErrCodeUnauthorized, requestID)
}

func SendForbidden(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusForbidden, message, ErrCodeForbidden, requestID)
}

func SendNotFound(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusNotFound, message, ErrCodeNotFound, requestID)
}

func SendConflict(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusConflict, message, ErrCodeAlreadyExists, requestID)
}

func SendInternalError(c *fiber.Ctx, message, requestID string) error {
	return SendError(c, fiber.StatusInternalServerError, message, ErrCodeInternalError, requestID)
}
