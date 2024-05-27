package http

type ResponseWriter interface {
	WriteBody(body interface{}) error
	StatusOk() ResponseWriter
	StatusNoContent() ResponseWriter
	Status(int) ResponseWriter
	Header(string, string) ResponseWriter
}
