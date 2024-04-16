package main

import (
	"car_catalog/cmd/initializers"
	"car_catalog/controller"
	"car_catalog/internal/config"
	"car_catalog/internal/storage/postgres"

	_ "car_catalog/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.MustLoad()             // read config file
	initializers.HandleLogging()  // init logging
	initializers.OpenConnection() // open db connection
}

// @title           Car Catalog API
// @version         1.0

// @host      localhost:8383
// @BasePath  /api/v1
func main() {
	r := gin.Default()
	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		cars := v1.Group("/cars")
		{
			cars.GET("", c.ListCars)
			cars.POST("stub/:num", postgres.CreateCarStub)
			cars.PATCH(":id", c.UpdateCar)
			cars.DELETE(":id", c.DeleteCar)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.Cfg.Address)
}
