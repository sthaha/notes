package graph

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errNodeExists   = errors.New("node already exists")
	errNodeNotFound = errors.New("node does not exist")
	errNotAdjacent  = errors.New("nodes are non adjacent")
)

type node interface{}

type graph interface {
}

// b : 8
type nodeWeight map[node]int

// a : b : 8 -> a is connected to b with a weight of 8
type connections map[node]nodeWeight

type undirected struct {
	conns connections
}

func newUndirected() *undirected {
	return &undirected{
		conns: connections{},
	}
}

func (g *undirected) addNode(x node) error {
	if g.contains(x) {
		return errNodeExists
	}

	g.conns[x] = nodeWeight{}
	return nil
}

func (g *undirected) size() int {
	return len(g.conns)
}

func (g *undirected) addEdge(a, b node, weight int) error {

	g.addNode(a)
	g.addNode(b)
	g.conns[a][b] = weight
	g.conns[b][a] = weight
	return nil
}

func (g *undirected) contains(x node) bool {
	_, ok := g.conns[x]
	return ok
}

func (g *undirected) isAdjascent(a, b node) bool {
	if !g.contains(a) || !g.contains(b) {
		return false
	}
	_, ok := g.conns[a][b]
	return ok
}

func (g *undirected) edgeWeight(a, b node) (int, error) {

	if !g.contains(a) || !g.contains(b) {
		return 0, errNodeNotFound
	}

	w, ok := g.conns[a][b]
	if !ok {
		return 0, errNotAdjacent
	}
	return w, nil
}

func (g *undirected) nodes() []node {
	keys := make([]node, 0, len(g.conns))
	for k := range g.conns {
		keys = append(keys, k)
	}
	return keys
}

func (g *undirected) adjacents(x node) []node {
	if !g.contains(x) {
		return []node{}
	}

	adj := g.conns[x]
	keys := make([]node, 0, len(adj))
	for k := range adj {
		keys = append(keys, k)
	}
	return keys
}

func printGraph(g *undirected) string {
	out := &strings.Builder{}

	for _, n := range g.nodes() {
		for _, adj := range g.adjacents(n) {
			out.WriteString(fmt.Sprintf("%v - %v\n", n, adj))
		}
	}
	return out.String()
}
