// Filename: internal/data/data.go

package data

import (
	"quiz3.castillojadah.net/internals/validator"
)

type Items struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`
}

func ValidateEntries(v *validator.Validator, entries *Items)  {
	//use the check method to execute our validation checks
	v.Check(entries.Name != "", "name", "must be provided")
	v.Check(len(entries.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entries.Description != "", "description ", "must be provided")
	v.Check(len(entries.Description ) <= 200, "description ", "must not be more than 200 bytes long")

	v.Check(entries.Status != "", "status", "must be provided")
	v.Check(len(entries.Status) <= 200, "status", "must not be more than 200 bytes long")

}