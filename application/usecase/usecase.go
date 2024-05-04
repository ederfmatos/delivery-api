package usecase

type UseCase[Input any, Output any] interface {
	Execute(input Input) (Output, error)
}

type UnitUseCase[Input any] interface {
	Execute(input Input) error
}
