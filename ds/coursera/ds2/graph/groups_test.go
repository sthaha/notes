package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup_create(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := &Undirected{}

	// a - b
	//  \
	//   - m
	g.Connect("a", "b")
	g.Connect("a", "m")

	//     b - c
	//       /
	//     m
	g.Connect("b", "c")
	g.Connect("m", "c")

	//         c - d
	//
	//
	g.Connect("c", "d")

	grp := groups(g)
	assert.Equal(t, 1, grp.count())

}

func TestGroup_2_groups(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := &Undirected{}

	// a - b
	//  \
	//   - m
	g.Connect("a", "b")
	g.Connect("a", "m")

	//     b - c
	//       /
	//     m
	g.Connect("b", "c")
	g.Connect("m", "c")

	//         c - d
	//
	//
	g.Connect("c", "d")

	g.Connect("x", "y")
	g.Connect("x", "z")

	grp := groups(g)
	assert.Equal(t, 2, grp.count())

}

func TestGroup_same(t *testing.T) {
	g := &Undirected{}
	// a - b - c - d
	g.Connect("a", "b")
	g.Connect("b", "c")
	g.Connect("c", "d")

	// a - * - c
	//  \    /
	//   - f
	g.Connect("a", "f")
	g.Connect("f", "c")

	g.Connect("x", "y")
	g.Connect("x", "z")

	g.Connect("m", "n")
	g.Connect("n", "o")

	grp := groups(g)
	assert.Equal(t, 3, grp.count())

	assert.True(t, grp.same("a", "c"))
	assert.True(t, grp.same("a", "c", g.Adjacents("d")...))

	assert.False(t, grp.same("a", "x"))
}
