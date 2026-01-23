package main

import "fmt"

type ServerPool struct {
	Backends []*Backend `json:"backends"`
	Current  uint64     `json:"current"` // Used for Round-Robin
}

func (s *ServerPool) GetNextValidPeer() *Backend {
	counter := 0
	initialValue := s.Current
	s.Current = (s.Current + 1) % uint64((len(s.Backends)))

	for !s.Backends[s.Current].Alive && counter < (len(s.Backends)) {
		s.Current = (s.Current + 1) % uint64((len(s.Backends)))
		counter++
	}

	if s.Backends[s.Current].Alive {
		return s.Backends[s.Current]
	} else {
		s.Current = initialValue
		fmt.Println("No backend found :(")
		return nil
	}
}
