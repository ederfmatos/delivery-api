package http

import "context"

type Body interface {
	BodyFieldString(name string) string
	BodyFieldFloat(name string) float64
}

type Request interface {
	Body
	Context() context.Context
}
