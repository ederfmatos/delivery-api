package product

import (
	"delivery-api/application/events"
	"delivery-api/application/events/products"
	"delivery-api/application/repository"
	"delivery-api/domain/entity"
	"delivery-api/domain/errors"
)

type CreateProductUseCase struct {
	productRepository repository.ProductRepository
	eventEmitter      events.EventEmitter
}

func NewCreateProductUseCase(
	productRepository repository.ProductRepository,
	eventEmitter events.EventEmitter,
) CreateProductUseCase {
	return CreateProductUseCase{
		productRepository: productRepository,
		eventEmitter:      eventEmitter,
	}
}

type Input struct {
	Name     string
	Price    float64
	Category string
}

func (useCase CreateProductUseCase) Execute(input Input) error {
	exists, err := useCase.productRepository.ExistsByNameAndCategory(input.Name, input.Category)
	if err != nil {
		return err
	}
	if exists {
		return errors.ProductAlreadyExistsError()
	}
	product, err := entity.NewProduct(input.Name, input.Price, input.Category)
	if err != nil {
		return err
	}
	err = useCase.productRepository.Create(*product)
	if err != nil {
		return err
	}
	return useCase.eventEmitter.Emit(products.NewProductCreatedEvent(*product))
}
