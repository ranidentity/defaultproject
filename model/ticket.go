package model

type TicketCategory struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}
