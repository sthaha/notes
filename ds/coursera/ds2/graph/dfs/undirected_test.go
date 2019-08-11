package dfs

import (
	"testing"

	"github.com/sthaha/ds/graph"
	"github.com/stretchr/testify/assert"
)

func TestSearch_dfs(t *testing.T) {
	// a - b - c - d
	//  \    /
	//   - m
	g := &graph.Undirected{}

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

	route := ForUndirected(g, "a")
	assert.NotNil(t, route)

	assert.Nil(t, route.To("foobar"))

	assert.Equal(t, graph.Path{"a"}, route.From("a"))

	// can take one of the path
	assert.Contains(t, []graph.Path{
		graph.Path{"a", "b", "c", "d"},
		graph.Path{"a", "m", "c", "d"},
	}, route.To("d"))

	assert.Contains(t, []graph.Path{
		graph.Path{"m", "c", "b", "a"},
		graph.Path{"m", "a"},
	}, route.From("m"))
}
