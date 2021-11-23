package logic

import (
	"sync"

	"github.com/Cyfell/BroQuiz/pkg/event"
)

type State struct {
	mu              sync.Mutex
	team            int
	eventsChannelID uint
	eventsChannels  map[uint]chan event.Event
}

func NewState() *State {
	return &State{
		team:            0,
		eventsChannelID: 0,
		eventsChannels:  make(map[uint]chan event.Event),
	}
}

func (s *State) GetHand(team int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.team != 0 {
		return false
	}

	s.team = team
	s.sendEvent(event.NewAnswer(team))

	return true
}

func (s *State) ClearHand() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.team = 0
	s.sendEvent(event.NewClear())
}

func (s *State) NewEventsChannel() (uint, chan event.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan event.Event)
	id := s.eventsChannelID

	s.eventsChannels[id] = ch
	s.eventsChannelID++

	return id, ch
}

func (s *State) RemoveEventsChannel(id uint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	close(s.eventsChannels[id])
	delete(s.eventsChannels, id)
}

func (s *State) sendEvent(e event.Event) {
	for _, ch := range s.eventsChannels {
		ch <- e
	}
}
