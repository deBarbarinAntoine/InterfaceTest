package api

// UserJSON represents a user in JSON format.
type UserJSON struct {
	ID         uint    `json:"id"`         // Unique identifier for the user
	FirstName  string  `json:"firstName"`  // User's first name
	LastName   string  `json:"lastName"`   // User's last name
	MaidenName string  `json:"maidenName"` // User's maiden name
	Age        int     `json:"age"`        // User's age
	Gender     string  `json:"gender"`     // User's gender (e.g., "male", "female")
	Email      string  `json:"email"`      // User's email address
	Phone      string  `json:"phone"`      // User's phone number
	Username   string  `json:"username"`   // User's username
	Password   string  `json:"password"`   // User's password (hashed)
	BirthDate  string  `json:"birthDate"`  // User's birth date in ISO format
	Image      string  `json:"image"`      // URL or path to the user's profile image
	BloodGroup string  `json:"bloodGroup"` // User's blood group
	Height     float64 `json:"height"`     // User's height in meters
	Weight     float64 `json:"weight"`     // User's weight in kilograms
	EyeColor   string  `json:"eyeColor"`   // User's eye color
	Hair       struct {
		Color string `json:"color"` // Color of the user's hair
		Type  string `json:"type"`  // Type of the user's hair (e.g., "straight", "curly")
	} `json:"hair"`            // Embedded struct for hair details
	Ip      string `json:"ip"` // User's IP address
	Address struct {
		Address     string `json:"address"`    // Street address of the location
		City        string `json:"city"`       // City where the location is located
		State       string `json:"state"`      // State or province where the location is located
		StateCode   string `json:"stateCode"`  // ISO code for the state or province
		PostalCode  string `json:"postalCode"` // Postal or ZIP code of the location
		Coordinates struct {
			Lat float64 `json:"lat"` // Latitude coordinate of the location
			Lng float64 `json:"lng"` // Longitude coordinate of the location
		} `json:"coordinates"`          // Embedded struct for geographic coordinates
		Country string `json:"country"` // Country where the location is located
	} `json:"address"`                    // Embedded struct for address details
	MacAddress string `json:"macAddress"` // User's MAC address
	University string `json:"university"` // User's university
	Bank       struct {
		CardExpire string `json:"cardExpire"` // Expiration date of the user's bank card
		CardNumber string `json:"cardNumber"` // Number of the user's bank card
		CardType   string `json:"cardType"`   // Type of the user's bank card (e.g., "Visa", "MasterCard")
		Currency   string `json:"currency"`   // Currency used by the user's bank account
		Iban       string `json:"iban"`       // International Bank Account Number of the user's bank account
	} `json:"bank"` // Embedded struct for bank details
	Company struct {
		Department string `json:"department"` // Department within the company where the user works
		Name       string `json:"name"`       // Name of the company
		Title      string `json:"title"`      // Title or position held by the user in the company
		Address    struct {
			Address     string `json:"address"`    // Street address of the company's location
			City        string `json:"city"`       // City where the company is located
			State       string `json:"state"`      // State or province where the company is located
			StateCode   string `json:"stateCode"`  // ISO code for the state or province
			PostalCode  string `json:"postalCode"` // Postal or ZIP code of the company's location
			Coordinates struct {
				Lat float64 `json:"lat"` // Latitude coordinate of the company's location
				Lng float64 `json:"lng"` // Longitude coordinate of the company's location
			} `json:"coordinates"`          // Embedded struct for geographic coordinates
			Country string `json:"country"` // Country where the company is located
		} `json:"address"` // Embedded struct for company address details
	} `json:"company"`                  // Embedded struct for company details
	Ein       string `json:"ein"`       // Employer Identification Number of the user's company
	Ssn       string `json:"ssn"`       // Social Security Number of the user
	UserAgent string `json:"userAgent"` // User agent string from the user's browser or device
	Crypto    struct {
		Coin    string `json:"coin"`    // Cryptocurrency used by the user
		Wallet  string `json:"wallet"`  // Wallet address of the user
		Network string `json:"network"` // Network on which the cryptocurrency wallet operates (e.g., "Bitcoin", "Ethereum")
	} `json:"crypto"`         // Embedded struct for cryptocurrency details
	Role string `json:"role"` // Role or permission level of the user in the system
}

// JsonProvider represents a model for handling JSON data.
type JsonProvider struct {
	
	// Path to the JSON file
	Path string
}
