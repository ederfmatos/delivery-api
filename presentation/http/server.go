package http

type Server interface {
	Add(controller Controller)
	Listen(port int32)
}
