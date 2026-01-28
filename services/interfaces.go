package services

import "kasir-api/models"

// ProductServiceInterface defines the interface for product service operations
type ProductServiceInterface interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

// CategoryServiceInterface defines the interface for category service operations
type CategoryServiceInterface interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id int) error
}
