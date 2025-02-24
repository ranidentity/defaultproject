package main

import (
	"defaultproject/server"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
