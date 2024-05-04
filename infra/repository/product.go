package repository

import (
	"delivery-api/domain/entity"
	"delivery-api/infra/repository/orm"
	"gorm.io/gorm"
)

type DefaultProductRepository struct {
	Database *gorm.DB
}

func NewDefaultProductRepository(database *gorm.DB) *DefaultProductRepository {
	return &DefaultProductRepository{Database: database}
}

func (repository *DefaultProductRepository) ExistsByNameAndCategory(name, category string) (bool, error) {
	var productORM orm.ProductORM
	err := repository.Database.First(&productORM, "name = ? AND category = ?", name, category).Error
	if err != nil && err.Error() != "record not found" {
		return false, err
	}
	return productORM.ID != "", nil
}

func (repository *DefaultProductRepository) Create(product entity.Product) error {
	productORM := orm.FromProduct(product)
	return repository.Database.Create(&productORM).Error
}
