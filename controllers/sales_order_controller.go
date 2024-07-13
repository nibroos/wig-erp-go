package controllers

import (
	"derp_api_go/services"

	"github.com/gofiber/fiber/v2"
)

type SalesOrderController struct {
    service services.SalesOrderService
}

func NewSalesOrderController(service services.SalesOrderService) *SalesOrderController {
    return &SalesOrderController{service}
}

func (c *SalesOrderController) GetSalesOrders(ctx *fiber.Ctx) error {
    isAll := false
    orders, err := c.service.GetSalesOrders(isAll)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(orders)
}
