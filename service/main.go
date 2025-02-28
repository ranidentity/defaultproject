package service

import (
	"defaultproject/serializer"
)

func Ping() (r serializer.Response, err error) {
	r = serializer.GeneralResponse("ping", nil)
	return
}
