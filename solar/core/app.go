package core

import (
	"fmt"
	"net/http"
)

type Application interface {
	Run()
}

type Rout struct {
	do func()
}

type App struct {
	Request   http.Request
	Response  http.Response
	Router    []Rout
	Component map[string]interface{}
	err       error
}

func (a *App) Run(port string) {
	fmt.Println("App Running!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})
	http.ListenAndServe(port, nil)
}
