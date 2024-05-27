package http

import (
	"delivery-api/presentation/http"
	"fmt"
	"log"
	http2 "net/http"
	"strconv"
)

type MuxHttpServer struct {
	mux *http2.ServeMux
}

func NewMuxHttpServer() *MuxHttpServer {
	return &MuxHttpServer{
		mux: http2.NewServeMux(),
	}
}

func (server *MuxHttpServer) Add(controller http.Controller) {
	pattern := fmt.Sprintf("%s %s", controller.Method(), controller.Path())
	server.mux.HandleFunc(pattern, func(writer http2.ResponseWriter, request *http2.Request) {
		server.handleRequest(writer, request, controller)
	})
}

func (server *MuxHttpServer) handleRequest(writer http2.ResponseWriter, request *http2.Request, controller http.Controller) {
	httpRequest := NewMuxHttpRequest(request)
	responseWriter := NewMuxResponseWriter(writer)
	err := controller.HandleRequest(httpRequest, responseWriter)
	if err != nil {
		response := map[string]interface{}{"error": err.Error()}
		_ = responseWriter.Status(422).WriteBody(response)
	}
}

func (server *MuxHttpServer) Listen(port int32) {
	log.Default().Println("Server running")
	err := http2.ListenAndServe(":"+strconv.Itoa(int(port)), server.mux)
	if err != nil {
		panic(err)
	}
}
