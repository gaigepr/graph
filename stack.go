package graph

import (
	"errors"
)

//
type Stack struct {
	s []*Vertex
}

//
func NewStack() *Stack {
	a := new(Stack)
	a.s = make([]*Vertex, 0)
	return a
}

//
func (s *Stack) Push(v *Vertex) {
	s.s = append(s.s, v)
}

//
func (s *Stack) Pop() (*Vertex, error) {
	l := len(s.s)
	if l == 0 {
		return &Vertex{}, errors.New("Stack underflow.")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

//
func (s *Stack) Size() int {
	return len(s.s)
}

//
func (s *Stack) Empty() bool {
	return len(s.s) == 0
}
