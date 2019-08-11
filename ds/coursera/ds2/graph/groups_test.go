package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup_create(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := newUndirected()

	// a - b
	//  \
	//   - m
	g.connect("a", "b")
	g.connect("a", "m")

	//     b - c
	//       /
	//     m
	g.connect("b", "c")
	g.connect("m", "c")

	//         c - d
	//
	//
	g.connect("c", "d")

	grp := groups(g)
	assert.Equal(t, 1, grp.count())

}

func TestGroup_2_groups(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := newUndirected()

	// a - b
	//  \
	//   - m
	g.connect("a", "b")
	g.connect("a", "m")

	//     b - c
	//       /
	//     m
	g.connect("b", "c")
	g.connect("m", "c")

	//         c - d
	//
	//
	g.connect("c", "d")

	g.connect("x", "y")
	g.connect("x", "z")

	grp := groups(g)
	assert.Equal(t, 2, grp.count())

}

func TestGroup_same(t *testing.T) {
	g := newUndirected()
	// a - b - c - d
	g.connect("a", "b")
	g.connect("b", "c")
	g.connect("c", "d")

	// a - * - c
	//  \    /
	//   - f
	g.connect("a", "f")
	g.connect("f", "c")

	g.connect("x", "y")
	g.connect("x", "z")

	g.connect("m", "n")
	g.connect("n", "o")

	grp := groups(g)
	assert.Equal(t, 3, grp.count())

	assert.True(t, grp.same("a", "c"))
	assert.True(t, grp.same("a", "c", g.adjacents("d")...))

	assert.False(t, grp.same("a", "x"))
}
