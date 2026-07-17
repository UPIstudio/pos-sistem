package models

import "time"

type Transaction struct {
	ID           uint              `json:"id" gorm:"primarykey"`
	MemberID     *uint             `json:"member_id"`
	Member       *Member           `json:"member"`
	TotalAmount  float64           `json:"total_amount"`
	Discount     float64           `json:"discount"`
	GrandTotal   float64           `json:"grand_total"`
	Status       string            `json:"status"`
	CustomerName string            `json:"customer_name"`
	Items        []TransactionItem `json:"items" gorm:"foreignKey:TransactionID"`
	CreatedAt    time.Time         `json:"created_at"`
}

type TransactionItem struct {
	ID            uint    `json:"id" gorm:"primarykey"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Quantity      int     `json:"quantity"`
	PriceAtSold   float64 `json:"price_at_sold"`
}
