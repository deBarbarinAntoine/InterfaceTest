package internal

import (
	"time"
	
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`   // Unique identifier for the user
	CreatedAt  time.Time `json:"-"`                      // Timestamp when the user was created (automatically managed by GORM)
	UpdatedAt  time.Time `json:"-"`                      // Timestamp when the user was last updated (automatically managed by GORM)
	FirstName  string    `json:"firstName"`              // User's first name
	LastName   string    `json:"lastName"`               // User's last name
	MaidenName string    `json:"maidenName"`             // User's maiden name
	Age        int       `json:"age"`                    // User's age
	Gender     string    `json:"gender"`                 // User's gender (e.g., "male", "female")
	Email      string    `json:"email" gorm:"unique"`    // User's email address, must be unique
	Phone      string    `json:"phone"`                  // User's phone number
	Username   string    `json:"username" gorm:"unique"` // User's username, must be unique
	Password   string    `json:"password"`               // User's password (hashed)
	BirthDate  string    `json:"birthDate"`              // User's birth date in ISO format
	Image      string    `json:"image"`                  // URL or path to the user's profile image
	BloodGroup string    `json:"bloodGroup"`             // User's blood group
	Height     float64   `json:"height"`                 // User's height in meters
	Weight     float64   `json:"weight"`                 // User's weight in kilograms
	EyeColor   string    `json:"eyeColor"`               // User's eye color
	Hair       struct {
		Color string `json:"color"` // Color of the user's hair
		Type  string `json:"type"`  // Type of the user's hair (e.g., "straight", "curly")
	} `json:"hair" gorm:"embedded;embeddedPrefix:hair_"` // Embedded struct for hair details
	Ip         string  `json:"ip"`                       // User's IP address
	AddressID  uint    `json:"-"`                        // Foreign key to the user's address (not exposed in JSON)
	Address    Address `json:"address"`                  // User's address information
	MacAddress string  `json:"macAddress"`               // User's MAC address
	University string  `json:"university"`               // User's university
	Bank       struct {
		CardExpire string `json:"cardExpire"` // Expiration date of the user's bank card
		CardNumber string `json:"cardNumber"` // Number of the user's bank card
		CardType   string `json:"cardType"`   // Type of the user's bank card (e.g., "Visa", "MasterCard")
		Currency   string `json:"currency"`   // Currency used by the user's bank account
		Iban       string `json:"iban"`       // International Bank Account Number of the user's bank account
	} `json:"bank" gorm:"embedded;embeddedPrefix:bank_"` // Embedded struct for bank details
	CompanyID uint    `json:"-"`                         // Foreign key to the user's company (not exposed in JSON)
	Company   Company `json:"company"`                   // User's company information
	Ein       string  `json:"ein"`                       // Employer Identification Number of the user's company
	Ssn       string  `json:"ssn"`                       // Social Security Number of the user
	UserAgent string  `json:"userAgent"`                 // User agent string from the user's browser or device
	Crypto    struct {
		Coin    string `json:"coin"`    // Cryptocurrency used by the user
		Wallet  string `json:"wallet"`  // Wallet address of the user
		Network string `json:"network"` // Network on which the cryptocurrency wallet operates (e.g., "Bitcoin", "Ethereum")
	} `json:"crypto" gorm:"embedded;embeddedPrefix:crypto_"` // Embedded struct for cryptocurrency details
	Role string `json:"role"`                                // Role or permission level of the user in the system
}

// Company represents a company associated with a user.
type Company struct {
	gorm.Model
	Department string  `json:"department"` // Department within the company where the user works
	Name       string  `json:"name"`       // Name of the company
	Title      string  `json:"title"`      // Title or position held by the user in the company
	AddressID  uint    `json:"-"`          // Foreign key to the company's address (not exposed in JSON)
	Address    Address `json:"address"`    // Company's address information
}

// Address represents an address associated with a user or company.
type Address struct {
	gorm.Model
	Address     string `json:"address"`    // Street address of the location
	City        string `json:"city"`       // City where the location is located
	State       string `json:"state"`      // State or province where the location is located
	StateCode   string `json:"stateCode"`  // ISO code for the state or province
	PostalCode  string `json:"postalCode"` // Postal or ZIP code of the location
	Coordinates struct {
		Lat float64 `json:"lat"` // Latitude coordinate of the location
		Lng float64 `json:"lng"` // Longitude coordinate of the location
	} `json:"coordinates" gorm:"embedded;embeddedPrefix:coords_"` // Embedded struct for geographic coordinates
	Country string `json:"country"`                               // Country where the location is located
}
