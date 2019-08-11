package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	g := newUndirected()
	assert.Equal(t, 0, g.size())

	g.addNode(1)
	assert.Equal(t, 1, g.size())

	g.addNode(2)
	assert.Equal(t, 2, g.size())
}

func TestAddNode_same_error(t *testing.T) {
	g := newUndirected()
	assert.Equal(t, 0, g.size())

	g.addNode(1)
	assert.Equal(t, 1, g.size())

	// adding the same node results in error
	err := g.addNode(1)
	assert.Error(t, err, "adding same node is an error")

	// size does not change
	assert.Equal(t, 1, g.size())
}

func TestContains(t *testing.T) {
	g := newUndirected()
	assert.Equal(t, 0, g.size())
	assert.False(t, g.contains(1))

	g.addNode(1)
	assert.True(t, g.contains(1))
	assert.False(t, g.contains(2))
}

func TestConnect(t *testing.T) {
	g := newUndirected()

	g.addNode("alice")
	g.addNode("bob")
	assert.Equal(t, 2, g.size())

	g.connect("alice", "bob")
	assert.Equal(t, 2, g.size())
}

func TestGraph_adjacents(t *testing.T) {
	g := newUndirected()
	// a - b - c - d
	//  \    /
	//   - m
	g.connect("a", "b")
	g.connect("b", "c")
	g.connect("c", "d")

	g.connect("a", "m")

	assert.Equal(t, 2, len(g.adjacents("a")))
	assert.Equal(t, 1, len(g.adjacents("d")))
	assert.Equal(t, 2, len(g.adjacents("b")))

	g.connect("m", "c")
	assert.Equal(t, 3, len(g.adjacents("c")))

}

func TestGraph_self_loop(t *testing.T) {
	g := newUndirected()
	// a - b - c - d
	//  \    /
	//   - m
	g.connect("a", "a")
	assert.Equal(t, 1, len(g.adjacents("a")))

}
