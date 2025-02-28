package repository

import (
	"defaultproject/cache"
	"defaultproject/model"
)

type OrderRepository struct {
	model.Cart
	model.EventTicket
}

func (ref *OrderRepository) CartChecker(user_id uint) (cart []model.Cart, err error) {
	db := model.DB.Model(ref.Cart)
	db.Where("user_id = ?", user_id)
	err = db.Find(&cart).
		Error
	return
}

// TODO add reserved seat to redis
func (ref *OrderRepository) ItemChecker(event_id uint, seat_no []string) (open_seats []model.EventTicket, err error) {
	var list []model.EventTicket
	db := model.DB.Model(ref.EventTicket)
	db.Where("event_id = ", event_id).
		Where("seat_no in ?", seat_no)
	err = db.Find(&list).
		Error
	for _, i := range list {
		flag, _ := cache.RedisClient.SIsMember(cache.RedisKey["RESERVED_SEAT"].Name, i).Result()
		if !flag && i.Status == model.SeatOpen {
			open_seats = append(open_seats, i)
			cache.RedisClient.SAdd(cache.RedisKey["RESERVED_SEAT"].Name, i)
		}
	}
	return
}

func (ref *OrderRepository) AddToCart(user_id uint, event_id uint, selected_tickets []model.EventTicket) (cart model.Cart, err error) {
	var carts []model.Cart
	for _, i := range selected_tickets {
		ea := model.Cart{
			CartItem:     model.CartItem{EventId: event_id, SeatNo: i.SeatNo, EventTicketId: 0},
			UserId:       user_id,
			Price:        i.Price,
			NumberOfItem: 1,
			TotalAmount:  i.Price,
		}
		carts = append(carts, ea)
	}
	db := model.DB
	tx := db.Begin()
	err = tx.Create(carts).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	// db.Transaction(func(tx *gorm.DB) error {
	// 	if err = tx.Create(carts).Error; err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
	return
}
