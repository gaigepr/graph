package graph

import (
	. "github.com/pborman/uuid"
)

// The Edge struct.
type Edge struct {
	uuid   UUID
	from   UUID
	to     UUID
	weight int
	//relation   string
	properties map[string]int
}

// Allocated a new Edge and returns a pointer to it.
////func NewEdge(weight int, relation string) *Edge {
func NewEdge(weight int) *Edge {
	var e *Edge = new(Edge)
	e.uuid = NewRandom()
	e.weight = weight
	return e
}

// A library method for changing the uuid
func (e *Edge) changeUUID(n string) {
	uuid := uuidFromString(n)
	e.uuid = uuid
}

// Returns the uuid of e as a UUID.
func (e *Edge) UUID() UUID {
	return e.uuid
}

// Returns the uuid of e as a string.
func (e *Edge) UUIDString() string {
	return e.uuid.String()
}

// Returns the from, to Vertex UUIDs as a pair.
func (e *Edge) Vertices() (UUID, UUID) {
	return e.from, e.to
}

// Returns the weight of e.
func (e *Edge) Weight() int {
	return e.weight
}

// Sets the weight for e.
//func (e *Edge) SetWeight(w int) {
//	e.weight = w
//}

// Returns the relationship type of e.
//func (e *Edge) Relation() string {
//	return e.relation
//}

// Returns a copy of the properties field of an Edge.
func (e *Edge) Properties() map[string]int {
	return e.properties
}

// Replaces the properties of e with p.
func (e *Edge) UpdateProperties(p map[string]int) {
	e.properties = p
}

// Returns the membership status of e in g.
func (e *Edge) Member(g *Graph) bool {
	if _, ex := g.edges[e.UUIDString()]; ex {
		return true
	}
	return false
}
