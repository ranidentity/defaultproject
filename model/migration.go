package model

func migration() {
	_ = DB.AutoMigrate(&Event{}, &EventLocation{}, &EventTicket{})
}
