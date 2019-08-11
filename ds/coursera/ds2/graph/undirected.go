package graph

import (
	"fmt"
	"strings"
)

type undirected struct {
	graph connections
}

func (g *undirected) init() {
	if g.graph != nil {
		return
	}
	g.graph = connections{}
}

func (g *undirected) size() int {
	g.init()
	return len(g.graph)
}

func (g *undirected) contains(x node) bool {
	g.init()
	_, ok := g.graph[x]
	return ok
}

func (g *undirected) addNode(x node) error {
	g.init()
	if g.contains(x) {
		return errNodeExists
	}

	g.graph[x] = &nodeset{}
	return nil
}

func (g *undirected) nodes() []node {
	g.init()
	keys := make([]node, 0, len(g.graph))
	for k := range g.graph {
		keys = append(keys, k)
	}
	return keys
}

func (g *undirected) connect(a, b node) error {
	g.addNode(a)
	g.addNode(b)
	g.graph[a].add(b)
	g.graph[b].add(a)
	return nil
}

func (g *undirected) isAdjascent(a, b node) bool {
	if !g.contains(a) || !g.contains(b) {
		return false
	}
	return g.graph[a].has(b)
}

func (g *undirected) adjacents(x node) []node {
	if !g.contains(x) {
		return []node{}
	}

	return g.graph[x].keys()
}

func (g *undirected) toString() string {
	out := &strings.Builder{}

	for _, n := range g.nodes() {
		for _, adj := range g.adjacents(n) {
			out.WriteString(fmt.Sprintf("%v - %v\n", n, adj))
		}
	}
	return out.String()
}
