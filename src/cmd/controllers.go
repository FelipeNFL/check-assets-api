package cmd

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func returnJSON(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func HealthCheckController(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func CreateNewAssetController(database *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		asset := CreateNewAssetDTO{}

		c.BindJSON(&asset)

		asset.Code = strings.ToUpper(asset.Code)

		usecase := NewCreateAssetUseCase(database)
		assetInserted, err := usecase.Create(asset.Code)

		returnJSON(c, err, assetInserted)
	}
}

func GetAllAssetsController(database *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		usecase := NewGetAssetListUseCase(database)

		order := strings.ToLower(c.Query("order"))
		assets, err := usecase.Get(order)
		returnJSON(c, err, assets)
	}
}

func GetAssetPriceController(c *gin.Context) {
	codes := strings.ToUpper(c.Query("code"))

	usecase := NewConsultAssetPriceUseCase()
	prices, err := usecase.Get(strings.Split(codes, ","))

	returnJSON(c, err, prices)
}

func SaveAssetOrdinationController(database *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := SaveAssetOrdinationDTO{}

		c.BindJSON(&body)

		usecase := NewSaveAssetOrdinationUseCase(database)
		saved, err := usecase.Save(body.Ordination, body.CustomOrder)
		returnJSON(c, err, saved)
	}
}
