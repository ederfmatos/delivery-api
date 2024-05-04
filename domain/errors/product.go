package errors

func ProductAlreadyExistsError() error {
	return NewDomainError("Produto já existente")
}
