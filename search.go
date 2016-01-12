package graph

//

//import "log"
//import "math"

// The return order is (map[child]parent, map[destination]cost)
//
func (g *Graph) Dijkstras(root *Vertex) (map[*Vertex]*Vertex, map[*Vertex]int) {
	ctop := make(map[*Vertex]*Vertex) // A mapping of child *Vertex to Parent *Vertex
	vtow := make(map[*Vertex]int)     // A mapping of *Vertex to cost from root

	return ctop, vtow
}

// The return order is (map[child]parent, map[destination]cost)
// This functions returns: (child -> parent *Vertex), (*Vertex -> cost).
// This allows a user to construct the shortest path for any *Vertex from root.
func (g *Graph) BFS(root *Vertex) (map[*Vertex]*Vertex, map[*Vertex]int) {
	ctop := make(map[*Vertex]*Vertex) // A mapping of child *Vertex to Parent *Vertex
	vtow := make(map[*Vertex]int)     // A mapping of *Vertex to cost from root
	q := new(Queue)
	for _, vert := range g.Vertices() {
		// Set the distance to initial values
		vtow[vert] = -1
	}

	vtow[root] = 0
	q.Push(root)
	for q.Length() != 0 {
		current := q.Poll()
		// sort adjacent!!
		for vertex, weight := range g.Adjacent(current) {
			if vtow[vertex] == -1 {
				// The distance from root to current
				// + the cost from current to vertex
				vtow[vertex] = vtow[current] + weight
				ctop[vertex] = current
				q.Push(vertex)
			}
		}
	}
	return ctop, vtow
}

// The return order is (map[child]parent, map[destination]cost)
func (g *Graph) DFS(root *Vertex) { // (map[*Vertex]*Vertex, map[*Vertex]int) {
	//ctop := make(map[*Vertex]*Vertex) // A mapping of child *Vertex to parent *Vertex
	//vtow := make(map[*Vertex]int)     // A mapping of *Vertex to cost from root
	discovered := make(map[*Vertex]bool)
	s := new(Stack)

	//vtow[root] = 0
	discovered[root] = true
	s.Push(root)
	for !s.Empty() {
		current, _ := s.Pop()
		if _, seen := discovered[current]; !seen {
			discovered[current] = true
			for vertex, _ := range g.Adjacent(current) {
				s.Push(vertex)
			}
		}
	}
	//return ctop, vtow
}
