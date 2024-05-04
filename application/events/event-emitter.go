package events

type Event interface {
	Id() string
	Name() string
}

type EventEmitter interface {
	Emit(event Event) error
}
