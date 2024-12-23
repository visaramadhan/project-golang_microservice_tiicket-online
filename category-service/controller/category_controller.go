package controller

import (
	// "net/http"
	// "ticket-online/helper"
	"strconv"
	"sync"
	"ticket-online/service"
	"ticket-onlinep/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController struct {
	service *service.AllService
	log     *zap.Logger
}

func NewCategoryController(service *service.AllService, log *zap.Logger) *CategoryController {
	return &CategoryController{
		service: service,
		log:     log,
	}
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Get a list of categories with optional pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} model.SuccessResponse{data=[]model.Category} "List of categories retrieved successfully"
// @Failure 500 {object} model.ErrorResponse "Failed to fetch categories"
// @Router /api/categories [get]
func (cc *CategoryController) GetAllCategory(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	categories, total, totalPages, err := cc.service.Category.ShowAllCategory(page, limit)
	if err != nil {
		cc.log.Error("Failed to fetch categories", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to fetch categories", nil)
		return
	}

	response := gin.H{
		"categories":  categories,
		"total":       total,
		"totalPages":  totalPages,
		"currentPage": page,
	}
	helper.Responses(c, http.StatusOK, "Categories retrieved successfully", response)
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Update the details of a category by its ID
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Category ID"
// @Param name formData string false "Category Name"
// @Param description formData string false "Category Description"
// @Param icon formData file false "Category Icon"
// @Success 200 {object} model.SuccessResponse{data=model.Category} "Category updated successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid category ID or data"
// @Failure 404 {object} model.ErrorResponse "Category not found"
// @Failure 500 {object} model.ErrorResponse "Failed to update category"
// @Router /api/categories/{id} [put]
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		cc.log.Error("Invalid category ID", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	// Fetch existing category
	category, err := cc.service.Category.GetCategoryByID(id)
	if err != nil {
		cc.log.Error("Category not found", zap.Error(err))
		helper.Responses(c, http.StatusNotFound, "Category not found", nil)
		return
	}

	// Read form data
	form, err := c.MultipartForm()
	if err != nil {
		cc.log.Error("Error reading form data", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "Invalid form data: "+err.Error(), nil)
		return
	}

	// Update file if provided
	files := form.File["icon"]
	if len(files) > 0 {
		var wg sync.WaitGroup
		responses, err := helper.Upload(&wg, files)
		if err != nil {
			cc.log.Error("Failed to upload icon", zap.Error(err))
			helper.Responses(c, http.StatusInternalServerError, "Failed to upload icon", nil)
			return
		}
		category.IconURL = responses[0].Data.Url
		cc.log.Info("Icon updated successfully", zap.String("iconURL", responses[0].Data.Url))
	}

	// Update other fields if provided
	if name := c.PostForm("name"); name != "" {
		category.Name = name
	}
	if description := c.PostForm("description"); description != "" {
		category.Description = description
	}

	// Save updates to database
	if err := cc.service.Category.UpdateCategory(category.ID, category); err != nil {
		cc.log.Error("Failed to update category", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to update category", nil)
		return
	}

	cc.log.Info("Category updated successfully", zap.String("categoryName", category.Name))
	helper.Responses(c, http.StatusOK, "Category updated successfully", category)
}
