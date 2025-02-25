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
	}
	return r
}
