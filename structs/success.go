package structs

// Success is a type definition for http success status
type Success struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
