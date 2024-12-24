package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/visaramadhan/project-golang_microservice_tiicket-online/category-service/service"
	"go.uber.org/zap"
)

type CategoryService = service.CategoryService

type CategoryController struct {
	service CategoryService
	log     *zap.Logger
}

func NewCategoryController(service CategoryService, log *zap.Logger) *CategoryController {
	return &CategoryController{
		service: service,
		log:     log,
	}
}

func (cc *CategoryController) GetAllCategory(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	categories, total, totalPages, err := cc.service.ShowAllCategory(page, limit)
	if err != nil {
		cc.log.Error("Failed to fetch categories", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories":  categories,
		"total":       total,
		"totalPages":  totalPages,
		"currentPage": page,
	})
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		cc.log.Error("Invalid category ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
		return
	}

	// Placeholder: Update logic here

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}
