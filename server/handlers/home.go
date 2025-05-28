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

    data, _ := newReactPageData(&homeData{
		Name:   "Michael",
		Number: 2001,
	})

	err = template.Execute(w, data)

	if err != nil {
		fmt.Println(err)
		return
	}
}
