package rest

import (
	log "github.com/sirupsen/logrus"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Router ...
type Router struct {
	*fiber.App
	Config *config.RESTServer
}

// NewRouter new fiber app with middlewares
func NewRouter(c *config.RESTServer) *Router {
	app := fiber.New()

	// Default configuration fiber middlewares
	// https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())
	// https://docs.gofiber.io/api/middleware/logger
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Europe/Moscow",
	}))
	AddRoutes(app)
	return &Router{
		app,
		c,
	}
}

// RunAPI start listen specified <addr:port>
func (r *Router) RunAPI() {
	if err := r.Listen("0.0.0.0:" + r.Config.Port); err != nil {
		log.Fatalf("cant Start server due: %s", err)
	}
}

// AddRoutes ...
func AddRoutes(app *fiber.App) {
	app.Get("/fibonacci", FibonacciSliceHandler)
}
