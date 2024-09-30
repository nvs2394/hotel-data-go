package main

import (
	"hotel-data-go/pkg/api"
	"hotel-data-go/pkg/repository"
	"hotel-data-go/pkg/service"
	"hotel-data-go/utils"
	"log"

	_ "hotel-data-go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	utils.LoadEnv()

	db, err := utils.ConnectDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	hotelRepo := repository.NewHotelRepository(db)

	hotelService := service.NewHotelService(hotelRepo)

	hotelHandler := api.NewHotelHandler(hotelService)

	r := gin.Default()

	// API documentation
	if utils.GetAppEnv() != "prd" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// API routes
	r.GET("/hotels", hotelHandler.GetHotels)

	if err := r.Run(("localhost:8080")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
