package controller

import (
	"produtos/service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	CreateProductService service.CreateProductService
}

func NewProductController(
	createProductService *service.CreateProductService) ProductController {
	return ProductController{CreateProductService: *createProductService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Get("/api/products", controller.List)
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	response := controller.CreateProductService.Execute()
	return c.JSON(response)
}
