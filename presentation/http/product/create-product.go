package product

import (
	"delivery-api/application/usecase/product"
	"delivery-api/presentation/http"
)

type CreateProductController struct {
	createProductUseCase product.CreateProductUseCase
}

func NewCreateProductController(createProductUseCase product.CreateProductUseCase) http.Controller {
	return CreateProductController{createProductUseCase: createProductUseCase}
}

func (controller CreateProductController) Path() string {
	return "/products"
}

func (controller CreateProductController) Method() http.Method {
	return http.POST
}

func (controller CreateProductController) HandleRequest(request http.Request) (*http.Response, error) {
	input := product.Input{
		Name:     request.BodyFieldString("name"),
		Price:    request.BodyFieldFloat("price"),
		Category: request.BodyFieldString("category"),
	}
	err := controller.createProductUseCase.Execute(input)
	return http.NoContent, err
}
