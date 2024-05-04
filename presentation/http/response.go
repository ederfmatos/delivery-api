package http

type Response struct {
	Status  int
	Body    *any
	Headers map[string]string
}

var NoContent = &Response{
	Status:  204,
	Body:    nil,
	Headers: nil,
}
