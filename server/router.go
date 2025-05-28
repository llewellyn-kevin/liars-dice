package server

import (
	"fmt"
	"log"
	"net/http"

	"llewellyn.dev/liars_dice/server/handlers"
	"llewellyn.dev/liars_dice/server/vite"
)

func Run(isDev bool) {
    vite.SetDevMode(isDev)

	webStack := makeStack(web, logger)

	http.Handle("/home", webStack(http.HandlerFunc(handlers.Home)))

	fmt.Println("Running server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
