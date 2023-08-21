package routes

import (
	"github.com/Razihmad/go-rest-api/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUP(app *fiber.App) {
	// app.Post("/cashier/:cashierId/login", controller.login)
	// app.Get("/cashier/:cashierId/logout", controller.logout)
	// app.Post("/cashier/cashierId/passcode", controller.passcode)

	app.Post("/cashier", controller.CreateCashier)
	app.Get("/cashier", controller.GetCashiers)
	app.Get("/cashier/:id", controller.GetCashierById)
	app.Put("/cashier/:id", controller.UpdateCashier)
	app.Delete("/cashier/:id", controller.DeleteCashier)

}
