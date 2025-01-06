package http

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"usdw/config"
	"usdw/internal/domain"
)

type BankFeedHandler struct {
	domain.BankFeedService
	config.Configuration
}

func NewBankFeedHandler(bankFeedService domain.BankFeedService, config *config.Configuration) BankFeedHandler {
	return BankFeedHandler{
		BankFeedService: bankFeedService,
		Configuration:   *config,
	}
}

func (h *BankFeedHandler) InitRoute(app fiber.Router) {
	app.Post("/feed-connections", h.CreateConnections)
	app.Get("/feed-connections", h.GetConnections)
	app.Get("/feed-connections/:id", h.GetConnectionByID)
	app.Delete("/feed-connections/:id", h.DeleteFeedConnection)
	app.Post("/statements", h.PostStatements)
	app.Get("/statements", h.GetStatements)
	app.Get("/statements/:id", h.GetStatementByID)
}

func (h *BankFeedHandler) CreateConnections(c *fiber.Ctx) error {
	var request domain.CreateConnectionsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	response, err := h.BankFeedService.CreateConnections(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *BankFeedHandler) GetConnections(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "20"))

	response, err := h.BankFeedService.GetConnections(c.Context(), page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *BankFeedHandler) GetConnectionByID(c *fiber.Ctx) error {
	feedConnectionID := c.Params("id")

	response, err := h.BankFeedService.GetConnectionByID(c.Context(), feedConnectionID)
	if err != nil {
		if err.Error() == "feed connection not found" {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *BankFeedHandler) DeleteFeedConnection(c *fiber.Ctx) error {
	feedConnectionID := c.Params("id")

	response, err := h.BankFeedService.DeleteConnection(c.Context(), feedConnectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *BankFeedHandler) PostStatements(c *fiber.Ctx) error {
	var request domain.PostStatementRequest

	// Parse the request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Call service layer to process statements
	response, err := h.BankFeedService.PostStatements(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (h *BankFeedHandler) GetStatements(c *fiber.Ctx) error {
	// Parse optional query parameters
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page parameter",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "50"))
	if err != nil || pageSize < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid pageSize parameter",
		})
	}

	response, err := h.BankFeedService.GetStatements(c.Context(), page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *BankFeedHandler) GetStatementByID(c *fiber.Ctx) error {
	statementID := c.Params("id")

	response, err := h.BankFeedService.GetStatementByID(c.Context(), statementID)
	if err != nil {
		if err.Error() == "statement not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Statement not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response)
}
