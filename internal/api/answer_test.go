package api

import (
	"net/http"
)

func (suite *APITestSuite) TestAnswer() {
	as := suite.Require()

	// Given a running server

	// When a first answer is triggered
	req, err := http.NewRequest("POST", "/answer/1", nil)
	as.NoError(err)
	response := suite.ExecuteRequest(suite.api.router, req)

	// Then the request is created
	as.Equal(http.StatusCreated, response.Code)

	// When a second answer is triggered by another team
	req, err = http.NewRequest("POST", "/answer/2", nil)
	as.NoError(err)
	response = suite.ExecuteRequest(suite.api.router, req)

	// Then the request is Conflict
	as.Equal(http.StatusConflict, response.Code)
}
