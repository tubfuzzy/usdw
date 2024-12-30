package app

import (
	"usdw/config"
	"usdw/pkg/cache"
	"usdw/pkg/db"
	"usdw/pkg/logger"

	"github.com/gofiber/fiber/v2"

	exampleusecasehandler "usdw/internal/usecase/exampleusecase/controller/http"
	exampleusecaserepositorydb "usdw/internal/usecase/exampleusecase/repository"
	exampleusecaseservice "usdw/internal/usecase/exampleusecase/service"
)

func NewApplication(app fiber.Router, logger logger.Logger, db *db.DB, cache cache.Engine, config *config.Configuration) {

	exampleRepository := exampleusecaserepositorydb.NewExampleRepository(db, logger, config)
	exampleService := exampleusecaseservice.NewExampleService(exampleRepository, config, cache, logger)
	exampleHandler := exampleusecasehandler.NewExampleHandler(exampleService, config)
	exampleHandler.InitRoute(app)
}
