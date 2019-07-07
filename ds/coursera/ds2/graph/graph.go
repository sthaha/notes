package graph

import (
	"errors"
)

var (
	errNodeExists     = errors.New("node already exists")
	errNodeNotFound   = errors.New("node does not exist")
	errNoConnectivity = errors.New("nodes are not directly connected")
)

type node interface{}

type graph interface {
}

// a : 8
type nodeWeight map[node]int
type connections map[node]nodeWeight

type undirected struct {
	nodes connections
}

func newUndirected() *undirected {
	return &undirected{
		nodes: connections{},
	}
}

func (g *undirected) addNode(x node) error {
	if g.contains(x) {
		return errNodeExists
	}

	g.nodes[x] = nodeWeight{}
	return nil
}

func (g *undirected) size() int {
	return len(g.nodes)
}

func (g *undirected) connect(a, b node, weight int) error {

	g.addNode(a)
	g.addNode(b)
	g.nodes[a][b] = weight
	g.nodes[b][a] = weight
	return nil
}

func (g *undirected) contains(x node) bool {
	_, ok := g.nodes[x]
	return ok
}

func (g *undirected) weightOf(a, b node) (int, error) {

	if !g.contains(a) || !g.contains(b) {
		return 0, errNodeNotFound
	}

	w, ok := g.nodes[a][b]
	if !ok {
		return 0, errNoConnectivity
	}
	return w, nil
}
