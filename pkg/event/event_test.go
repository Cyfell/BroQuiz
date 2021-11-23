package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestEventsSuite(t *testing.T) {
	suite.Run(t, new(EventsTestSuite))
}

type EventsTestSuite struct {
	suite.Suite
}

func (suite *EventsTestSuite) TestEventToAnswer() {
	as := suite.Require()

	answer := Answer{
		Team: 1,
	}
	evt := NewAnswer(1)

	b, err := json.Marshal(evt)
	as.NoError(err)

	recvEvent := Event{}
	as.NoError(json.Unmarshal(b, &recvEvent))

	recvAnswer, ok := recvEvent.Answer()
	as.True(ok)
	as.Equal(answer, recvAnswer)
}

func (suite *EventsTestSuite) TestEventToClear() {
	as := suite.Require()

	clear := Clear{}
	evt := NewClear()

	b, err := json.Marshal(evt)
	as.NoError(err)

	recvEvent := Event{}
	as.NoError(json.Unmarshal(b, &recvEvent))

	recvClear, ok := recvEvent.Clear()
	as.True(ok)
	as.Equal(clear, recvClear)
}
