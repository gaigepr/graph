package graph

import (
	. "github.com/pborman/uuid"
)

// The Vertex struct.
type Vertex struct {
	uuid       UUID
	key        int // integer representation of the UUID???
	vtype      string
	properties map[string]int
}

// Allocates a new Vertex and return a pointer to it.
func NewVertex(vtype string, key int) *Vertex {
	var v *Vertex = new(Vertex)
	v.uuid = NewRandom()
	v.vtype = vtype
	v.key = key
	return v
}

// A Key method for use in search things.
func (v *Vertex) Key() int {
	return v.key
}

//
func (v *Vertex) Equal(x *Vertex) bool {
	return v.UUIDString() == x.UUIDString()
}

// Returns the uuid of v as a UUID which is an alias for [16]byte
func (v *Vertex) UUID() UUID {
	return v.uuid
}

// Returns the string representation of the uuid for v.
func (v *Vertex) UUIDString() string {
	return v.uuid.String()
}

//
func (v *Vertex) Type() string {
	return v.vtype
}

// Returns a copy of the props map of v.
func (v *Vertex) GetProperties() map[string]int {
	return v.properties
}

// Replaces v's props with the supplied p map[string]int.
func (v *Vertex) UpdateProperties(p map[string]int) bool {
	v.properties = p
	return true
}

// Returns the membership status of v in g.
func (v *Vertex) Member(g *Graph) bool {
	if _, ex := g.vertices[v.UUIDString()]; ex {
		return true
	}
	return false
}
