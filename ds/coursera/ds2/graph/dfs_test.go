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

	route := dfs(g, "a")
	assert.NotNil(t, route)

	assert.Nil(t, route.to("foobar"))

	assert.Equal(t, path{"a"}, route.from("a"))

	// can take one of the path
	assert.Contains(t, []path{
		path{"a", "b", "c", "d"},
		path{"a", "m", "c", "d"},
	}, route.to("d"))

	assert.Contains(t, []path{
		path{"m", "c", "b", "a"},
		path{"m", "a"},
	}, route.from("m"))
}
