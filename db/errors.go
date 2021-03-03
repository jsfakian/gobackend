package db

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrorStruct struct{}

var Error = ErrorStruct{}

//ErrNotFound is an function that checks if an object is not found in the db
func (e ErrorStruct) ErrNotFound(err error, what string) error {
	return wrapDBErr(err, "Could not find "+what)
}

//ErrUpdate is an function that checks if there is an error updating an object in the db
func (e ErrorStruct) ErrUpdate(err error, what string) error {
	if err != nil {
		fmt.Println(what, " ", err)
	}
	return wrapDBErr(err, "Could not update "+what)
}

//ErrDelete is an function that checks if there is an error deleting an object from the db
func (e ErrorStruct) ErrDelete(err error, what string) error {
	if err != nil {
		fmt.Println(what, " ", err)
	}
	return wrapDBErr(err, "Could not delete "+what)
}

//ErrCreate is an function that checks if there is an error creating an object in the db
func (e ErrorStruct) ErrCreate(err error, what string) error {
	if err != nil {
		fmt.Println(what, " ", err)
	}
	return wrapDBErr(err, "Could not create "+what)
}

//ErrIsEmpty is an function that checks if the return from a search query of db is empty
func (e ErrorStruct) ErrIsEmpty(err error, whats string) error {
	return wrapDBErr(err, "List of "+whats+" is empty.")
}

func wrapDBErr(err error, message string) error {
	if err != nil {
		return errors.Wrap(err, "DB: "+message)
	}
	return err
}
