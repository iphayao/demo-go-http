package main

import (
	"demo-go-http/web"
	"log"
	"net/http"
)

type handler struct {}

func (h *handler) Hello(c web.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}

func main() {

	r := web.New()
	h := handler{}

	r.GET("/hello", h.Hello)

	log.Fatal(r.Start(":8082"))
}
