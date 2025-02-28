package service

import (
	"defaultproject/repository"
	"defaultproject/serializer"
)

type OrderService struct {
	repository.OrderRepository
}

// TODO remove invalid items
func (ref *OrderService) GetUserCart(user_id uint) (r serializer.Response, err error) {
	carts, _ := ref.GetUserCart(user_id)
	r = serializer.GeneralResponse("", carts)
	return
}

// Step:: add cart with user id
func (ref *OrderService) AddToCart(user_id uint, event_id uint, seats []string) (r serializer.Response, err error) {
	reserved_seats, _ := ref.ItemChecker(event_id, seats)
	carts, _ := ref.OrderRepository.AddToCart(user_id, event_id, reserved_seats)
	r = serializer.GeneralResponse("", carts)
	return
}
