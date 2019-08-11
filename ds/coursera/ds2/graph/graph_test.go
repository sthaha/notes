package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	g := &Undirected{}
	assert.Equal(t, 0, g.Size())

	g.AddNode(1)
	assert.Equal(t, 1, g.Size())

	g.AddNode(2)
	assert.Equal(t, 2, g.Size())
}

func TestAddNode_same_error(t *testing.T) {
	g := &Undirected{}
	assert.Equal(t, 0, g.Size())

	g.AddNode(1)
	assert.Equal(t, 1, g.Size())

	// adding the same node results in error
	err := g.AddNode(1)
	assert.Error(t, err, "adding same node is an error")

	// Size does not change
	assert.Equal(t, 1, g.Size())
}

func TestContains(t *testing.T) {
	g := &Undirected{}
	assert.Equal(t, 0, g.Size())
	assert.False(t, g.Contains(1))

	g.AddNode(1)
	assert.True(t, g.Contains(1))
	assert.False(t, g.Contains(2))
}

func TestConnect(t *testing.T) {
	g := &Undirected{}

	g.AddNode("alice")
	g.AddNode("bob")
	assert.Equal(t, 2, g.Size())

	g.Connect("alice", "bob")
	assert.Equal(t, 2, g.Size())
}

func TestGraph_adjacents(t *testing.T) {
	g := &Undirected{}
	// a - b - c - d
	//  \    /
	//   - m
	g.Connect("a", "b")
	g.Connect("b", "c")
	g.Connect("c", "d")

	g.Connect("a", "m")

	assert.Equal(t, 2, len(g.Adjacents("a")))
	assert.Equal(t, 1, len(g.Adjacents("d")))
	assert.Equal(t, 2, len(g.Adjacents("b")))

	g.Connect("m", "c")
	assert.Equal(t, 3, len(g.Adjacents("c")))

}

func TestGraph_self_loop(t *testing.T) {
	g := &Undirected{}
	// a - b - c - d
	//  \    /
	//   - m
	g.Connect("a", "a")
	assert.Equal(t, 1, len(g.Adjacents("a")))

}
