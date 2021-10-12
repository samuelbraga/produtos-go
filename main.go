package main

import (
	"produtos/config"
	"produtos/controller"
	"produtos/exception"
	"produtos/repository"
	"produtos/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.New()

	productRepository := repository.NewProductRepository()
	createProductService := service.NewCreateProductService(&productRepository)
	productController := controller.NewProductController(&createProductService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	productController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
