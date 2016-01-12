package graph

import (
	"log"

	. "github.com/pborman/uuid"
)

// Could make it so that the mapping of adjacency and revAdjacency is to weights.
// Also the adj maps don't account for multiple types of Edges connecting Vertices.
type Graph struct {
	vertices     map[string]*Vertex
	edges        map[string]*Edge
	adjacency    map[*Vertex]map[*Vertex]int
	revAdjacency map[*Vertex]map[*Vertex]int
}

//
func NewGraph() *Graph {
	var g *Graph = new(Graph)
	g.vertices = make(map[string]*Vertex)
	g.edges = make(map[string]*Edge)
	g.adjacency = make(map[*Vertex]map[*Vertex]int)
	g.revAdjacency = make(map[*Vertex]map[*Vertex]int)
	return g
}

// Get a *Vertex by UUID from g.
func (g *Graph) Vertex(u UUID) (*Vertex, bool) {
	v, ex := g.vertices[u.String()]
	return v, ex
}

// Get a *Edge by UUID from g.
func (g *Graph) Edge(u UUID) (*Edge, bool) {
	e, ex := g.edges[u.String()]
	return e, ex
}

// Add a Vertex v to g.
// This is a constant time operation.
func (g *Graph) AddVertex(v *Vertex) bool {
	if _, exists := g.Vertex(v.UUID()); !exists { // v does not exists in g
		g.vertices[v.UUIDString()] = v // copy the contents of v into the graph
		vert, ex := g.Vertex(v.UUID())
		if !ex {
			log.Fatal("AddVertex: Failed to add v to g.")
		} else {
			log.Println("Succeeded", vert.UUIDString())
		}
		// Update adjacency
		g.adjacency[vert] = make(map[*Vertex]int)
		// Update reverse adjacency
		g.revAdjacency[vert] = make(map[*Vertex]int)
		return true
	}
	return false
}

// Add an edge from source to destination.
// This is a constant time operation.
func (g *Graph) AddEdge(e *Edge, src, dst *Vertex) bool {
	// can lock the global instances of e from src to dst?
	gsrc, ex1 := g.Vertex(src.UUID())
	gdst, ex2 := g.Vertex(dst.UUID())

	// check for the existence of both Vertices
	if !ex1 || !ex2 {
		return false
	}

	if gsrc.Equal(src) && gdst.Equal(dst) {
		// The pointers passed in DO point to the Vertex they name in g.
		e.from = gsrc.uuid
		e.to = gdst.uuid
		g.edges[e.UUIDString()] = e // copy the contents of e into g.

		// Update adjacency
		g.adjacency[gsrc][gdst] = e.weight
		// Update reverse adjacency
		g.revAdjacency[gdst][gsrc] = e.weight

		return true
	}
	return false
}

// Remove an Edge e between Vertex src and Vertex dst from g.
// always returns true.
// This function deletes 1 entry in each of g.adjacency and g.revAdjacency.
// Finally it deletes e from g.edges.
// This deleteion does not account for multiple Edges between two Vertices.
// Constant time operation.
func (g *Graph) RemoveEdge(e *Edge, src, dst *Vertex) bool {
	// src is not longer adjacent to dst via e.relation
	delete(g.adjacency[src], dst)
	delete(g.revAdjacency[dst], src)
	// remove e from g.edges
	delete(g.edges, e.UUIDString())
	return true
}

// Remove a Vertex v from g also removing all Edges with their source at v.
// Runs in time proportional to:
// (number of Vertices to which v is adjacent)
// + (number of Vertices that are adjacent to v)
// + (number of edges in/out of v)
func (g *Graph) RemoveVertex(v *Vertex) bool {
	vuuid := v.UUIDString()
	gv, _ := g.Vertex(v.UUID())
	if v == gv { // v points to a vertex in g.
		// Remove all connected edges
		for _, edge := range g.edges {
			efromuuid := edge.from
			etouuid := edge.to
			if efromuuid.String() == vuuid || etouuid.String() == vuuid {
				// Either the from or to of an edge is connected to v.
				from, _ := g.Vertex(efromuuid)
				to, _ := g.Vertex(etouuid)
				removed := g.RemoveEdge(edge, from, to)
				if !removed {
					log.Fatal("ERROR: RemoveVertex: Failed to delete edge:", edge.UUIDString(), "from:", efromuuid, "to:", etouuid)
				}
			}
		}

		// Update adjacency
		for vert, _ := range g.revAdjacency { // For each Vertex vert from which Vertex v has an incoming edge
			delete(g.adjacency[vert], v) // removes v from vert's adjacency map
		}
		// Update reverse adjacency
		for vert, _ := range g.adjacency { // For each Vertex vert to which Vertex v has outgoing edge
			delete(g.revAdjacency[vert], v) // remove v from vert's reverse adjacency map
		}

		delete(g.adjacency, v)             // delete v's entry from g.adjacency
		delete(g.revAdjacency, v)          // delete v's entry from g.revAdjacency
		delete(g.vertices, v.UUIDString()) // delete v from g.vertices
		return true
	}
	return false
}

// Returns a mapping of UUID.String() to *Vertex in g.
func (g *Graph) Vertices() map[string]*Vertex {
	m := make(map[string]*Vertex)
	for k, v := range g.vertices {
		m[k] = v
	}
	return m
}

// Returns a mapping of UUID.String() to *Edge in g.
func (g *Graph) Edges() map[string]*Edge {
	m := make(map[string]*Edge)
	for k, e := range g.edges {
		m[k] = e
	}
	return m
}

// Return the adjacency map of v in g.
func (g *Graph) Adjacent(v *Vertex) map[*Vertex]int {
	m, _ := g.adjacency[v]
	return m
}

// Return the reverse adjacency map of v in g.
func (g *Graph) ReverseAdjacent(v *Vertex) map[*Vertex]int {
	m, _ := g.revAdjacency[v]
	return m
}

// Returns a mapping of UUID.String() to *Edge in g where *Edge.to == v.
func (g *Graph) EdgesTo(v *Vertex) map[string]*Edge {
	m := make(map[string]*Edge)
	for k, e := range g.edges {
		if e.to.String() == v.UUIDString() {
			m[k] = e
		}
	}
	return m
}

// Returns a mapping of UUID.String() to *Edge in g where *Edge.from == v.
func (g *Graph) EdgesFrom(v *Vertex) map[string]*Edge {
	m := make(map[string]*Edge)
	for k, e := range g.edges {
		if e.from.String() == v.UUIDString() {
			m[k] = e
		}
	}
	return m
}

// Returns a map that is the combination of g.EdgesTo() and g.EdgesFrom().
func (g *Graph) EdgesOf(v *Vertex) map[string]*Edge {
	m := g.EdgesFrom(v)
	for k, e := range g.EdgesTo(v) {
		m[k] = e
	}
	return m
}

// Get the In-degree of a Vertex.
// A return value of -1 indicates the Vertex has no entry in the revAdjacency map.
// -1 implies v does not exist.
func (g *Graph) InDegree(v *Vertex) int {
	if _, ex := g.revAdjacency[v]; !ex {
		if _, ex1 := g.vertices[v.UUIDString()]; ex1 {
			return 0
		}
		return -1
	}
	return len(g.revAdjacency[v])
}

// Get the Out-degree of a Vertex.
// A return value of -1 indicates the Vertex has no entry in the adjacency map.
// -1 implies v does not exist.
func (g *Graph) OutDegree(v *Vertex) int {
	if _, ex := g.adjacency[v]; !ex {
		if _, ex1 := g.vertices[v.UUIDString()]; ex1 {
			return 0
		}
		return -1
	}
	return len(g.adjacency[v])
}
