package vite

import (
	"os"

	"github.com/olivere/vite"
)

var devFlagSingleton bool = false

func SetDevMode(devMode bool) {
	devFlagSingleton = devMode
}

func GetViteFragment() (*vite.Fragment, error) {
	if devFlagSingleton {
		return vite.HTMLFragment(vite.Config{
			FS:      os.DirFS("."),
			IsDev:   true,
			ViteURL: "http://localhost:5173",
		})
	}
	return vite.HTMLFragment(vite.Config{
		FS:    os.DirFS("dist"),
		IsDev: false,
	})
}
