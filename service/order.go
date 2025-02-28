package service

import (
	"defaultproject/repository"
	"defaultproject/serializer"
)

type OrderService struct {
	repository.OrderRepository
}

// Steps:: add cart with user id
func (ref *OrderService) AddToCart(user_id uint, event_id uint, seats []string) (r serializer.Response, err error) {
	reserved_seats, _ := ref.CartChecker(event_id, seats)
	carts, _ := ref.OrderRepository.AddToCart(user_id, event_id, reserved_seats)
	r = serializer.GeneralResponse(200, "", carts)
	return
}
