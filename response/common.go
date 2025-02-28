package response

type GeneralRequest struct {
	EventId uint     `form:"event_id" json:"event_id"`
	Seats   []string `form:"seats" json:"seats"`
}
