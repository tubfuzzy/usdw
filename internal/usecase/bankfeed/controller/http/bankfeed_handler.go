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

// @Summary Create a new feed connection
// @Description Creates a new feed connection for bank feeds
// @Tags FeedConnections
// @Accept json
// @Produce json
// @Success 201 {object} domain.CreateConnectionsResponse
// @Failure 400 {object} exception.ErrorResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /feed-connections [post]
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

// @Summary Get all feed connections
// @Description Retrieves a paginated list of feed connections
// @Tags FeedConnections
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} domain.ConnectionsResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /feed-connections [get]
func (h *BankFeedHandler) GetConnections(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "20"))

	response, err := h.BankFeedService.GetConnections(c.Context(), page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

// @Summary Get feed connection by ID
// @Description Retrieves details of a feed connection by its ID
// @Tags FeedConnections
// @Accept json
// @Produce json
// @Param id path string true "Feed Connection ID"
// @Success 200 {object} domain.Connection
// @Failure 404 {object} exception.ErrorResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /feed-connections/{id} [get]
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

// @Summary Delete a feed connection
// @Description Deletes a feed connection by its ID
// @Tags FeedConnections
// @Accept json
// @Produce json
// @Param id path string true "Feed Connection ID"
// @Success 200 {object} domain.DeleteResult
// @Failure 500 {object} exception.ErrorResponse
// @Router /feed-connections/{id} [delete]
func (h *BankFeedHandler) DeleteFeedConnection(c *fiber.Ctx) error {
	feedConnectionID := c.Params("id")

	response, err := h.BankFeedService.DeleteConnection(c.Context(), feedConnectionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(response)
}

// @Summary Create a new statement
// @Description Creates a new statement with details
// @Tags Statements
// @Accept json
// @Produce json
// @Success 202 {object} domain.PostStatementResponse
// @Failure 400 {object} exception.ErrorResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /statements [post]
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

// @Summary Get all statements
// @Description Retrieves a paginated list of statements
// @Tags Statements
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} domain.GetStatementsResponse
// @Failure 400 {object} exception.ErrorResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /statements [get]
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

// @Summary Get statement by ID
// @Description Retrieves details of a statement by its ID
// @Tags Statements
// @Accept json
// @Produce json
// @Param id path string true "Statement ID"
// @Success 200 {object} domain.StatementResult
// @Failure 404 {object} exception.ErrorResponse
// @Failure 500 {object} exception.ErrorResponse
// @Router /statements/{id} [get]
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
