package graph

import (
	"fmt"
	"strings"
)

type Undirected struct {
	graph connections
}

func (g *Undirected) init() {
	if g.graph != nil {
		return
	}
	g.graph = connections{}
}

func (g *Undirected) Size() int {
	g.init()
	return len(g.graph)
}

func (g *Undirected) Contains(x Node) bool {
	g.init()
	_, ok := g.graph[x]
	return ok
}

func (g *Undirected) AddNode(x Node) error {
	g.init()
	if g.Contains(x) {
		return ErrNodeExists
	}

	g.graph[x] = &nodeset{}
	return nil
}

func (g *Undirected) Nodes() []Node {
	g.init()
	keys := make([]Node, 0, len(g.graph))
	for k := range g.graph {
		keys = append(keys, k)
	}
	return keys
}

func (g *Undirected) Connect(a, b Node) error {
	g.AddNode(a)
	g.AddNode(b)
	g.graph[a].add(b)
	g.graph[b].add(a)
	return nil
}

func (g *Undirected) IsAdjascent(a, b Node) bool {
	if !g.Contains(a) || !g.Contains(b) {
		return false
	}
	return g.graph[a].has(b)
}

func (g *Undirected) Adjacents(x Node) []Node {
	if !g.Contains(x) {
		return []Node{}
	}

	return g.graph[x].keys()
}

func (g *Undirected) ToString() string {
	out := &strings.Builder{}

	for _, n := range g.Nodes() {
		for _, adj := range g.Adjacents(n) {
			out.WriteString(fmt.Sprintf("%v - %v\n", n, adj))
		}
	}
	return out.String()
}
