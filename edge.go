package graph

import (
	. "github.com/pborman/uuid"
)

//
type Edge struct {
	uuid       UUID
	from       UUID
	to         UUID
	weight     int
	etype      string
	properties interface{}
}

//
func NewEdge(weight int, etype string) *Edge {
	return &Edge{
		uuid:   NewRandom(),
		weight: weight,
		etype:  etype,
	}
}

//
func (e *Edge) UUID() UUID {
	return e.uuid
}

//
func (e *Edge) UUIDString() string {
	return e.uuid.String()
}

//
func (e *Edge) Weight() int {
	return e.weight
}

//
func (e *Edge) SetWeight(newWeight int) {
	e.weight = newWeight
}

//
func (e *Edge) RelationString() string {
	return e.etype
}

//
func (e *Edge) Vertices() (UUID, UUID) {
	return e.from, e.to
}
