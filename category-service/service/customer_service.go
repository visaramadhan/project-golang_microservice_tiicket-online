package service

import (
	"go.uber.org/zap"
)

type CategoryService interface {
	ShowAllCategory(page, limit int) (*[]model.Category, int, int, error)
	UpdateCategory(categoryID uint, category *model.Category) error
}

type categoryService struct {
	repo *repository.AllRepository
	log  *zap.Logger
}

func NewCategoryService(repo *repository.AllRepository, log *zap.Logger) CategoryService {
	return &categoryService{repo: repo, log: log}
}

func (cs *categoryService) ShowAllCategory(page, limit int) (*[]model.Category, int, int, error) {
	cs.log.Info("Fetching all category", zap.Int("page", page), zap.Int("limit", limit))

	category, count, totalPages, err := cs.repo.Category.ShowAllCategory(page, limit)
	if err != nil {
		cs.log.Error("Error fetching category", zap.Error(err))
		return nil, 0, 0, err
	}

	cs.log.Info("Successfully fetched category", zap.Int("count", count), zap.Int("totalPages", totalPages))
	return category, count, totalPages, nil
}

func (ps *categoryService) UpdateCategory(categoryID uint, category *model.Category) error {
	ps.log.Info("Updating category", zap.Uint("categoryID", categoryID), zap.String("name", category.Name))

	if err := ps.repo.Category.UpdateCategory(categoryID, category); err != nil {
		ps.log.Error("Error updating category", zap.Error(err))
		return err
	}

	ps.log.Info("Successfully updated category", zap.Uint("categoryID", categoryID))
	return nil
}
