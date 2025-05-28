package handlers

import (
	"github.com/olivere/vite"
	appVite "llewellyn.dev/liars_dice/server/vite"
)

type pageProps interface{}

type pageData struct {
	Vite  *vite.Fragment
	Props pageProps
}

func newPageData(props pageProps) *pageData {
	return &pageData{
		Vite:  nil,
		Props: props,
	}
}

func newReactPageData(props pageProps) (*pageData, error) {
	viteTags, err := appVite.GetViteFragment()
	if err != nil {
		return &pageData{}, err
	}

	return &pageData{
		Vite:  viteTags,
		Props: props,
	}, nil
}
