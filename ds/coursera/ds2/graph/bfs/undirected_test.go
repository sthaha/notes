package bfs

import (
	"testing"

	"github.com/sthaha/ds/graph"
	"github.com/stretchr/testify/assert"
)

func TestSearch_bfs_route(t *testing.T) {
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
	assert.Nil(t, route.to("foobar"))

	assert.Equal(t, graph.Path{"a"}, route.from("a"))

	fromM := route.from("m")
	assert.Equal(t, graph.Path{"m", "a"}, fromM)

	toD := route.to("d") // can be through b or m
	assert.Contains(t, []graph.Path{
		graph.Path{"a", "b", "c", "d"},
		graph.Path{"a", "m", "c", "d"},
	}, toD)

}

func TestSearch_bfs_depth(t *testing.T) {
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
	_, err := route.depthTo("foobar")
	assert.Error(t, graph.ErrNodeNotFound, err)

	d, err := route.depthTo("d")
	assert.NoError(t, err)
	assert.Equal(t, 3, d)

}
