package api

import (
	"net/http"
	"strconv"

	"github.com/Cyfell/BroQuiz/pkg/answer"
	"github.com/gorilla/mux"
)

// AnswerHandler will handle the answer route
// swagger:route POST /answer/{teamID} Answer AnswerRequest
//
// Request the server for an answer
//
// Produces:
// 	- application/json
//
// Schemes: http, https
//
// Responses:
// 	default: GenericError
// 	200: AnswerResponse
func (s *API) AnswerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamNB, err := strconv.ParseUint(vars["teamID"], 10, 32)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	Respond(w, http.StatusOK, answer.Response{
		TeamNB:  teamNB,
		HasHand: s.state.GetHand(teamNB),
	})
}
