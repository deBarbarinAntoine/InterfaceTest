package db

import (
	"gorm.io/gorm"
)

// DatabaseProvider represents a model for handling database data.
type DatabaseProvider struct {
	
	// The connection to the database using GORM
	*gorm.DB
}
