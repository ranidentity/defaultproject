package model

func (User) TableName() string {
	return "user"
}

type User struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
