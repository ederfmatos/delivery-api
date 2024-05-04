package products

import "delivery-api/domain/entity"

type ProductCreatedEvent struct {
	ProductId string
}

func NewProductCreatedEvent(product entity.Product) *ProductCreatedEvent {
	return &ProductCreatedEvent{ProductId: product.Id}
}

func (event *ProductCreatedEvent) Id() string {
	return event.ProductId
}

func (event *ProductCreatedEvent) Name() string {
	return "PRODUCT_CREATED"
}
