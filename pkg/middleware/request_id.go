package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDKey = "x-request-id"

// RequestIDMiddleware ensures every request has a unique x-request-id
func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(RequestIDKey)

		// Generate a new request ID if missing
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Set request ID in response headers
		c.Set(RequestIDKey, requestID)

		// Pass request ID to context
		ctx := context.WithValue(c.Context(), RequestIDKey, requestID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
