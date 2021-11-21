package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cyfell/BroQuiz/pkg/er"
)

func Respond(w http.ResponseWriter, status int, v interface{}) error {
	resp, err := json.Marshal(v)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return err
	}
	w.WriteHeader(status)
	fmt.Fprint(w, string(resp))
	return nil
}

func RespondError(w http.ResponseWriter, status int, err error) {
	resp, err := json.Marshal(&er.GenericError{
		Error: err.Error(),
	})
	if err != nil {
		http.Error(w, err.Error(), status)
	} else {
		http.Error(w, string(resp), status)
	}
}
