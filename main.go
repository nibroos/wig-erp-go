package main

import (
	"derp_api_go/controllers"
	"derp_api_go/database"
	"derp_api_go/repositories"
	"derp_api_go/routes"
	"derp_api_go/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Initialize the database
    database.InitDatabase()

    // Initialize repositories, services, and controllers
    salesOrderRepo := repositories.NewSalesOrderRepository(database.DB)
    salesOrderService := services.NewSalesOrderService(salesOrderRepo)
    salesOrderController := controllers.NewSalesOrderController(salesOrderService)

    // Setup routes
    routes.SetupRoutes(app, salesOrderController)

    // Start the application
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(app.Listen(":" + port))
}
