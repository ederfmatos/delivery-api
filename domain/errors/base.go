package errors

type DomainError struct {
	Message string
}

func NewDomainError(message string) *DomainError {
	return &DomainError{Message: message}
}

func (domainError *DomainError) Error() string {
	return domainError.Message
}
