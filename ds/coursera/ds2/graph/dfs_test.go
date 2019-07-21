package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch_dfs(t *testing.T) {
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

	route := dfs(g, "a")
	assert.NotNil(t, route)

	assert.Nil(t, route.to("foobar"))

	toD := route.to("d")
	assert.NotNil(t, toD)

	assert.Equal(t, path{"a", "b", "c", "d"}, toD)

	fromM := route.from("m")
	assert.Equal(t, path{"m", "c", "b", "a"}, fromM)
}
