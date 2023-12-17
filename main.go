package main

import (
	"SocialMediaAPI/datastore"
	"SocialMediaAPI/handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/post/{id}", h.GetByID)
	app.POST("/post", h.Create)
	app.PUT("/post/{id}", h.Update)
	app.DELETE("/post/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
	app.Start()
}
