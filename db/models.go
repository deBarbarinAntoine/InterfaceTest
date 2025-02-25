package db

import (
	"gorm.io/gorm"
)

// DbModel represents a model for handling database data.
type DbModel struct {
	
	// The connection to the database using GORM
	*gorm.DB
}
