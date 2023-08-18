package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *applicatiion) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applicatiion) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *applicatiion) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
