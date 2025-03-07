package model

func (Book) TableName() string {
	return "book"
}

type Book struct {
	BaseModel
	Title           string `json:"title"`
	AvailableCopies int    `json:"available_copies"`
}
