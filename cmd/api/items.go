//Filename:cmd/api/items.go
package main

import	(
	"fmt"
	"net/http"
	"quiz3.castillojadah.net/internals/data"
	"quiz3.castillojadah.net/internals/validator"
)
//create entry handler for the POST items endpoint
func (app *application) createToDoHandler(w http.ResponseWriter, r *http.Request){
	
	//our target decode destination
	var input struct{
		Name string `json:"name"`
		Description string `json:"description"`
		Status string `json:"status"`
	}

	err := app.readJSON(w, r, &input )
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	entries := &data.Items{
		Name: input.Name,
		Description: input.Description,
		Status: input.Status,
	}
	//initialize a new validator instance
	v := validator.New()
	//check the map to determine if there were any validation errors
	if data.ValidateEntries(v,entries); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)

}