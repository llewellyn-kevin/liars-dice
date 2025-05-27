package handlers

import "fmt"

func layout(filename string) string {
	return fmt.Sprintf("resources/layouts/%s.tmpl.html", filename)
}

func resource(filename string) string {
	return fmt.Sprintf("resources/templates/%s.tmpl.html", filename)
}
