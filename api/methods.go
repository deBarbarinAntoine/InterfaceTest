package api

import (
	"encoding/json"
	"errors"
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

func (j JsonProvider) getJsonData() ([]UserJSON, error) {
	
	// Read the JSON file from the specified path
	data, err := os.ReadFile(j.Path)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}
	
	// Unmarshal the JSON data into a slice of UserJSON structs
	var usersJSON []UserJSON
	if err := json.Unmarshal(data, &usersJSON); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	
	return usersJSON, nil
}

// Migrate is here to implement the UserCRUD interface, nothing else.
func (j JsonProvider) Migrate(model any) error {
	return nil
}

// Count gets the number of users in the JSON file.
func (j JsonProvider) Count(model any) (int, error) {
	
	// Get users from JSON file
	usersJSON, err := j.getJsonData()
	if err != nil {
		return 0, err
	}
	
	// Return the number of items in the list
	return len(usersJSON), nil
}

// Create append data in the JSON file.
func (j JsonProvider) Create(data any) error {
	
	// Checking the data type
	var users []internal.User
	switch data.(type) {
	case internal.User:
		users = append(users, data.(internal.User))
	case []internal.User:
		users = data.([]internal.User)
	default:
		return errors.New("invalid data type")
	}
	
	// Get all users from JSON file.
	allUsers, err := j.GetAll()
	if err != nil {
		return fmt.Errorf("error getting all users from JSON: %w", err)
	}
	
	// Look for new users
	var isNewUser bool
	for _, user := range users {
		for _, existentUser := range allUsers {
			if user.ID != existentUser.ID && user.Email != existentUser.Email && user.Username != existentUser.Username {
				isNewUser = true
				allUsers = append(allUsers, &user)
			}
		}
	}
	
	// Update JSON file if there are new users
	if isNewUser {
		
		// Convert allUsers to JSON format
		newData, err := json.MarshalIndent(allUsers, "", "\t")
		if err != nil {
			return fmt.Errorf("error marshalling JSON: %w", err)
		}
		
		// Write allUsers to JSON file
		err = os.WriteFile(j.Path, newData, 0644)
		if err != nil {
			return fmt.Errorf("error writing JSON: %w", err)
		}
	}
	
	return nil
}

// GetByID retrieves a user by their ID from the JSON file.
func (j JsonProvider) GetByID(id uint) (*internal.User, error) {
	
	// Get users from JSON file
	usersJSON, err := j.getJsonData()
	if err != nil {
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
func (j JsonProvider) GetByUsername(username string) (*internal.User, error) {
	
	// Get users from JSON file
	usersJSON, err := j.getJsonData()
	if err != nil {
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
func (j JsonProvider) GetByEmail(email string) (*internal.User, error) {
	
	// Get users from JSON file
	usersJSON, err := j.getJsonData()
	if err != nil {
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
func (j JsonProvider) GetAll() ([]*internal.User, error) {
	
	// Get users from JSON file
	usersJSON, err := j.getJsonData()
	if err != nil {
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
