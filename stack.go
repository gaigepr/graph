package graph

import (
	"errors"
)

//
type Stack struct {
	s      []*Vertex
	length int
}

//
func (s *Stack) Push(v *Vertex) {
	s.s = append(s.s, v)
	s.length += 1
}

//
func (s *Stack) Pop() (*Vertex, error) {
	l := s.length
	if l == 0 {
		return &Vertex{}, errors.New("Stack underflow.")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	s.length -= 1
	return res, nil
}

//
func (s *Stack) Size() int {
	return s.length
}

//
func (s *Stack) Empty() bool {
	return s.length == 0
}
