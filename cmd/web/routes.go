package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/for-students", http.HandlerFunc(app.forStudents))
	mux.Get("/for-staff", http.HandlerFunc(app.forStaff))
	mux.Get("/for-applicants", http.HandlerFunc(app.forApplicants))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet)) // Moved down
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}