package api

import (
	"encoding/json"
	"fmt"
	"os"
	
	"InterfaceTest/internal"
)

// toUser converts a UserJSON object to an internal.User object.
func (user UserJSON) toUser() internal.User {
	return internal.User{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MaidenName: user.MaidenName,
		Age:        user.Age,
		Gender:     user.Gender,
		Email:      user.Email,
		Phone:      user.Phone,
		Username:   user.Username,
		Password:   user.Password,
		BirthDate:  user.BirthDate,
		Image:      user.Image,
		BloodGroup: user.BloodGroup,
		Height:     user.Height,
		Weight:     user.Weight,
		EyeColor:   user.EyeColor,
		Hair:       user.Hair,
		Ip:         user.Ip,
		Address: internal.Address{
			Address:     user.Address.Address,
			City:        user.Address.City,
			State:       user.Address.State,
			StateCode:   user.Address.StateCode,
			PostalCode:  user.Address.PostalCode,
			Coordinates: user.Address.Coordinates,
			Country:     user.Address.Country,
		},
		MacAddress: user.MacAddress,
		University: user.University,
		Bank:       user.Bank,
		Company: internal.Company{
			Department: user.Company.Department,
			Name:       user.Company.Name,
			Title:      user.Company.Title,
			Address: internal.Address{
				Address:     user.Company.Address.Address,
				City:        user.Company.Address.City,
				State:       user.Company.Address.State,
				StateCode:   user.Company.Address.StateCode,
				PostalCode:  user.Company.Address.PostalCode,
				Coordinates: user.Company.Address.Coordinates,
				Country:     user.Company.Address.Country,
			},
		},
		Ein:       user.Ein,
		Ssn:       user.Ssn,
		UserAgent: user.UserAgent,
		Crypto:    user.Crypto,
		Role:      user.Role,
	}
}

// GetByID retrieves a user by their ID from the JSON file.
func (j JsonModel) GetByID(id uint) (*internal.User, error) {
	
	// Read the JSON file from the specified path
	data, err := os.ReadFile(j.Path)
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the JSON data into a slice of UserJSON structs
	var usersJSON []UserJSON
	if err := json.Unmarshal(data, &usersJSON); err != nil {
		return nil, err
	}
	
	// Iterate through each UserJSON struct to find a match by ID
	for _, userJSON := range usersJSON {
		if userJSON.ID == id {
			
			// Convert the UserJSON struct to an internal.User struct
			user := userJSON.toUser()
			return &user, nil
		}
	}
	
	// Return an error if no matching user is found
	return nil, fmt.Errorf("userJSON with id %d not found", id)
}

// GetByUsername retrieves a user by their username from the JSON file.
func (j JsonModel) GetByUsername(username string) (*internal.User, error) {
	
	// Read the JSON file from the specified path
	data, err := os.ReadFile(j.Path)
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the JSON data into a slice of UserJSON structs
	var usersJSON []UserJSON
	if err := json.Unmarshal(data, &usersJSON); err != nil {
		return nil, err
	}
	
	// Iterate through each UserJSON struct to find a match by username
	for _, userJSON := range usersJSON {
		if userJSON.Username == username {
			
			// Convert the UserJSON struct to an internal.User struct
			user := userJSON.toUser()
			return &user, nil
		}
	}
	
	// Return an error if no matching user is found
	return nil, fmt.Errorf("user with username %s not found", username)
}

// GetByEmail retrieves a user by their email from the JSON file.
func (j JsonModel) GetByEmail(email string) (*internal.User, error) {
	
	// Read the JSON file from the specified path
	data, err := os.ReadFile(j.Path)
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the JSON data into a slice of UserJSON structs
	var usersJSON []UserJSON
	if err := json.Unmarshal(data, &usersJSON); err != nil {
		return nil, err
	}
	
	// Iterate through each UserJSON struct to find a match by email
	for _, userJSON := range usersJSON {
		if userJSON.Email == email {
			
			// Convert the UserJSON struct to an internal.User struct
			user := userJSON.toUser()
			return &user, nil
		}
	}
	
	// Return an error if no matching user is found
	return nil, fmt.Errorf("user with email %s not found", email)
}

// GetAll retrieves all users from the JSON file.
func (j JsonModel) GetAll() ([]*internal.User, error) {
	// Read the JSON file from the specified path
	data, err := os.ReadFile(j.Path)
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the JSON data into a slice of UserJSON structs
	var usersJSON []UserJSON
	if err := json.Unmarshal(data, &usersJSON); err != nil {
		return nil, err
	}
	
	// Create a slice to hold the converted internal.User structs
	var users []*internal.User
	for _, userJSON := range usersJSON {
		
		// Convert each UserJSON struct to an internal.User struct and append it to the slice
		user := userJSON.toUser()
		users = append(users, &user)
	}
	
	// Return the slice of internal.User structs
	return users, nil
}
