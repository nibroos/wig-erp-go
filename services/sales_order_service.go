package services

import (
	"derp_api_go/models"
	"derp_api_go/repositories"
)

type SalesOrderService interface {
    GetSalesOrders(isAll bool) ([]models.SalesOrder, error)
}

type salesOrderService struct {
    repo repositories.SalesOrderRepository
}

func NewSalesOrderService(repo repositories.SalesOrderRepository) SalesOrderService {
    return &salesOrderService{repo}
}

func (s *salesOrderService) GetSalesOrders(isAll bool) ([]models.SalesOrder, error) {
    var orders []models.SalesOrder

    query := s.repo.IndexV2ParentBuilder()

    if isAll {
        query.Select("sales_orders.id, sales_orders.total_amount, sales_orders.total_discount, sales_orders.status, sales_orders.total_pph23, sales_orders.vat, sales_orders.grand_total")
    } else {
        query.Select("sales_orders.id, sales_orders.order_number, sales_orders.buyer_po_number, sales_orders.customer_id, sales_orders.order_date, sales_orders.shipping_date, sales_orders.production_date, sales_orders.status, sales_orders.currency_id, sales_orders.validator_1_id, sales_orders.validated_1_at, sales_orders.created_by as created_by_id, users.name as created_by_name, updated_by.name as updated_by_name, customers.name as customer_name, currencies.symbol as currency_symbol, validator_1_user.name as validator_1_name, sales_orders.total_amount, sales_orders.total_discount, sales_orders.total_pph23, sales_orders.vat, sales_orders.grand_total, sales_orders.total_qty as qty, sales_orders.total_tmp_high_amount as total_high_amount").Limit(10).Offset(0)
    }

    if err := query.Find(&orders).Error; err != nil {
        return nil, err
    }

    return orders, nil
}
