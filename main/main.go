package main

import (
	"delivery-api/main/factory"
)

func main() {
	server := factory.MakeHttpServer()
	server.Listen(8080)
}
