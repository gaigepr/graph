package graph

import ()

//
func (g *Graph) IterativeDFS(v *Vertex) map[*Vertex]bool {
	s := NewStack()
	k := make(map[*Vertex]bool) // a hash set to keep track of vertices that have been visited
	s.Push(v)
	for !s.Empty() {
		c, err := s.Pop()
		if err != nil {
			break
		}
		if !k[c] {
			k[c] = true
			neighbors := g.Adjacent(c) // []*Vertex
			for _, vert := range neighbors {
				S.Push(vert)
			}
		}
	}
	return
}

//
func (g *Graph) IterativeBFS(v *Vertex) map[*Vertex]bool {
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
		for _, adj := range g.Adjacent(c) {
			if d[adj] == INFINITY {
				d[adj] = d[c] + 1
				p[adj] = c
				q.Push(adj)
			}
		}
	}
	return
}
