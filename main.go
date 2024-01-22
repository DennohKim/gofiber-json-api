package main

import (
	"flag"

	"github.com/DennohKim/gofiber-json-api/api"
	"github.com/gofiber/fiber/v2"
)

func main(){
	listenAddr := flag.String("listenAddr", ":5000", "API server listen address")
	flag.Parse()
	app:= fiber.New()

	apiv1 := app.Group("api/v1");

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}


