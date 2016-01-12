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
	edgesOut map[*Vertex]int    = make(map[*Vertex]int)
	edgesIn  map[*Vertex]int    = make(map[*Vertex]int)
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
	log.Println(records)

	for _, r := range records {
		a, _ := strconv.Atoi(r[1])
		from := NewVertex(r[0])
		edge := NewEdge(a)
		to := NewVertex(r[2])

		if _, ex := vl[r[0]]; !ex {
			// if from is not already in the graph as per our record
			vl[r[0]] = from
			vertices[from.UUIDString()] = from
			added := g.AddVertex(from)
			if !added {
				log.Fatal("Failed to add vertex .")
			}
		} else {
			log.Println("vertex exists.", r[0])
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
			}
		} else {
			log.Println("vertex exists.", r[2])
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
		}
	}
}

func TestCollectionOperations(t *testing.T) {
	a := g.Vertices()
	if len(a) != 4 {
		log.Fatal("Not all Vertices present in g.")
	}

	b := g.Edges()
	if len(b) != 5 {
		log.Fatal("Not all edges present in g.")
	}
}

func TestDeletions(t *testing.T) {
	// Test Deletion of Edge
	//g.RemoveEdge()
	// Delete Vertex with Known Edge(s)
	// Make sure Edge(s) and the Vertex were deleted
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
