package repositories

import (
	"gorm.io/gorm"
)

type SalesOrderRepository interface {
    IndexV2ParentBuilder() *gorm.DB
}

type salesOrderRepository struct {
    db *gorm.DB
}

func NewSalesOrderRepository(db *gorm.DB) SalesOrderRepository {
    return &salesOrderRepository{db}
}

func (r *salesOrderRepository) IndexV2ParentBuilder() *gorm.DB {
    var query = r.db.Table("sales_orders").
        Joins("LEFT JOIN sales_has_styles ON sales_orders.id = sales_has_styles.sales_order_id").
        Joins("JOIN customers ON sales_orders.customer_id = customers.id").
        Joins("JOIN order_types ON sales_orders.order_type_id = order_types.id").
        Joins("JOIN currencies ON sales_orders.currency_id = currencies.id").
        Joins("LEFT JOIN users ON sales_orders.created_by = users.id").
        Joins("LEFT JOIN users as updated_by ON sales_orders.updated_by_id = updated_by.id").
        Joins("LEFT JOIN users as validator_1_user ON sales_orders.validator_1_id = validator_1_user.id").
        Joins("LEFT JOIN styles ON sales_has_styles.style_id = styles.id").
        Joins("LEFT JOIN sales_order_and_style_has_color_methods ON sales_orders.id = sales_order_and_style_has_color_methods.sales_order_id AND styles.id = sales_order_and_style_has_color_methods.style_id").
        Joins("LEFT JOIN color_methods ON sales_order_and_style_has_color_methods.color_method_id = color_methods.id").
        Joins("LEFT JOIN selected_color_variants ON sales_orders.id = selected_color_variants.sales_order_id AND styles.id = selected_color_variants.style_id AND color_methods.id = selected_color_variants.color_method_id").
        Joins("LEFT JOIN color_variants ON selected_color_variants.color_variant_id = color_variants.id").
        Joins("LEFT JOIN coloring_so ON sales_has_styles.id = coloring_so.sales_has_style_id").
        Joins("LEFT JOIN group_methods ON coloring_so.group_method_id = group_methods.id").
        Joins("LEFT JOIN color_groups ON group_methods.group_id = color_groups.id").
        Joins("LEFT JOIN styles as cg_styles ON color_groups.style_id = cg_styles.id").
        Joins("LEFT JOIN color_methods as gm_color_methods ON group_methods.color_method_id = gm_color_methods.id").
        Joins("LEFT JOIN color_variants as gm_color_variants ON group_methods.color_variant_id = gm_color_variants.id").
        Where("sales_orders.deleted_at IS NULL").
        Group("sales_orders.id")

    return query
}
