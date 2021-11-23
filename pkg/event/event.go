package event

type Type string

const (
	TypeIsAnswer Type = "answer"
	TypeIsClear  Type = "clear"
)

type Event struct {
	Type    Type
	Content interface{}
}

type Answer struct {
	Team int `json:"team"`
}

func NewAnswer(team int) Event {
	return Event{
		Type: TypeIsAnswer,
		Content: Answer{
			Team: team,
		},
	}
}

func (e Event) Answer() (Answer, bool) {
	if e.Type != TypeIsAnswer {
		return Answer{}, false
	}

	m, ok := e.Content.(map[string]interface{})
	if !ok {
		return Answer{}, false
	}
	answer := Answer{}

	teamInterface, ok := m["team"]
	if !ok {
		return Answer{}, false
	}

	if team, ok := teamInterface.(float64); ok {
		answer.Team = int(team)
	} else {
		return Answer{}, false
	}

	return answer, true
}

type Clear struct {
}

func NewClear() Event {
	return Event{
		Type:    TypeIsClear,
		Content: Clear{},
	}
}

func (e Event) Clear() (Clear, bool) {
	if e.Type != TypeIsClear {
		return Clear{}, false
	}

	return Clear{}, true
}
