package graph

import (
	//

	. "github.com/pborman/uuid"
)

//
type Index interface {
	Name() string
	UUID() UUID
	UUIDString() string
	VertexAdd(v *Vertex)
	VertexRemove(v *Vertex)
	EdgeAdd(e *Edge, from, to *Vertex)
	EdgeRemove(e *Edge, from, to *Vertex)
}

//
type Graph struct {
	name             string
	vertices         map[string]Vertex
	edges            map[string]Edge
	adjacency        map[*Vertex][]*Vertex
	reverseAdjacency map[*Vertex][]*Vertex

	vertexType       map[string][]*Vertex                   // "vertex-type" -> []*Vertex x: x contains all vertices of "vertex-type"
	EdgeType         map[string][]*Edge                     // "edge-type" -> []*Edge e: e contains all edges of "edge-type"
	outgoingEdgeType map[string]map[*Vertex][]*Edge         // "edge-type" -> *Vertex x -> []*Edge e: e contains all Edges of "edge-type" with their origin at x.
	incomingEdgeType map[string]map[*Vertex][]*Edge         // "edge-type" -> *Vertex x -> []Edge e: e contains all Edges of "edge-type" with their destination at x.
	priorityEdgeType map[string]map[*Vertex]map[*Vertex]int // "edge-type" -> *Vertex x -> *Vertex y -> weight i: i is the lowest weight of "edge-type" from x to y.

	//indexes   map[string]Index
}

//
type DegreeStat struct {
	MinInDegree    int
	MaxInDegree    int
	MinOutDegree   int
	MaxOutDegree   int
	MinInDegreeID  UUID
	MaxInDegreeID  UUID
	MinOutDegreeID UUID
	MaxOutDegreeID UUID
	AvgOutDegree   float64
}

//
func NewGraph(name string) *Graph {
	g := &Graph{
		name:             name,
		vertices:         make(map[string]Vertex),
		edges:            make(map[string]Edge),
		adjacency:        make(map[*Vertex][]*Vertex),
		reverseAdjacency: make(map[*Vertex][]*Vertex),
		//indexes:   make(map[string]Index),
	}

	// Load in all desired indexes
	//g.indexes["adjacency"] = NewAdjIndex()

	return g
}

//
func (g *Graph) Name() string {
	return g.name
}

//
func (g *Graph) VertexCount() int {
	return len(g.vertices)
}

//
func (g *Graph) EdgeCount() int {
	return len(g.edges)
}

//
func (g *Graph) Adjacent(v *Vertex) []*Vertex {
	return g.adjacency[v]
}

//
func (g *Graph) ReverseAdjacent(v *Vertex) []*Vertex {
	return g.reverseAdjacency[v]
}

//
func (g *Graph) VertexExists(v *Vertex) bool {
	_, exists := g.vertices[v.uuid.String()]
	return exists
}

//
func (g *Graph) EdgeExists(e *Edge) bool {
	_, exists := g.edges[e.UUIDString()]
	return exists
}

//
func (g *Graph) VertexByUUID(uuid UUID) *Vertex {
	u := uuid.String()
	var v Vertex = g.vertices[u]
	var vptr *Vertex = &v
	return vptr
}

//
func (g *Graph) EdgeByUUID(uuid UUID) *Edge {
	u := uuid.String()
	var e Edge = g.edges[u]
	var eptr *Edge = &e
	return eptr
}

//
func (g *Graph) AddVertex(v *Vertex) bool {
	if g.VertexExists(v) {
		return false
	}

	// WARN: Potentially unsafe operation!
	// If the user did not use NewVertex(...) this could be nil
	// However since the UUID string will be "" at worst, it won't be too bad.
	g.vertices[v.UUIDString()] = *v

	// TODO: Update indexes
	// adjacency nil
	// reverseAdjacency nil

	return true
}

//
func (g *Graph) AddEdge(e *Edge, from, to *Vertex) bool {
	if !g.EdgeExists(e) && g.VertexExists(from) && g.VertexExists(to) {
		e.from = from.uuid
		e.to = to.uuid
		g.edges[e.UUIDString()] = *e

		// TODO: Update indexes
		// adjacency
		g.adjacency[from] = append(g.adjacency[from], to)
		// reverseAdjacency
		g.reverseAdjacency[to] = append(g.reverseAdjacency[to], from)

		return true
	}
	return false
}

// Always returns true.
func (g *Graph) RemoveEdge(e *Edge, from, to *Vertex) bool {
	// TODO: Update indexes
	// adjacency
	for key, val := range g.adjacency[from] {
		if val == to {
			g.adjacency[from] = append(g.adjacency[from][key:], g.adjacency[from][:key+1]...)
			break
		}
	}
	// reverseAdjacency
	for key, val := range g.reverseAdjacency[to] {
		if val == from {
			g.reverseAdjacency[to] = append(g.reverseAdjacency[to][key:], g.reverseAdjacency[to][:key+1]...)
			break
		}
	}

	delete(g.edges, e.UUIDString())
	return true
}

//
func (g *Graph) RemoveVertex(v *Vertex) bool {
	if g.VertexExists(v) {
		// Remove all connected edges
		for _, edge := range g.edges {
			// If either Vertex of this edge are equal to the Vertex being deleted:
			if edge.from.String() == v.UUIDString() || edge.to.String() == v.UUIDString() {
				from := g.VertexByUUID(edge.from)
				to := g.VertexByUUID(edge.to)
				g.RemoveEdge(&edge, from, to)
			}
		}

		// TODO: Update indexes
		// adjacency
		for key, val := range g.adjacency { // for each adjacency []*Vertex entry
			for ix, vertex := range val { // for each *Vertex in []*Vertex
				if vertex == v { // Vertex to delete is in this []*Vertex
					g.adjacency[key] = append(g.adjacency[key][ix:], g.adjacency[key][:ix+1]...)
					break
				}
			}
		}
		// reverseAdjacency
		for key, val := range g.reverseAdjacency { // for each reverseAdjacency []*Vertex entry
			for ix, vertex := range val { // for each *Vertex in []*Vertex
				if vertex == v { // Vertex to delete is in this []*Vertex
					g.reverseAdjacency[key] = append(g.reverseAdjacency[key][ix:], g.reverseAdjacency[key][:ix+1]...)
					break
				}
			}
		}

		delete(g.vertices, v.UUIDString())
		return true
	}
	return false
}

//
func (g *Graph) Vertices() []*Vertex {
	var vertices []*Vertex
	for _, vertex := range g.vertices {
		vertices = append(vertices, &vertex)
	}
	return vertices
}

//
func (g *Graph) VerticesByUUID(uuids ...UUID) []*Vertex {
	var vertices []*Vertex
	for _, uuid := range uuids {
		for key, vert := range g.vertices {
			if key == uuid.String() {
				vertices = append(vertices, &vert)
			}
		}
	}
	return vertices
}

//
func (g *Graph) Edges() []*Edge {
	var edges []*Edge
	for _, edge := range g.edges {
		edges = append(edges, &edge)
	}
	return edges
}

//
func (g *Graph) EdgesByUUID(uuids ...UUID) []*Edge {
	var edges []*Edge
	for _, uuid := range uuids {
		for key, edge := range g.edges {
			if key == uuid.String() {
				edges = append(edges, &edge)
			}
		}
	}
	return edges
}

//
func (g *Graph) InDegree(v *Vertex) int {
	if _, e := g.reverseAdjacency[v]; !e {
		return 0
	}
	return len(g.reverseAdjacency[v])
}

//
func (g *Graph) OutDegree(v *Vertex) int {
	if _, e := g.adjacency[v]; !e {
		return 0
	}
	return len(g.adjacency[v])
}

//
func (g *Graph) DegreeStats() DegreeStat {
	s := DegreeStat{}
	return s
}
