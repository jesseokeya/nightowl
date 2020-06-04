package structs

// Error is a type definition for http error status
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
