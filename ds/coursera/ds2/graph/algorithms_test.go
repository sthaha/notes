package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegree(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := &Undirected{}
	g.Connect("a", "b")
	g.Connect("a", "m")

	g.Connect("b", "c")
	g.Connect("m", "c")

	g.Connect("c", "d")

	assert.Equal(t, 3, degree(g, "c"))
	assert.Equal(t, 2, degree(g, "a"))
	assert.Equal(t, 1, degree(g, "d"))
}

func TestMaxDegree(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := &Undirected{}
	g.Connect("a", "b")
	g.Connect("a", "m")

	g.Connect("b", "c")
	g.Connect("m", "c")

	g.Connect("c", "d")

	assert.Equal(t, 3, maxDegree(g))
}
