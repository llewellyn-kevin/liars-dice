package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type homeData struct {
	Name   string
	Number int
}

func Home(w http.ResponseWriter, _ *http.Request) {
	template, err := template.ParseFiles(
		layout("app"),
		resource("home"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = template.Execute(w, &homeData{
		Name:   "Anna",
		Number: 24,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}
