package main

import (
	"github.com/FelipeNFL/check-assets-api/cmd"
	"github.com/gin-gonic/gin"
)

func main() {
	DATABASE_NAME := "api"

	database := cmd.GetMongoDatabase(DATABASE_NAME)

	r := gin.Default()
	r.GET("/health", cmd.HealthCheckController)
	r.POST("/asset", cmd.CreateNewAssetController(database))
	r.Run()
}
