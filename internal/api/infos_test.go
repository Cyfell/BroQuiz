package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Cyfell/BroQuiz/pkg/infos"
)

func (suite *APITestSuite) TestInfos() {
	as := suite.Require()

	expectedResponse := infos.Infos{
		Time: time.Now().UTC(),
	}

	req, err := http.NewRequest("GET", "/infos", nil)
	as.NoError(err)
	response := suite.ExecuteRequest(suite.api.router, req)

	as.Equal(http.StatusOK, response.Code)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	as.NoError(err)

	var resp infos.Infos
	as.NoError(json.Unmarshal(bodyBytes, &resp))

	as.True(expectedResponse.Time.Add(-time.Second).Before(resp.Time))
	as.True(expectedResponse.Time.Add(time.Second).After(resp.Time))
	as.Equal(time.UTC, resp.Time.Location())
}
