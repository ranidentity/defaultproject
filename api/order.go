package api

import (
	"defaultproject/model"
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/service"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var reqbody response.GeneralRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		c.JSON(500, serializer.ErrResponse(serializer.CodeGeneralError, "", err))
		c.Abort()
		return
	}
	_user, _ := c.Get("user")
	user := _user.(model.User)
	var order service.OrderService
	if res, err := order.AddToCart(user.ID, reqbody.EventId, reqbody.Seats); err == nil {
		c.JSON(200, res)
	} else {
		c.JSON(500, res)
	}
}
