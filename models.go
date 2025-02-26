package main

import (
	"InterfaceTest/internal"
)

// CRUDStrategy interface defines methods for CRUD operations on users.
type CRUDStrategy interface {
	
	// Migrate creates the database (only provided for databases kind of provider)
	Migrate(model any) error
	
	// Count gets the number of users stored
	Count(model any) int
	
	// Create inserts data in the provider.
	Create(data any) error
	
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
	DB   CRUDStrategy
	JSON CRUDStrategy
}
