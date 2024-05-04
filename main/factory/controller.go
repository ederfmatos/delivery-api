package factory

import (
	"delivery-api/presentation/http"
	"delivery-api/presentation/http/product"
)

func MakeControllers() []http.Controller {
	infraFactory := NewInfrastructureFactory()
	useCaseFactory := NewUseCaseFactory(infraFactory)
	return []http.Controller{
		product.NewCreateProductController(useCaseFactory.CreateProductUseCase),
	}
}
