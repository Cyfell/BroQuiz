package logic

import (
	"sync"
)

type State struct {
	mu     sync.Mutex
	teamNb uint64
}

func NewState() *State {
	return &State{
		teamNb: 0,
	}
}

func (s *State) GetHand(teamNb uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.teamNb != 0 {
		return false
	}

	s.teamNb = teamNb
	return true
}
