package http

import (
	"delivery-api/presentation/http"
	"encoding/json"
	http2 "net/http"
)

type MuxResponseWriter struct {
	Writer http2.ResponseWriter
}

func NewMuxResponseWriter(writer http2.ResponseWriter) *MuxResponseWriter {
	return &MuxResponseWriter{Writer: writer}
}

func (writer *MuxResponseWriter) WriteBody(body interface{}) error {
	responseJson, _ := json.Marshal(body)
	_, err := writer.Writer.Write(responseJson)
	return err
}

func (writer *MuxResponseWriter) StatusOk() http.ResponseWriter {
	writer.Writer.WriteHeader(200)
	return writer
}

func (writer *MuxResponseWriter) StatusNoContent() http.ResponseWriter {
	writer.Writer.WriteHeader(204)
	return writer
}

func (writer *MuxResponseWriter) Status(status int) http.ResponseWriter {
	writer.Writer.WriteHeader(status)
	return writer
}

func (writer *MuxResponseWriter) Header(key string, value string) http.ResponseWriter {
	writer.Writer.Header().Set(key, value)
	return writer
}
