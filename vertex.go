package graph

import (
	. "github.com/pborman/uuid"
)

// The Vertex struct.
type Vertex struct {
	uuid       UUID
	vtype      string
	properties map[string]int
}

// Allocates a new Vertex and return a pointer to it.
func NewVertex(vtype string) *Vertex {
	var v *Vertex = new(Vertex)
	v.uuid = NewRandom()
	v.vtype = vtype
	return v
}

//
/*
func uuidFromString(s string) UUID {
	// method Parse(s) takes a string like:
	//xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	// and created a valid UUID type
	// Use this for testing so that the small strig supplied can be used as a uuid
	// This requires s to be unqiue still.
	// Sugar for using "a" and "b" as easy to remember UUIDs

	// WARN: Can try to access out of bounds on `base`
	var base string = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	var temp string = ""
	for k, v := range s {
		if base[k] == '-' {
			temp = temp + "-"
		} else {
			temp = temp + string(v)
		}
	}
	return Parse(base)
}

// A library method for changing the uuid
func (v *Vertex) changeUUID(n string) {
	uuid := uuidFromString(n)
	v.uuid = uuid
}
*/

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
