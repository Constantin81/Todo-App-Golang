package main

import (
	"log"

	todo "github.com/Constantin81/Todo-app"
	"github.com/Constantin81/Todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s ", err.Error())
	}
}