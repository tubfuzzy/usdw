package http

import (
	"usdw/config"
	"usdw/internal/domain"

	"git.matador.ais.co.th/esport-development-team/common/go-common-sdk/requestcontext"
	"github.com/gofiber/fiber/v2"
)

type ExampleHandler struct {
	domain.ExampleService
	config.Configuration
}

func NewExampleHandler(exampleService domain.ExampleService, config *config.Configuration) ExampleHandler {
	return ExampleHandler{
		ExampleService: exampleService,
		Configuration:  *config,
	}
}

func (h ExampleHandler) InitRoute(app fiber.Router) {
	app.Get("/test", h.TestHandler)
}

func (h ExampleHandler) TestHandler(ctx *fiber.Ctx) error {
	context := requestcontext.SetRequestID(ctx.Context(), "TestHandler")
	data, err := h.ExampleService.GetUserTest(context)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"code":    "50000",
			"message": "Operation unsuccessful",
		})
	}

	resp := map[string]interface{}{
		"code":    "20000",
		"message": "Operation successful",
		"data":    data,
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
