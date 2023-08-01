package main

import (
	"github.com/gofiber/fiber/v2"
	"rateService/Storage"
	"rateService/User"
)

func main() {
	Storage.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {
	app.Get("/getAll", User.GetUsers)
	app.Get("/delete/:id", User.DeleteUserById)
	app.Get("/getById/:id", User.GetUserById)
	app.Post("/create", User.CreateUser)
	app.Post("/update/:id", User.UpdateById)
}
