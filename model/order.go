package model

import "github.com/shopspring/decimal"

// Order will check out cart
func (Order) TableName() string {
	return "order"
}

type Order struct {
	BaseModel
	CartId      uint            `json:"cart_id"`
	Status      int8            `json:"status"`
	TotalAmount decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"TotalAmount"`
	FinalAmount decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"FinalAmount"`
}

func (Cart) TableName() string {
	return "cart"
}

type Cart struct {
	BaseModel
	EventTicketId uint            `json:"event_ticket_id"` //event_ticket
	Price         decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"price"`
	NumberOfItem  int             `json:"number_of_item"`
	TotalAmount   decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"TotalAmount"`
}

func (Payment) TableName() string {
	return "payment"
}

// only 1 payment per order
type Payment struct {
	BaseModel
	Status int8            `json:"status"`
	Amount decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"amount"`
}
