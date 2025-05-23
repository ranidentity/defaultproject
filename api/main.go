package api

import (
	"defaultproject/service"

	"github.com/gin-gonic/gin"
)

func ApiPing(c *gin.Context) {
	if res, err := service.Ping(); err == nil {
		c.JSON(200, res)
	} else {
		c.JSON(500, res)
	}
}
