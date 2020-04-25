package middleware

import (
	"demo-go-http/web"
	"log"
)

func Logger() web.HandlerFunc {
	return func(c *web.Context) error {
		log.Println("In Logger middleware, Method: " + c.Method)
		return nil
	}
}