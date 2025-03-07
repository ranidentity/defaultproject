package api

import (
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/service"
	"defaultproject/status"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	var reqbody response.BookStoreRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		c.JSON(status.CodeGeneralError, serializer.ErrResponse(status.CodeGeneralError, "", err))
		c.Abort()
		return
	}
	if res, err := service.BookStoreGetBook(reqbody.Title); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}

func Borrow(c *gin.Context) {
	var reqbody response.BookStoreRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		c.JSON(status.CodeGeneralError, serializer.ErrResponse(status.CodeGeneralError, "", err))
		c.Abort()
		return
	}
	if res, err := service.BookStoreGetBook(reqbody.Title); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}
