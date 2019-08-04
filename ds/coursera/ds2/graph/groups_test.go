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
	g.connect("a", "b", 1)
	g.connect("a", "m", 1)

	//     b - c
	//       /
	//     m
	g.connect("b", "c", 1)
	g.connect("m", "c", 1)

	//         c - d
	//
	//
	g.connect("c", "d", 1)

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
	g.connect("a", "b", 1)
	g.connect("a", "m", 1)

	//     b - c
	//       /
	//     m
	g.connect("b", "c", 1)
	g.connect("m", "c", 1)

	//         c - d
	//
	//
	g.connect("c", "d", 1)

	g.connect("x", "y", 1)
	g.connect("x", "z", 1)

	grp := groups(g)
	assert.Equal(t, 2, grp.count())

}

func TestGroup_same(t *testing.T) {
	g := newUndirected()
	// a - b - c - d
	g.connect("a", "b", 1)
	g.connect("b", "c", 1)
	g.connect("c", "d", 1)

	// a - * - c
	//  \    /
	//   - f
	g.connect("a", "f", 1)
	g.connect("f", "c", 1)

	g.connect("x", "y", 1)
	g.connect("x", "z", 1)

	g.connect("m", "n", 1)
	g.connect("n", "o", 1)

	grp := groups(g)
	assert.Equal(t, 3, grp.count())

	assert.True(t, grp.same("a", "c"))
	assert.True(t, grp.same("a", "c", g.adjacents("d")...))

	assert.False(t, grp.same("a", "x"))
}
