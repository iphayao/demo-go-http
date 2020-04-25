package main

import (
	"demo-go-http/web"
	"log"
	"net/http"
)

type handler struct {}

type Message struct {
	Message string
}

func (h *handler) Hello(c *web.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}

func (h *handler) Post(c *web.Context) error {
	return c.JSON(http.StatusOK, "Post Handled")
}

func (h *handler) GetById(c *web.Context) error {
	log.Println(c.Params["id"])
	return c.JSON(http.StatusOK, Message{"Get By Id Handled"})
}

func main() {

	r := web.New()
	h := handler{}

	//r.Use(middleware.Logger())

	r.GET("/hello", h.Hello)
	r.GET("/hello/:id", h.GetById)
	r.POST("/hello", h.Post)

	log.Fatal(r.Start(":8082"))

}







