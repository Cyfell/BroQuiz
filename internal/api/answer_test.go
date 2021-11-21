package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Cyfell/BroQuiz/pkg/answer"
)

func (suite *APITestSuite) TestAnswer() {
	as := suite.Require()

	// Given a running server

	// When a first answer is triggered
	req, err := http.NewRequest("POST", "/answer/1", nil)
	as.NoError(err)
	response := suite.ExecuteRequest(suite.api.router, req)

	// Then the request is OK
	as.Equal(http.StatusOK, response.Code)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	as.NoError(err)

	// And the answer returns the team number and the status with the hand
	var resp answer.Response
	as.NoError(json.Unmarshal(bodyBytes, &resp))

	expectedResp := answer.Response{
		TeamNB:  1,
		HasHand: true,
	}
	as.Equal(expectedResp, resp)

	// When a second answer is triggered by another team
	req, err = http.NewRequest("POST", "/answer/2", nil)
	as.NoError(err)
	response = suite.ExecuteRequest(suite.api.router, req)

	// Then the request is OK
	as.Equal(http.StatusOK, response.Code)

	bodyBytes, err = ioutil.ReadAll(response.Body)
	as.NoError(err)

	// And the answer returns the team number and the status without the hand
	as.NoError(json.Unmarshal(bodyBytes, &resp))

	expectedResp = answer.Response{
		TeamNB:  2,
		HasHand: false,
	}
	as.Equal(expectedResp, resp)
}
