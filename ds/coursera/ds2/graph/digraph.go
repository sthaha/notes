package graph

import (
	"fmt"
	"strings"
)

type Directed struct {
	Undirected
}

func (g *Directed) Connect(a, b Node) error {
	g.AddNode(a)
	g.AddNode(b)
	g.graph[a].add(b)
	return nil
}

func (g *Directed) ToString() string {
	out := &strings.Builder{}

	for _, n := range g.Nodes() {
		for _, adj := range g.Adjacents(n) {
			out.WriteString(fmt.Sprintf("%v -> %v\n", n, adj))
		}
	}
	return out.String()
}
