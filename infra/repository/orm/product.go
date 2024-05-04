package orm

import (
	"delivery-api/domain/entity"
)

type ProductORM struct {
	ID       string  `gorm:"primaryKey"`
	Name     string  `gorm:"name,omitempty"`
	Price    float64 `gorm:"price,omitempty"`
	Category string  `gorm:"category,omitempty"`
}

func FromProduct(product entity.Product) ProductORM {
	return ProductORM{
		ID:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
	}
}

func (orm ProductORM) TableName() string {
	return "products"
}
