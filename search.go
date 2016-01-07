package graph

import ()

//
func (g *Graph) IterativeDFS(v *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex) {
	k := make(map[*Vertex]bool)    // a hash set to keep track of vertices that have been visited
	d := make(map[*Vertex]int)     // a map of *Vertex to distance from v
	p := make(map[*Vertex]*Vertex) // maps *Vertex to parent *Vertex

	for _, vert := range g.Vertices() {
		k[vert] = false
		d[vert] = INFINITY
		p[vert] = nil
	}

	d[v] = 0 // The distance of the start node is 0

	s := NewStack()
	s.Push(v)

	for !s.Empty() {
		c, err := s.Pop()
		if err != nil {
			break
		}
		if !k[c] {
			k[c] = true
			neighbors := g.Adjacent(c) // []*Vertex
			for _, adj := range neighbors {
				if d[adj] == INFINITY {
					d[adj] = d[c] + 1
					p[adj] = c
				}
				s.Push(adj)
			}
		}
	}
	return d, p
}

//
func (g *Graph) IterativeBFS(v *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex) {
	d := make(map[*Vertex]int)     // vertex to weight distance from root
	p := make(map[*Vertex]*Vertex) // maps a *Vertex to its parent *Vertex in BFS terms
	for _, vert := range g.Vertices() {
		d[vert] = INFINITY
		p[vert] = nil
	}
	q := NewQueue()
	d[v] = 0
	q.Push(v)
	for q.Length() > 0 {
		c := q.Poll()
		if c == nil {
			break
		}
		for _, adj := range g.Adjacent(c) {
			if d[adj] == INFINITY {
				d[adj] = d[c] + 1
				p[adj] = c
				q.Push(adj)
			}
		}
	}
	return d, p
}
