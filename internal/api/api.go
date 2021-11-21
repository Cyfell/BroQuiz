package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type API struct {
	router *mux.Router
	url    string
}

func New(ctx context.Context, c Config) (*API, error) {
	s := &API{
		router: newRouter(ctx),
		url:    fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	return s, nil
}

func newRouter(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/infos", InfosHandler(ctx)).Methods("GET")

	return r
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
