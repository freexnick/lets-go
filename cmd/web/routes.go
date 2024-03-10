package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("/", dynamic.ThenFunc(app.home))
	mux.Handle("/snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
