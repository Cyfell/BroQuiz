package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// EventsHandler will handle the events route
// swagger:route GET /events Events events
//
// Return a websocket which returns events from the server
//
// Produces:
// 	- application/json
//
// Schemes: http, https
//
// Responses:
// 	default: GenericError
// 	200:
func (s *API) EventsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}
	defer conn.Close()

	id, ch := s.state.NewEventsChannel()
	defer s.state.RemoveEventsChannel(id)

	for {
		event := <-ch
		err := conn.WriteJSON(event)
		if err != nil {
			fmt.Println(err)
		}
	}
}
