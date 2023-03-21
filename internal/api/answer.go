package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AnswerHandler will handle the answer route
// swagger:route POST /answer/{team} Answer AnswerRequest
//
// # Request the server for an answer
//
// Produces:
//   - application/json
//
// Schemes: http, https
//
// Responses:
//
//	default: GenericError
//	200: AnswerResponse
func (s *API) AnswerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Get team number from path
	team, err := strconv.Atoi(vars["team"])
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	// Get correct code
	var code int
	hasHand := s.state.GetHand(team)
	if hasHand {
		code = http.StatusCreated
	} else {
		code = http.StatusConflict
	}

	// Respond with correct code
	Respond(w, code, nil)
}
