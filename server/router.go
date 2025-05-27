package server

import (
	"fmt"
	"log"
	"net/http"

	"llewellyn.dev/liars_dice/server/handlers"
)

func Run() {
	webStack := makeStack(web, logger)

	http.Handle("/home", webStack(http.HandlerFunc(handlers.Home)))

	fmt.Println("Running server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
