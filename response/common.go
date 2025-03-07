package response

type GeneralRequest struct {
	EventId uint     `form:"event_id" json:"event_id"`
	Seats   []string `form:"seats" json:"seats"`
}

type BookStoreRequest struct {
	Title          string `form:"title" json:"title"`
	NameOfBorrower string `form:"name_of_borrower" json:"name_of_borrower"`
}
