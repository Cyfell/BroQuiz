package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Cyfell/BroQuiz/pkg/event"
	"github.com/gorilla/websocket"
)

func (suite *APITestSuite) TestEvents() {
	as := suite.Require()

	// Given a server
	s := httptest.NewServer(suite.api.router)
	defer s.Close()

	// When establishing the websocket connection
	u := "ws" + strings.TrimPrefix(s.URL, "http") + "/events"
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	as.NoError(err)
	defer ws.Close()

	// And a team tries to get hand
	req, err := http.NewRequest("POST", "/answer/1", nil)
	as.NoError(err)
	response := suite.ExecuteRequest(suite.api.router, req)
	as.Equal(http.StatusCreated, response.Code)

	// Then we receive the corresponding event
	evt := event.Event{}
	as.NoError(ws.ReadJSON(&evt))
	as.Equal(event.TypeIsAnswer, evt.Type)
	answer, ok := evt.Answer()
	as.True(ok)
	as.Equal(event.Answer{
		Team: 1,
	}, answer)

	// And the hand is released
	req, err = http.NewRequest("POST", "/clear", nil)
	as.NoError(err)
	response = suite.ExecuteRequest(suite.api.router, req)
	as.Equal(http.StatusOK, response.Code)

	// Then we receive the corresponding event
	evt = event.Event{}
	as.NoError(ws.ReadJSON(&evt))
	as.Equal(event.TypeIsClear, evt.Type)
	clear, ok := evt.Clear()
	as.True(ok)
	as.Equal(event.Clear{}, clear)
}
