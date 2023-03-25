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
		assets, err := usecase.Get()
		returnJSON(c, err, assets)
	}
}

func GetAssetPrice(c *gin.Context) {
	code := strings.ToUpper(c.Param("code"))

	usecase := NewConsultAssetPriceUseCase()
	price, err := usecase.Get(code)

	returnJSON(c, err, gin.H{"price": price})
}
