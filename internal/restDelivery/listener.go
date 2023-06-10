package restDelivery

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"ozon/internal/restDelivery/restHendlers"
)

func Hearing(app *fiber.App) {
	app.Get("/get", restHendlers.Get)
	app.Post("/post", restHendlers.Post)
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error with building connection")
	}
}
