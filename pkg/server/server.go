package server

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"usdw/config"

	apiv1 "usdw/internal/app"
	cachePkg "usdw/pkg/cache"
	"usdw/pkg/common/exception"
	dbPkg "usdw/pkg/db"
	loggerPkg "usdw/pkg/logger"
)

type Server struct {
	app    *fiber.App
	conf   *config.Configuration
	logger loggerPkg.Logger
	cache  cachePkg.Engine
	db     *dbPkg.DB
}

func New() (*Server, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := loggerPkg.NewLogger(conf)
	cacheEngine, err := cachePkg.NewCache(conf)
	if err != nil {
		return nil, err
	}
	db, err := dbPkg.NewDB(conf)
	if err != nil {
		return nil, err
	}
	// db := &db.DB{}
	app := NewFiberApp(conf, logger, cacheEngine, db)

	return &Server{
		conf:   conf,
		logger: logger,
		cache:  cacheEngine,
		db:     db,
		app:    app,
	}, nil
}

func NewFiberApp(
	conf *config.Configuration,
	logger loggerPkg.Logger,
	cacheEngine cachePkg.Engine,
	db *dbPkg.DB,
) *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		ReadTimeout:  time.Second * conf.Server.ReadTimeout,
		WriteTimeout: time.Second * conf.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	app.Use(fiberLog.New(fiberLog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	apiv1.NewApplication(v1, logger, db, cacheEngine, conf)

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Health check successful",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		panic(exception.NotFoundError{Message: "path " + c.Path() + " does not exist."})
	})

	return app
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Configuration {
	return serv.conf
}

func (serv Server) Logger() loggerPkg.Logger {
	return serv.logger
}

func (serv Server) DB() *dbPkg.DB {
	return serv.db
}

func (serv Server) Cache() cachePkg.Engine {
	return serv.cache
}
