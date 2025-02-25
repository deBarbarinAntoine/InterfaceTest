package db

import (
	"InterfaceTest/internal"
)

// GetByID retrieves a user by its ID from the database, including its associated company and address information.
func (d DbModel) GetByID(id uint) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified ID, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.id = ?", id).First(&user)
	
	return &user, nil
}

// GetByUsername retrieves a user by its username from the database, including its associated company and address information.
func (d DbModel) GetByUsername(username string) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified username, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.username = ?", username).First(&user)
	
	return &user, nil
}

// GetByEmail retrieves a user by its email from the database, including its associated company and address information.
func (d DbModel) GetByEmail(email string) (*internal.User, error) {
	var user internal.User
	
	// Query the database for a user with the specified email, joining with Company, Company.Address, and Address tables
	d.Model(&user).Joins("Company").Joins("Company.Address").Joins("Address").Where("users.email = ?", email).First(&user)
	
	return &user, nil
}

// GetAll retrieves all users from the database, including their associated company and address information.
func (d DbModel) GetAll() ([]*internal.User, error) {
	var users []*internal.User
	
	// Query the database for all users, joining with Company, Company.Address, and Address tables
	d.Model(&users).Joins("Company").Joins("Company.Address").Joins("Address").Find(&users)
	
	return users, nil
}
