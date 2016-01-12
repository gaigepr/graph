package graph

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

var (
	// A pointer for a graph
	g *Graph = NewGraph()

	vl       map[string]*Vertex = make(map[string]*Vertex)
	el       map[string]*Edge   = make(map[string]*Edge)
	vertices map[string]*Vertex = make(map[string]*Vertex)
	edges    map[string]*Edge   = make(map[string]*Edge)
	//edgesOut map[*Vertex]int    = make(map[*Vertex]int)
	//edgesIn  map[*Vertex]int    = make(map[*Vertex]int)
	// adjCount  map[*Vertex]int    = make(map[*Vertex]int)
	// radjCount map[*Vertex]int    = make(map[*Vertex]int)
)

// A function to load a sample graph from csv
func LoadGraphFromCSV(filename string) [][]string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to load CSV file:", filename)
	}

	r := csv.NewReader(strings.NewReader(string(data)))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV:", err)
	}
	return records
}

//
func TestGraphFromCSV(t *testing.T) {

	records := LoadGraphFromCSV("test-graph.csv")
	log.Println("Loading:", records)

	log.Println("Adding Edges and Vertices.")
	for _, r := range records {
		a, _ := strconv.Atoi(r[1])
		from := NewVertex(r[0], 0)
		edge := NewEdge(a)
		to := NewVertex(r[2], 1)

		if _, ex := vl[r[0]]; !ex {
			// if from is not already in the graph as per our record
			vl[r[0]] = from
			vertices[from.UUIDString()] = from
			added := g.AddVertex(from)
			if !added {
				log.Fatal("Failed to add vertex .")
			} else {
				log.Println("Added Vertex:", from.Type())
			}
		} else {
			//log.Println("vertex exists.", r[0])
			for _, v := range vertices {
				if v.Type() == r[0] {
					from, _ = g.Vertex(v.UUID())
				}
			}
		}

		if _, ex := vl[r[2]]; !ex {
			// if from is not already in the graph as per our record
			vl[r[2]] = to
			vertices[to.UUIDString()] = to
			added := g.AddVertex(to)
			if !added {
				log.Fatal("Failed to add vertex .")
			} else {
				log.Println("Added Vertex:", to.Type())
			}
		} else {
			//log.Println("vertex exists.", r[2])
			for _, v := range vertices {
				if v.Type() == r[2] {
					to, _ = g.Vertex(v.UUID())
				}
			}
		}

		el[r[1]] = edge
		edges[edge.UUIDString()] = edge
		added := g.AddEdge(edge, from, to)
		if !added {
			log.Fatal("Failed to add edge.")
		} else {
			log.Println("Added Edge:", edge.Weight(), "from:", from.Type(), "to:", to.Type())
		}
	}
}

func TestCollectionOperations(t *testing.T) {
	log.Println("Testing Vertices() operation.")
	a := g.Vertices()
	if len(a) != 4 {
		log.Fatal("Not all Vertices present in g.")
	}

	log.Println("Testing Edges() operation.")
	b := g.Edges()
	if len(b) != 6 {
		log.Fatal("Not all edges present in g.")
	}
}

func TestDeletions(t *testing.T) {
	log.Println("Deletion tests.")
	// Delete Vertex with Known Edge(s)
	r1 := g.RemoveVertex(vl["b"]) // THIS will remove 3 edges and 1 vertex. b has 3 connections.
	if !r1 {
		log.Fatal("failed to remove b.")
	}
	a := g.Vertices()
	if len(a) != 3 {
		log.Fatal("Too many Vertices present in g.")
	}
	b := g.Edges()
	if len(b) != 3 {
		log.Fatal("Too many edges present in g.")
	}
	log.Println("Removed b and all associated edges.")

	va, ex := g.Vertex(vl["a"].UUID())
	if !ex {
		log.Fatal("a did not exist in g.")
	}

	if len(g.Adjacent(va)) != 1 {
		log.Fatal("Too many adjacency entries for a in g.")
	}
	// Test Deletion of Edge
	vae := g.EdgesOf(va)
	if len(vae) != 2 {
		log.Fatal("Wrong number of in/out edges for a.")
	}
	for _, edge := range vae {
		to, _ := g.Vertex(edge.to)
		removed := g.RemoveEdge(edge, va, to)
		if !removed {
			log.Fatal("Failed to remove edge from:", va.Type(), "to:", to.Type())
		}
	}
	log.Println("Removed all edges to and from a.")
	// Make sure Edge(s) and the Vertex were deleted
	for str, vert := range vl {
		_, ex := g.Vertex(vert.UUID())
		log.Println("Vertex:", str, "exists:", ex)
	}

	count := 0
	for _, edge := range el {
		e, ex := g.Edge(edge.UUID())
		if !ex {
			continue
		}
		count += 1
		from, exfrom := g.Vertex(e.from)
		to, exto := g.Vertex(e.to)
		if exfrom && exto {
			log.Println("edge:", e.Weight(), "From:", from.Type(), "to:", to.Type())
		} else {
			log.Fatal("Not both Vertices exists for edge:", e.Weight())
		}
	}
	if count > 1 {
		log.Fatal("Too many edges left. Should be 1.")
	} else {
		log.Println("Done with deletion tests.")
	}
}

func TestCopyOperations(t *testing.T) {
	//var t *Graph = g.Transpose()
	// Compare t to g, all edges (x -> y) in t should be (y -> x) in g
	//var c *Graph = g.Copy()
	// Compare c to g, all edges (x -> y) in c should be (x -> x) in g
}

func TestGraphizInput(t *testing.T) {
	// test g.NewGraphiz(str) -> *Graph
}

func TestGraphizOutput(t *testing.T) {
	// Test g.Graphiz() -> string
}
