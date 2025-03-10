package main

import (
	"defaultproject/model"
	"defaultproject/server"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))
	// model.Database(os.Getenv("MYSQL_DSN"))
	model.PostgresDB(os.Getenv("POSTGRESQL_DSN"))
	r := server.NewRouter()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
