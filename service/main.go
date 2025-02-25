package service

import (
	"defaultproject/serializer"
)

func Ping() (r serializer.Response[any], err error) {
	r = serializer.GeneralResponse[any](0, "ping", nil)
	return
}
