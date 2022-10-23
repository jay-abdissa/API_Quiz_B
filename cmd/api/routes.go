//Filename:cmd/api/routes.go
package main

import	(
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func (app *application) routes() *httprouter.Router{
	//Create new httprouter instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed =http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/todo", app.createToDoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/todo/:id", app.showToDoHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/todo/:id", app.updateToDoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/todo/:id", app.deleteToDoHandler)
	return router
}