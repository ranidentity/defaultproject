package model

import (
	"time"

	"github.com/shopspring/decimal"
)

func (Event) TableName() string {
	return "event"
}

type Event struct {
	BaseModel
	Name           string    `json:"name"`
	ShortName      string    `json:"short_name"`
	Location       string    `json:"location"`
	StartDateTime  time.Time `json:"start_date"`
	EndDateTime    time.Time `json:"end_date"`
	TicketingStart time.Time `json:"ticketing_start"`
	TicketingEnd   time.Time `json:"ticketing_end"`
	Slot           int       `json:"slot"`   // total tickets
	Status         int8      `json:"status"` // use main/code.go
}

func (EventLocation) TableName() string {
	return "event_location"
}

type EventLocation struct {
	BaseModel
	EventId       uint      `json:"event_id"`
	LocationId    uint      `json:"location_id"`
	StartDateTime time.Time `json:"start_date"`
	EndDateTime   time.Time `json:"end_date"`
	Status        int8      `json:"status"`
}

func (EventTicket) TableName() string {
	return "event_ticket"
}

type EventTicket struct {
	BaseModel
	EventId          uint            `json:"event_id"`
	TicketCategoryId uint            `json:"ticket_category_id"`
	Count            int             `json:"count"`
	Price            decimal.Decimal `gorm:"type:DECIMAL(10,2);not null" json:"price"`
	Status           int8            `json:"status"`
}
