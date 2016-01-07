package graph

import (
	"log"
	"testing"
)

var g *Graph
var v1 *Vertex
var v2 *Vertex
var v3 *Vertex
var e1 *Edge
var e2 *Edge

// Test the creation of a new Graph, Vertex, and Edge.
func TestNewFuncs(t *testing.T) {
	g = NewGraph("test")
	v1 = NewVertex("type1")
	v2 = NewVertex("type2")
	v3 = NewVertex("type2")
	e1 = NewEdge(10, "knows")
	e2 = NewEdge(1, "knows")
	if g == nil || v1 == nil || v2 == nil || v3 == nil || e1 == nil || e2 == nil {
		log.Fatal("Unable to call NewSomething.")
		t.Fail()
	}
}

// Test adding a Vertex to a Graph.
func TestAddVertex(t *testing.T) {
	success := g.AddVertex(v1)
	if !success {
		log.Fatal("Failed to add vertex v1.")
		t.Fail()
	}
	suc2 := g.AddVertex(v2)
	if !suc2 {
		log.Fatal("Failed to add vertex v2.")
		t.Fail()
	}
	suc3 := g.AddVertex(v3)
	if !suc3 {
		log.Fatal("Failed to add vertex v3.")
		t.Fail()
	}
}

// Test adding an Edge to a Graph.
func TestAddEdge(t *testing.T) {
	success := g.AddEdge(e1, v1, v2)
	if !success {
		log.Fatal("Failed to add edge e1 from v1 to v2.")
		t.Fail()
	}

	success = g.AddEdge(e2, v2, v3)
	if !success {
		log.Fatal("Failed to add edge e2 from v2 to v3.")
		t.Fail()
	}
}

// Test existence checking functions.
func TestExistsFuncs(t *testing.T) {
	f1 := g.VertexExists(v1)
	f2 := g.VertexExists(v2)
	f3 := g.VertexExists(v3)
	f4 := g.EdgeExists(e1)
	f5 := g.EdgeExists(e2)
	if !(f1 && f2 && f3 && f4 && f5) {
		log.Fatal("Failed to locate all entites.")
		t.Fail()
	}
}

//
func TestAdjacency(t *testing.T) {
	for _, v := range g.Adjacent(v1) {
		log.Println("v1 adjacent to v2:", v.UUIDString())
	}

	for _, v := range g.Adjacent(v2) {
		log.Println("v2 adjacent to v3:", v.UUIDString())
	}

	for _, v := range g.Adjacent(v3) {
		log.Println("v3 adjacent to:", v.UUIDString())
	}
}

//
func TestReverseAdjacency(t *testing.T) {
	for _, v := range g.ReverseAdjacent(v1) {
		log.Println("v1 revadj to none:", v.UUIDString())
	}

	for _, v := range g.ReverseAdjacent(v2) {
		log.Println("v2 revadj to v1:", v.UUIDString())
	}

	for _, v := range g.ReverseAdjacent(v3) {
		log.Println("v3 revadj to v2:", v.UUIDString())
	}
}

//
func TestVertices(t *testing.T) {
	a := g.Vertices()
	if len(a) != 3 {
		log.Fatal("Wrong number of vertices.")
	}
	log.Println("vertices:", a)
}

//
func TestEdges(t *testing.T) {
	a := g.Edges()
	if len(a) != 2 {
		log.Fatal("Wrong number of edges.")
	}
	log.Println("edges:", a)
}

//
func TestRemove(t *testing.T) {
	//var vc int = g.VertexCount()
	//var ec int = g.EdgeCount()
	var fail bool = false
	log.Println("Vertex Count:", g.VertexCount())
	log.Println("Edge Count:", g.EdgeCount())

	// This will also remove e2
	if !g.RemoveVertex(v3) {
		log.Fatal("Failed to delete v3.")
		fail = true
	}
	log.Println("Deleted e2, then v3.")
	log.Println("Vertex Count:", g.VertexCount())
	log.Println("Edge Count:", g.EdgeCount())

	// This should return false becase the previous try deleted it
	// g.RemoveEdge always returns true bevause it is a simple delete() call
	if !g.RemoveEdge(e2, v2, v3) {
		log.Fatal("failed to delete e2.")
		fail = true
	}

	if !g.RemoveEdge(e1, v1, v2) {
		log.Fatal("failed to delete e1.")
		fail = true
	}
	log.Println("Deleted e1.")

	log.Println("Vertex Count:", g.VertexCount())
	log.Println("Edge Count:", g.EdgeCount())

	if !g.RemoveVertex(v1) {
		log.Fatal("failed to delete v1.")
		fail = true
	}

	if !g.RemoveVertex(v2) {
		log.Fatal("failed to delete v2.")
		fail = true
	}
	log.Println("Deleted v1, then v2.")

	log.Println("Vertex Count:", g.VertexCount())
	log.Println("Edge Count:", g.EdgeCount())

	if fail {
		t.Fail()
	}
}
