package api

import (
	"net/http"
	"time"

	"github.com/Cyfell/BroQuiz/pkg/infos"
)

// InfosHandler will handle the infos route
// swagger:route GET /infos Miscellaneous infos
//
// Return informations on the server
//
// Produces:
// 	- application/json
//
// Schemes: http, https
//
// Responses:
// 	default: GenericError
// 	200: Infos
func InfosHandler(w http.ResponseWriter, r *http.Request) {
	Respond(w, http.StatusOK, infos.Infos{
		Time: time.Now().UTC(),
	})
}
