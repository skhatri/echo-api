package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/api-router-go/router/functions"
	"github.com/skhatri/api-router-go/router/model"
	"log"
	"net/http"
)

func Echo(request *router.WebRequest) *model.Container {
	container := model.Response(request)
	for k, v := range request.Headers {
		container.AddHeader(k, v)
	}
	printData(container)
	return container
}

func printData(container *model.Container) {
	buff := bytes.Buffer{}
	err := json.NewEncoder(&buff).Encode(container)
	if err == nil {
		fmt.Println(buff.String())
	}
}

func main() {

	httpRouter := router.NewHttpRouterBuilder().
		WithOptions(router.HttpRouterOptions{
			LogRequest: true,
		}).
		Configure(func(configurer router.ApiConfigurer) {
			configurer.Get("/", functions.IndexFunc)
			configurer.Get("/echo", Echo)
			configurer.Post("/echo", Echo)
		}).
		Build()

	var address = "0.0.0.0:8080"
	log.Printf("Listening on %s\n", address)
	http.ListenAndServe(address, httpRouter)

}
