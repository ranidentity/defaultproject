package model

func (Location) TableName() string {
	return "location"
}

type Location struct {
	BaseModel
	Name      string  `json:"name"`
	ShortName string  `json:"short_name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Seat      int     `json:"seat"`
	Status    int8    `json:"status"`
}

// every location has a floor plan
type LocationFloorPlan struct {
	BaseModel
	LocationId uint   `json:"location_id"`
	SeatNo     string `json:"seat_no"`
	Status     int8   `json:"status"`
}
