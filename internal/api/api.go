package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Cyfell/BroQuiz/internal/logic"
	"github.com/gorilla/mux"
)

type Config struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type API struct {
	router *mux.Router
	url    string
	state  *logic.State
}

func New(ctx context.Context, c Config) (*API, error) {
	s := &API{
		url: fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	return s, nil
}

func (s *API) Init() {
	r := mux.NewRouter()

	r.HandleFunc("/infos", InfosHandler).Methods("GET")
	r.HandleFunc("/answer/{teamID}", s.AnswerHandler).Methods("POST")

	s.router = r
	s.state = logic.NewState()
}

func (s *API) Run() error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
