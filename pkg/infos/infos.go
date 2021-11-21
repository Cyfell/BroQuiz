package infos

import "time"

// swagger:model Infos
type Infos struct {
	// in:body
	// The response generation time (in UTC)
	// Required: true
	Time time.Time `json:"time"`
}
