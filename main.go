package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	// Взял fasthttp, тк он один из самых быстрых(в данной задаче скорость не нужна, но привык его юзать).
	server := &fasthttp.Server{
		Handler: handler,
	}

	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal(err)
	}
}
