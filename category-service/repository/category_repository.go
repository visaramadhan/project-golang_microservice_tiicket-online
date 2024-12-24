package repository

import (
	"fmt"
	"github.com/visaramadhan/project-golang_microservice_tiicket-online/category-service/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
)

type Category = models.Category

// Repository interface
type CategoryRepository interface {
	ShowAllCategory(page, limit int) (*[]Category, int, int, error)
	UpdateCategory(categoryID uint, category *Category) error
}

// Repository implementation
type categoryRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCategoryRepo(db *gorm.DB, log *zap.Logger) CategoryRepository {
	return &categoryRepository{db, log}
}

func (cr *categoryRepository) ShowAllCategory(page, limit int) (*[]Category, int, int, error) {
	cr.log.Info("Fetching all category", zap.Int("page", page), zap.Int("limit", limit))

	var categories []Category
	var totalRecords int64

	// Count total records
	if err := cr.db.Model(&Category{}).Count(&totalRecords).Error; err != nil {
		cr.log.Error("Error counting category", zap.Error(err))
		return nil, 0, 0, err
	}

	// Fetch paginated results
	offset := (page - 1) * limit
	if err := cr.db.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		cr.log.Error("Error fetching categories", zap.Error(err))
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	cr.log.Info("Successfully fetched categories", zap.Int("totalRecords", int(totalRecords)), zap.Int("totalPages", totalPages))
	return &categories, int(totalRecords), totalPages, nil
}

func (cr *categoryRepository) UpdateCategory(categoryID uint, category *Category) error {
	cr.log.Info("Updating category", zap.Uint("categoryID", categoryID))
	result := cr.db.Model(&Category{}).Where("id = ?", categoryID).Updates(map[string]interface{}{
		"name":        category.Name,
		"description": category.Description,
		"icon_url":    category.IconURL,
	})
	if result.Error != nil {
		cr.log.Error("Failed to update category", zap.Uint("categoryID", categoryID), zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		cr.log.Warn("No category found to update", zap.Uint("categoryID", categoryID))
		return fmt.Errorf("no category found with id %d", categoryID)
	}

	cr.log.Info("Successfully updated category", zap.Uint("categoryID", categoryID))
	return nil
}
