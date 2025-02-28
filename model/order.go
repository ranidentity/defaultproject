package model

import "github.com/shopspring/decimal"

const (
	// Shared
	Pending = 0

	//event
	Scheduled    = 1
	PastComplete = -1
	Delayed      = -2
	Cancelled    = -99

	// location
	Active               = 1
	TemporaryUnavailable = -1
	Closed               = -99

	// FloorPlan
	SeatOpen        = 1
	SeatMaintenance = -1
	SeatClose       = -2

	// order
	Complete       = 2
	PayingStage    = 1
	OrderCancelled = -99

	// payment
	Paid   = 1
	Failed = -99
)

// Order will check out cart
func (Order) TableName() string {
	return "order"
}

type Order struct {
	BaseModel
	UserId      uint            `json:"user_id"`
	CartId      uint            `json:"cart_id"`
	Status      int8            `json:"status"`
	TotalAmount decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"TotalAmount"`
	FinalAmount decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"FinalAmount"`
}

func (Cart) TableName() string {
	return "cart"
}

// depend on type of item sold in this system
type CartItem struct {
	EventId       uint   `json:"event_id"`
	SeatNo        string `json:"seat_no"`
	EventTicketId uint   `json:"event_ticket_id"`
}
type Cart struct {
	BaseModel
	CartItem
	Price        decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"price"`
	NumberOfItem int             `json:"number_of_item"`
	TotalAmount  decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"TotalAmount"`
	UserId       uint            `json:"user_id"`
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
