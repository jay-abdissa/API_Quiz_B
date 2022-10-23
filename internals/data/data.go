// Filename: internal/data/data.go

package data

import (
	"database/sql"

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
	v.Check(len(entries.Description ) <= 500, "description ", "must not be more than 500 bytes long")

	v.Check(entries.Status != "", "status", "must be provided")
	v.Check(len(entries.Status) <= 200, "status", "must not be more than 200 bytes long")

}
//Define an item model which wraps a sql.DB connection pool
type ItemModel struct {
	DB *sql.DB
}
//Insert allows us to create another to do list
func (m ItemModel) Insert(items *Items) error {
	query := `
		INSERT INTO items (name, decription, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	// Collect the data fields into a slice
	args := []interface{}{
		items.Name, items.Description,
		items.Status,
	}
	return m.DB.QueryRowContext(query, args...).Scan(&items.ID)
}

//Insert allows us to get another to do list
func (m ItemModel) Get(id int64) (*Items, error) {
	return nil, nil
}

//Insert allows us to update another to do list
func (m ItemModel) Update(items *Items) error {
	return nil
}

//Insert allows us to delete another to do list
func (m ItemModel) Delete(id int64) error {
	return nil
}