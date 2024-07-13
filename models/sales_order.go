package models

import "time"

type SalesOrder struct {
    ID              uint      `gorm:"primaryKey"`
    OrderNumber     string    `gorm:"column:order_number"`
    BuyerPONumber   string    `gorm:"column:buyer_po_number"`
    CustomerID      uint      `gorm:"column:customer_id"`
    OrderDate       time.Time `gorm:"column:order_date"`
    ShippingDate    time.Time `gorm:"column:shipping_date"`
    Status          string    `gorm:"column:status"`
    TotalAmount     float64   `gorm:"column:total_amount"`
    TotalDiscount   float64   `gorm:"column:total_discount"`
    TotalPPH23      float64   `gorm:"column:total_pph23"`
    VAT             float64   `gorm:"column:vat"`
    GrandTotal      float64   `gorm:"column:grand_total"`
    DeletedAt       *time.Time `gorm:"column:deleted_at"`
}

func (SalesOrder) TableName() string {
    return "sales_orders"
}
