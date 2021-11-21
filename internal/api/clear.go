package api

import "net/http"

// ClearHandler will handle the clear route
// swagger:route POST /clear Answer ClearRequest
//
// Request the server to clear the team answerer
//
// Produces:
// 	- application/json
//
// Schemes: http, https
//
// Responses:
// 	default: GenericError
// 	200:
func (s *API) ClearHandler(w http.ResponseWriter, r *http.Request) {
	s.state.ClearHand()
	Respond(w, http.StatusOK, nil)
}
