package factory

import (
	"delivery-api/application/usecase/product"
)

type UseCaseFactory struct {
	CreateProductUseCase product.CreateProductUseCase
}

func NewUseCaseFactory(infraFactory InfraFactory) UseCaseFactory {
	return UseCaseFactory{
		CreateProductUseCase: product.NewCreateProductUseCase(infraFactory.ProductRepository, infraFactory.EventEmitter),
	}
}
