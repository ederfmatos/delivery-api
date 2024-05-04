package http

import (
	"delivery-api/presentation/http"
	"encoding/json"
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
	var responseBody interface{}
	httpRequest := NewMuxHttpRequest(request)
	httpResponse, err := controller.HandleRequest(httpRequest)
	if err != nil {
		writer.WriteHeader(422)
		responseBody = map[string]interface{}{"error": err.Error()}
	} else {
		writer.WriteHeader(httpResponse.Status)
		if httpResponse.Body != nil {
			responseBody, _ = json.Marshal(httpResponse.Body)
		}
	}
	responseJson, _ := json.Marshal(responseBody)
	_, _ = writer.Write(responseJson)
}

func (server *MuxHttpServer) Listen(port int32) {
	log.Default().Println("Server running")
	err := http2.ListenAndServe(":"+strconv.Itoa(int(port)), server.mux)
	if err != nil {
		panic(err)
	}
}
