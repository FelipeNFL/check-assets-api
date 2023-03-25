package main

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeNFL/check-assets-api/cmd"
	"github.com/FelipeNFL/check-assets-api/commom"
)

const DATABASE_NAME = "api"

func main() {
	database := commom.GetMongoDatabase(DATABASE_NAME)

	r := gin.Default()
	r.GET("/health", cmd.HealthCheckController)
	r.POST("/asset", cmd.CreateNewAssetController(database))
	r.GET("/asset", cmd.GetAllAssetsController(database))
	r.GET("/asset/:code/price", cmd.GetAssetPrice)
	r.Run()
}
