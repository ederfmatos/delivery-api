package factory

import (
	http2 "delivery-api/main/http"
	"delivery-api/presentation/http"
)

func MakeHttpServer() http.Server {
	server := http2.NewMuxHttpServer()
	controllers := MakeControllers()
	for _, controller := range controllers {
		server.Add(controller)
	}
	return server
}
