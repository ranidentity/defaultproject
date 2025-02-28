package api

import (
	"defaultproject/model"
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/service"
	"defaultproject/status"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var reqbody response.GeneralRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		c.JSON(status.CodeGeneralError, serializer.ErrResponse(status.CodeGeneralError, "", err))
		c.Abort()
		return
	}
	_user, _ := c.Get("user")
	user := _user.(model.User)
	var order service.OrderService
	if res, err := order.AddToCart(user.ID, reqbody.EventId, reqbody.Seats); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}

func GetCart(c *gin.Context) {
	_user, _ := c.Get("user")
	user := _user.(model.User)
	var order service.OrderService
	if res, err := order.GetUserCart(user.ID); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}
