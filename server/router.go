package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func layout(filename string) string {
	return fmt.Sprintf("resources/layouts/%s.tmpl.html", filename)
}

func resource(filename string) string {
	return fmt.Sprintf("resources/templates/%s.tmpl.html", filename)
}

type homeData struct {
	Name   string
	Number int
}

func home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(
		layout("app"),
		resource("home"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = template.Execute(w, &homeData{
		Name:   "Kevin",
		Number: 70,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func Run() {
	homeHandler := http.HandlerFunc(home)

	webStack := makeStack(web, logger)

	http.Handle("/home", webStack(homeHandler))

	fmt.Println("Running server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
