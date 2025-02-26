package main

import (
	"fmt"
	"os"
	
	"InterfaceTest/api"
	"InterfaceTest/db"
	"InterfaceTest/internal"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// exitError prints the error message and exits the application
func exitError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

// main is the entry point of the application
func main() {
	
	// Establish the connection to the database
	dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		exitError(fmt.Errorf("failed to connect database: %w", err))
	}
	
	// Create the application instance with its configuration
	app := application{
		DB:   db.DatabaseProvider{DB: dbConn},
		JSON: api.JsonProvider{Path: "data.json"},
	}
	
	// Migrate the schema
	err = app.DB.Migrate(&internal.User{})
	if err != nil {
		exitError(fmt.Errorf("migrations failed: %w", err))
	}
	
	// Get the users from JSON
	users, err := app.JSON.GetAll()
	if err != nil {
		exitError(fmt.Errorf("error fetching json data: %w", err))
	}
	
	// Create users if they don't already exist in the database
	if count, err := app.DB.Count(&internal.User{}); count == 0 {
		fmt.Println("No users found: adding users...")
		err = app.DB.Create(users)
		if err != nil {
			exitError(fmt.Errorf("creating users failed: %w", err))
		}
	} else {
		fmt.Printf("%d Users already exist: skipping...\n", count)
	}
	
	// Get the user Emily Smith from JSON file by ID
	emilyJsonByID, err := app.JSON.GetByID(1)
	if err != nil {
		exitError(fmt.Errorf("error fetching json data: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyJsonByID.ID)
	fmt.Printf("%+v\n", emilyJsonByID)
	
	// Get the user Emily Smith from JSON file by Email
	emilyJsonByEmail, err := app.JSON.GetByEmail("emily.johnson@x.dummyjson.com")
	if err != nil {
		exitError(fmt.Errorf("error fetching json data: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyJsonByEmail.ID)
	fmt.Printf("%+v\n", emilyJsonByEmail)
	
	// Get the user Emily Smith from JSON file by Username
	emilyJsonByUsername, err := app.JSON.GetByUsername("emilys")
	if err != nil {
		exitError(fmt.Errorf("error fetching json data: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyJsonByUsername.ID)
	fmt.Printf("%+v\n", emilyJsonByUsername)
	
	// Get the user Emily Smith from DB by ID
	emilyDbByID, err := app.DB.GetByID(1)
	if err != nil {
		exitError(fmt.Errorf("error fetching data from database: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyDbByID.ID)
	fmt.Printf("%+v\n", emilyDbByID)
	
	// Get the user Emily Smith from DB by Email
	emilyDbByEmail, err := app.DB.GetByEmail("emily.johnson@x.dummyjson.com")
	if err != nil {
		exitError(fmt.Errorf("error fetching data from database: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyDbByEmail.ID)
	fmt.Printf("%+v\n", emilyDbByEmail)
	
	// Get the user Emily Smith from DB by Username
	emilyDbByUsername, err := app.DB.GetByUsername("emilys")
	if err != nil {
		exitError(fmt.Errorf("error fetching data from database: %w", err))
	}
	// Print the user found
	fmt.Println()
	fmt.Printf("User Nº%d\n", emilyDbByUsername.ID)
	fmt.Printf("%+v\n", emilyDbByUsername)
	
	// Get all users from DB
	usersDb, err := app.DB.GetAll()
	if err != nil {
		exitError(fmt.Errorf("error fetching all data from database: %w", err))
	}
	
	// Print all users one by one
	fmt.Println()
	fmt.Println("All Users of the database:")
	for _, user := range usersDb {
		fmt.Println()
		fmt.Printf("User Nº%d:\n", user.ID)
		fmt.Printf("%+v\n", *user)
	}
}
