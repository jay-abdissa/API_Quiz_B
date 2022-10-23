// Filename: internal/data/data.go

package data

import (
	"time"
	"quiz3.castillojadah.net/internals/validator"
)

type Mystruct struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"createdat"`
	Name string `json:"name"`
	Year string `json:"year"`
	Contact string `json:"contact"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Website string `json:"website"`
	Address string `json:"address"`
}

func ValidateEntries(v *validator.Validator, entries *Mystruct)  {
	//use the check method to execute our validation checks
	v.Check(entries.Name != "", "name", "must be provided")
	v.Check(len(entries.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entries.Year != "", "year", "must be provided")
	v.Check(len(entries.Year) <= 200, "year", "must not be more than 200 bytes long")

	v.Check(entries.Contact != "", "contact", "must be provided")
	v.Check(len(entries.Contact) <= 200, "contact", "must not be more than 200 bytes long")

	v.Check(entries.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(entries.Phone,validator.PhoneRegex), "phone", "must provide a valid phone number")
	v.Check(len(entries.Phone) <= 300, "phone", "must not be more than 300 bytes long")

	v.Check(entries.Email != "", "email", "must be provided")
	v.Check(len(entries.Email) <= 300, "email", "must not be more than 300 bytes long")
	v.Check(validator.Matches(entries.Email,validator.EmailRegex), "phone", "must provide a valid email")

	v.Check(entries.Website != "", "website", "must be provided")
	v.Check(len(entries.Website) <= 200, "website", "must not be more than 200 bytes long")
	v.Check(validator.ValidWebsite(entries.Website), "website", "must be a valid URL")

	v.Check(entries.Address != "", "Address", "must be provided")
	v.Check(len(entries.Address) <= 200, "Address", "must not be more than 200 bytes long")
	
}