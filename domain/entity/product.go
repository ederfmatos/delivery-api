package entity

import "fmt"

type Product struct {
	Id       string
	Name     string
	Price    float64
	Category string
}

func NewProduct(name string, price float64, category string) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("O nome do produto é obrigatório")
	}
	if category == "" {
		return nil, fmt.Errorf("A categoria do produto é obrigatório")
	}
	if price <= 0 {
		return nil, fmt.Errorf("O valor do produto deve ser maior que 0,00")
	}
	return &Product{
		Id:       NewId(),
		Name:     name,
		Price:    price,
		Category: category,
	}, nil
}
