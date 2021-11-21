package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

type APITestSuite struct {
	suite.Suite
	api *API
}

func (suite *APITestSuite) ExecuteRequest(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func (suite *APITestSuite) BeforeTest(suiteName, testName string) {
	ctx := context.Background()

	api, err := New(
		ctx,
		Config{
			Host: "127.0.0.1",
			Port: 2000,
		},
	)

	if suite.NoError(err) {
		suite.api = api
	}
}
