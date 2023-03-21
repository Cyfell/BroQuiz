package api

import (
	"net/http"
)

func (suite *APITestSuite) TestClear() {
	as := suite.Require()

	// Given a running server and a first answer is triggered
	req, err := http.NewRequest("POST", "/answer/1", nil)
	as.NoError(err)
	response := suite.ExecuteRequest(suite.api.router, req)
	as.Equal(http.StatusCreated, response.Code)

	// When the clear command is issued
	req, err = http.NewRequest("POST", "/clear", nil)
	as.NoError(err)
	response = suite.ExecuteRequest(suite.api.router, req)
	as.Equal(http.StatusOK, response.Code)

	// And a second answer is triggered by another team
	req, err = http.NewRequest("POST", "/answer/2", nil)
	as.NoError(err)
	response = suite.ExecuteRequest(suite.api.router, req)

	// Then the request is OK
	as.Equal(http.StatusCreated, response.Code)
}
