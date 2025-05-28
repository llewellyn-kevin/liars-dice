package handlers

import (
	"encoding/json"

	"github.com/olivere/vite"
	appVite "llewellyn.dev/liars_dice/server/vite"
)

type pageProps = interface{}

type pageData struct {
	Vite  *vite.Fragment
	Props pageProps
}

func NewPageData(props pageProps) *pageData {
	return &pageData{
		Vite:  nil,
		Props: props,
	}
}

func NewReactPageData(props pageProps) (*pageData, error) {
	viteTags, err := appVite.GetViteFragment()
	if err != nil {
		return &pageData{}, err
	}

	return &pageData{
		Vite:  viteTags,
		Props: props,
	}, nil
}

type ReactComponent struct {
	Name string
	Data string
}

func NewReactComponent(name string, data interface{}) (ReactComponent, error) {
	dataString, err := json.Marshal(data)
	if err != nil {
		return ReactComponent{}, err
	}

	return ReactComponent{
		Name: name,
		Data: string(dataString),
	}, nil
}
