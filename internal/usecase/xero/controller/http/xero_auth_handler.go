package http

import (
	"github.com/gofiber/fiber/v2"
	"usdw/config"
)

// XeroAuthHandler handles OAuth 2.0 authentication for Xero
type XeroAuthHandler struct{}

// NewXeroAuthHandler initializes Xero auth routes

func NewXeroAuthHandler(app fiber.Router) {
	handler := &XeroAuthHandler{}
	app.Get("/oauth/callback", handler.HandleOAuthCallback)
}

// HandleOAuthCallback processes the redirect from Xero
func (h *XeroAuthHandler) HandleOAuthCallback(c *fiber.Ctx) error {
	// Get authorization code from query parameters
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization code not found"})
	}

	// Exchange authorization code for access token
	token, err := config.XeroOAuthConfig.Exchange(c.Context(), code)
	if err != nil {
		//logger.Error("Failed to exchange token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// Store token (in-memory for now; use a DB or cache in production)
	//logger.Info("Xero Access Token:", token.AccessToken)

	// Redirect user to success page or API response
	return c.JSON(fiber.Map{
		"message":      "Xero OAuth successful!",
		"access_token": token.AccessToken,
		"expires_in":   token.Expiry,
	})
}
