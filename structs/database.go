package structs

// Database is a type definition for connecting to the db via gorm
type Database struct {
	name    string `json:"name"`
	dialect string `json:"dialect"`
}
