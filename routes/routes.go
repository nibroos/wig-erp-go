package routes

import (
	"derp_api_go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, salesOrderController *controllers.SalesOrderController) {
    app.Get("/orders/sales-orders", salesOrderController.GetSalesOrders)
}
