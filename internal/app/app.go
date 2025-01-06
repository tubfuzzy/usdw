package app

import (
	"github.com/go-resty/resty/v2"
	"usdw/config"
	"usdw/pkg/cache"
	"usdw/pkg/db"
	"usdw/pkg/logger"

	"github.com/gofiber/fiber/v2"

	bankfeedhandler "usdw/internal/usecase/bankfeed/controller/http"
	bankfeedrepository "usdw/internal/usecase/bankfeed/repository"
	bankfeedservice "usdw/internal/usecase/bankfeed/service"

	xeroauthhandler "usdw/internal/usecase/xero/controller/http"
)

func NewApplication(app fiber.Router, logger logger.Logger, client *resty.Client, _ *db.DB, cache cache.Engine, config *config.Configuration) {
	bankFeedRepository := bankfeedrepository.NewBankFeedRepository(client, config)
	bankFeedService := bankfeedservice.NewBankFeedService(bankFeedRepository, config, cache, logger)
	bankFeedHandler := bankfeedhandler.NewBankFeedHandler(bankFeedService, config)
	bankFeedHandler.InitRoute(app)

	xeroauthhandler.NewXeroAuthHandler(app)
}
