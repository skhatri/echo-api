package main

import (
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/api-router-go/router/functions"
	"log"
	"net/http"
)

func main() {

	httpRouter := router.NewHttpRouterBuilder().
		WithOptions(router.HttpRouterOptions{
			LogRequest: true,
		}).
		Configure(func(configurer router.ApiConfigurer) {
			configurer.Get("/", functions.IndexFunc)
			configurer.Get("/echo", functions.EchoFunc)
			configurer.Post("/echo", functions.EchoFunc)
		}).
		Build()

	var address = "0.0.0.0:8080"
	log.Printf("Listening on %s\n", address)
	http.ListenAndServe(address, httpRouter)

}
