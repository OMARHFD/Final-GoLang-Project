package main

type ServerPool struct {
	Backends []*Backend `json:"backends"`
	Current  uint64     `json:"current"` // Used for Round-Robin
}

//func (s ServerPool) GetNextValidPeer() *Backend {
//	counter := 0
//	initialValue := s.Current
//	for !s.Backends[s.Current].Alive && counter < (len(s.Backends)) {
		s.Current = s.Current % uint64((len(s.Backends)))
		s.Current++
		counter++
//	}
//
//	if s.Backends[s.Current].Alive {
//		return s.Backends[s.Current]
//	} else {
//		s.Current = initialValue
//		// errror handling
//	}
//
//}
