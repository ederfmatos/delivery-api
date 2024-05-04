package http

type Method string

const (
	POST Method = "POST"
)

type Controller interface {
	Path() string
	Method() Method
	HandleRequest(request Request) (*Response, error)
}
