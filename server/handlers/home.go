package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type homeData struct {
	Name     string
	Number   int
	Counters []ReactComponent
}

func Home(w http.ResponseWriter, _ *http.Request) {
	template, err := template.ParseFiles(
		layout("app"),
		partial("react-component"),
		resource("home"),
	)

	counters := make([]ReactComponent, 0)
	for i := int32(5); i < 6; i += 4 {
		counter, err := NewReactComponent("counter", CounterProps{StartCount: i})
		if err != nil {
			panic(err)
		}
		counters = append(counters, counter)
	}

	data, err := NewReactPageData(homeData{
		Name:     "Dave",
		Number:   2001,
		Counters: counters,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		fmt.Println("template error", err)
		return
	}
}
