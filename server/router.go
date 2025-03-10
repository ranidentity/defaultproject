package server

import (
	"defaultproject/api"
	"defaultproject/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	v1 := r.Group("")
	{
		v1.GET("/ping", api.ApiPing)
		v1.GET("/Book", api.GetBook)
		v1.GET("/Borrow", api.BorrowBook)
		v1.GET("/Extend", api.ExtendLoan)
		v1.GET("/Return", api.ReturnBook)
	}
	return r
}
