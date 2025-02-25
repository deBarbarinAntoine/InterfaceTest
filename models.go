package main

import (
	"InterfaceTest/api"
	"InterfaceTest/db"
	"InterfaceTest/internal"
)

// UserCRUD interface defines methods for CRUD operations on users.
type UserCRUD interface {
	// GetByID retrieves a user by their ID.
	GetByID(id uint) (*internal.User, error)
	
	// GetByUsername retrieves a user by their username.
	GetByUsername(username string) (*internal.User, error)
	
	// GetByEmail retrieves a user by their email address.
	GetByEmail(email string) (*internal.User, error)
	
	// GetAll retrieves all users.
	GetAll() ([]*internal.User, error)
}

// application struct holds dependencies for the application.
type application struct {
	DB   db.DbModel
	JSON api.JsonModel
}
