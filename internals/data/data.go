// Filename: internal/data/data.go

package data

import (
	"context"
	"database/sql"
	"errors"
	//"fmt"
	"time"

	"quiz3.castillojadah.net/internals/validator"
)

type Items struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`
	Mode string `json:"mode"`
}

func ValidateItems(v *validator.Validator, entries *Items)  {
	//use the check method to execute our validation checks
	v.Check(entries.Name != "", "name", "must be provided")
	v.Check(len(entries.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entries.Description != "", "description", "must be provided")
	v.Check(len(entries.Description ) <= 500, "description", "must not be more than 500 bytes long")

	v.Check(entries.Status != "", "status", "must be provided")
	v.Check(len(entries.Status) <= 200, "status", "must not be more than 200 bytes long")

	v.Check(entries.Mode != "", "mode", "must be provided")
	v.Check(len(entries.Mode ) <= 500, "mode", "must not be more than 500 bytes long")

}
//Define an item model which wraps a sql.DB connection pool
type ItemModel struct {
	DB *sql.DB
}
//Insert allows us to create another to do list
func (m ItemModel) Insert(items *Items) error {
	query := `
		INSERT INTO items (name, description, status, mode)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	// Collect the data fields into a slice
	args := []interface{}{
		items.Name, items.Description,
		items.Status, items.Mode,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&items.ID)
}

//Insert allows us to get another to do list
func (m ItemModel) Get(id int64) (*Items, error) {
	// Ensure that there is a valid id
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Create the query
	query := `
		SELECT pg_sleep(15), id, name, description, status, mode
		FROM items
		WHERE id = $1
	`
	// Declare a items variable to hold the returned data
	var items Items
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query using QueryRow()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&[]byte{},
		&items.ID,
		&items.Name,
		&items.Description,
		&items.Status,
		&items.Mode,
	)
	// Handle any errors
	if err != nil {
		// Check the type of error
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Success
	return &items, nil
}

//Insert allows us to update another to do list
func (m ItemModel) Update(items *Items) error {
	// Create a query
	query := `
		UPDATE items
		SET name = $1, description = $2, status = $3, mode = $4
		WHERE id = $5
	`
	args := []interface{}{
		items.Name,
		items.Description,
		items.Status,
		items.Mode,
		items.ID,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&items.ID)
	// Check for edit conflicts
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

//Insert allows us to delete another to do list
func (m ItemModel) Delete(id int64) error {
	// Ensure that there is a valid id
	if id < 1 {
		return ErrRecordNotFound
	}
	// Create the delete query
	query := `
		DELETE FROM items
		WHERE id = $1
	`
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Check how many rows were affected by the delete operation. We
	// call the RowsAffected() method on the result variable
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check if no rows were affected
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}