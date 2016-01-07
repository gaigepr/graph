package graph

import (
	. "github.com/pborman/uuid"
)

//
type Vertex struct {
	uuid       UUID
	vtype      string
	properties interface{}
}

//
func NewVertex(vtype string) *Vertex {
	return &Vertex{
		uuid:  NewRandom(),
		vtype: vtype,
	}
}

//
func (v *Vertex) UUID() UUID {
	return v.uuid
}

//
func (v *Vertex) UUIDString() string {
	return v.uuid.String()
}
