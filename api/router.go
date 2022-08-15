package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	_ "github.com/wil-g2/unico-challenge/docs"
)

func SetupRoutes(
	app fiber.Router, fairHandler FairHandler) {
	app.Get("/", monitor.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/fairs", fairHandler.List)
	app.Get("/fairs/:id", fairHandler.Find)
	app.Get("/fairs/name/:name", fairHandler.FindByName)
	app.Put("/fairs/:id", fairHandler.Update)
	app.Post("/fairs", fairHandler.Create)
	app.Delete("/fairs/:id", fairHandler.Delete)
}
