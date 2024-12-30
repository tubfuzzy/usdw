package exception

import (
	"usdw/internal/domain/constant"

	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var resp *ErrorResponse
	switch err.(type) {
	case BadRequestError:
		resp = DefaultErrBadRequest
		resp.Errors = err.Error()
	case NotFoundError:
		resp = DefaultErrNotFound
		resp.Errors = err.Error()
	case UnauthorizedError:
		resp = DefaultErrUnauthenticated
		resp.Errors = err.Error()
	default:
		resp = &DefaultErrorResponse
		resp.Errors = err.Error()

	}
	var e *fiber.Error
	if errors.As(err, &e) {
		resp.Code = constant.FRAMEWORK_ERROR
		resp.Message = e.Message
	}

	return ctx.Status(*resp.HTTPStatus).JSON(resp)
}
