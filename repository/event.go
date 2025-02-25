package repository

import "defaultproject/model"

type EventRepository struct {
	model.Event
}

func (ref *EventRepository) GetEvent() (result []model.Event) {
	db := model.DB.Model(ref.Event)
	err = db.Find(&result).
		Error
	return
}
