package main

import (
	
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Category{})

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := NewCategoryRepo(db, logger)
	service := NewCategoryService(repo, logger)
	controller := NewCategoryController(service, logger)

	router := gin.Default()

	router.GET("/api/categories", controller.GetAllCategory)
	router.PUT("/api/categories/:id", controller.UpdateCategory)

	router.Run(":8087")
}
