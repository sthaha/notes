package graph

import (
	"fmt"
	"strings"
)

type directed struct {
	undirected
}

func (g *directed) connect(a, b node) error {
	g.addNode(a)
	g.addNode(b)
	g.graph[a].add(b)
	return nil
}

func (g *directed) toString() string {
	out := &strings.Builder{}

	for _, n := range g.nodes() {
		for _, adj := range g.adjacents(n) {
			out.WriteString(fmt.Sprintf("%v -> %v\n", n, adj))
		}
	}
	return out.String()
}
