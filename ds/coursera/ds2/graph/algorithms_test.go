package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegree(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := newUndirected()
	g.connect("a", "b", 1)
	g.connect("a", "m", 1)

	g.connect("b", "c", 1)
	g.connect("m", "c", 1)

	g.connect("c", "d", 1)

	assert.Equal(t, 3, degree(g, "c"))
	assert.Equal(t, 2, degree(g, "a"))
	assert.Equal(t, 1, degree(g, "d"))
}

func TestMaxDegree(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := newUndirected()
	g.connect("a", "b", 1)
	g.connect("a", "m", 1)

	g.connect("b", "c", 1)
	g.connect("m", "c", 1)

	g.connect("c", "d", 1)

	assert.Equal(t, 3, maxDegree(g))
}
