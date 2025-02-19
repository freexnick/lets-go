package main

import (
	"lets-go-snippetbox/ui"
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(ui.Files))

	mux.Handle("/static/{filepath...}", fileServer)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.HandleFunc("/ping", ping)

	mux.Handle("/", dynamic.ThenFunc(app.home))

	mux.Handle("/snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("/user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("/user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("/about", dynamic.ThenFunc(app.about))
	mux.Handle("/account/view", dynamic.ThenFunc(app.accountView))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("/snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))
	mux.Handle("/account/password/update", protected.ThenFunc(app.accountPasswordUpdate))
	mux.Handle("POST /account/password/update", protected.ThenFunc(app.accountPasswordUpdatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
