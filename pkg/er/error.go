package er

// Generic API error
// swagger:model GenericError
type GenericError struct {
	// in:body
	// Message describing the error
	// Required: true
	Error string `json:"error"`
}
