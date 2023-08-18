package main

import "net/http"

func (app *applicatiion) routes() *http.ServeMux {
	mux := http.NewServeMux()

	httpFileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", httpFileServer))

	mux.HandleFunc("/", app.homePage)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snipptCreate)

	return mux
}
