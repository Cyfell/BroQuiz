package answer

// swagger:model AnswerResponse
type Response struct {
	// in:body
	// The team number that as attempted an answer
	// Required: true
	Team int
	// Informs if the team has the hand or not
	// Required: true
	HasHand bool
}
