package repository

import "delivery-api/domain/entity"

type ProductRepository interface {
	ExistsByNameAndCategory(name, category string) (bool, error)
	Create(product entity.Product) error
}
