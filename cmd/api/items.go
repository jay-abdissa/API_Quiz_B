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
	//create a to do list
	err = app.models.Items.Insert(items)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// Create a Location header for the newly created resource/todo
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/todo/%d", items.ID))
	// Write the JSON response with 201 - Created status code with the body
	// being the item data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"items": items}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}