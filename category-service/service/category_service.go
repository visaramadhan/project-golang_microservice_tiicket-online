package service

import (
	"github.com/visaramadhan/project-golang_microservice_tiicket-online/category-service/models"
	"github.com/visaramadhan/project-golang_microservice_tiicket-online/category-service/repository"
	"go.uber.org/zap"
)

type Category = models.Category
type CategoryRepository = repository.CategoryRepository

// Service interface
type CategoryService interface {
	ShowAllCategory(page, limit int) (*[]Category, int, int, error)
	UpdateCategory(categoryID uint, category *Category) error
}

type categoryService struct {
	repo CategoryRepository
	log  *zap.Logger
}

func NewCategoryService(repo CategoryRepository, log *zap.Logger) CategoryService {
	return &categoryService{repo: repo, log: log}
}

func (cs *categoryService) ShowAllCategory(page, limit int) (*[]Category, int, int, error) {
	cs.log.Info("Fetching all category", zap.Int("page", page), zap.Int("limit", limit))

	categories, total, totalPages, err := cs.repo.ShowAllCategory(page, limit)
	if err != nil {
		cs.log.Error("Error fetching categories", zap.Error(err))
		return nil, 0, 0, err
	}

	cs.log.Info("Successfully fetched categories", zap.Int("count", total), zap.Int("totalPages", totalPages))
	return categories, total, totalPages, nil
}

func (cs *categoryService) UpdateCategory(categoryID uint, category *Category) error {
	cs.log.Info("Updating category", zap.Uint("categoryID", categoryID))
	return cs.repo.UpdateCategory(categoryID, category)
}
