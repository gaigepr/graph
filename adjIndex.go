package graph

import (
	. "github.com/pborman/uuid"
)

//
type AdjIndex struct {
	name string
	uuid UUID
	data map[*Vertex][]*Vertex
}

//
func NewAdjIndex() *AdjIndex {
	i := &AdjIndex{
		uuid: NewRandom(),
		data: make(map[*Vertex][]*Vertex),
	}
	return i
}

//
func (a *AdjIndex) Name() string {
	return a.name
}

//
func (a *AdjIndex) UUID() UUID {
	return a.uuid
}

//
func (a *AdjIndex) UUIDString() string {
	return a.uuid.String()
}

//
func (a *AdjIndex) VertexAdd(v *Vertex) {
	return
}

//
func (a *AdjIndex) VertexRemove(v *Vertex) {
	for key, val := range a.data {
		for ix, vertex := range val {
			if vertex == v {
				a.data[key] = append(a.data[key][ix:], a.data[key][:ix+1]...)
				break
			}
		}
	}
}

//
func (a *AdjIndex) EdgeAdd(e *Edge, from, to *Vertex) {
	a.data[from] = append(a.data[from], to)
}

//
func (a *AdjIndex) EdgeRemove(e *Edge, from, to *Vertex) {
	var ix int
	for key, val := range a.data[from] {
		if val == to {
			ix = key
			break
		}
	}
	a.data[from] = append(a.data[from][ix:], a.data[from][:ix+1]...)
}

//
func (a *AdjIndex) AdjacentTo(v *Vertex) []*Vertex {
	return a.data[v]
}
