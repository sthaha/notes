package graph

import (
	"errors"
)

var (
	ErrNodeExists   = errors.New("node already exists")
	ErrNodeNotFound = errors.New("node does not exist")
	ErrNotAdjacent  = errors.New("nodes are non adjacent")
)

type Node interface{}

// b : 8
//type nodeWeight map[node]int

// a : b : 8 -> a is connected to b with a weight of 8
type connections map[Node]*nodeset

type Graph interface {
	AddNode(x Node) error
	Size() int
	Nodes() []Node
	Adjacents(x Node) []Node

	Contains(x Node) bool
	ToString() string
}

// ## algorithms
// 1. is a and c connected
// 2. shorted path between a and b
// 3. is there a cycle in the graph
// 4. Eulers tour of graph -> cycle that visit each edge only once
// 5. hamilton tour
// ### Connectivity
// 6. are all vertices connected
// 7. MST ?? what is the best way to connect all vertices
// 8. is there a vertex whose removal disconnects the graph
