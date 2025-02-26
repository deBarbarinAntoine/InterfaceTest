package db

import (
	"InterfaceTest/internal"
)

// Migrate creates the database according to the model passed.
func (d DatabaseProvider) Migrate(model any) error {
	return d.DB.AutoMigrate(model)
}

// Count gets the number of rows in a table according to the model passed.
func (d DatabaseProvider) Count(model any) (int, error) {
	var count int64
	d.Model(model).Count(&count)
	return int(count), nil
}

// Create adds data in the database according to the model passed.
func (d DatabaseProvider) Create(model any) error {
	d.DB.Create(model)
	return nil
}

// GetByID retrieves a user by its ID from the database, including its associated company and address information.
func (d DatabaseProvider) GetByID(id uint) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified ID, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.id = ?", id).First(&user)
	
	return &user, nil
}

// GetByUsername retrieves a user by its username from the database, including its associated company and address information.
func (d DatabaseProvider) GetByUsername(username string) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified username, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.username = ?", username).First(&user)
	
	return &user, nil
}

// GetByEmail retrieves a user by its email from the database, including its associated company and address information.
func (d DatabaseProvider) GetByEmail(email string) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified email, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.email = ?", email).First(&user)
	
	return &user, nil
}

// GetAll retrieves all users from the database, including their associated company and address information.
func (d DatabaseProvider) GetAll() ([]*internal.User, error) {
	var users []*internal.User
	
	// Query the database for all users, joining with Company, Company.Address, and Address tables
	d.Model(&users).Joins("Company").Joins("Company.Address").Joins("Address").Find(&users)
	
	return users, nil
}
