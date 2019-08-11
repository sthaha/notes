package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch_bfs_route(t *testing.T) {
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

	route := bfs(g, "a")
	assert.NotNil(t, route)
	assert.Nil(t, route.to("foobar"))

	assert.Equal(t, path{"a"}, route.from("a"))

	fromM := route.from("m")
	assert.Equal(t, path{"m", "a"}, fromM)

	toD := route.to("d") // can be through b or m
	assert.Contains(t, []path{
		path{"a", "b", "c", "d"},
		path{"a", "m", "c", "d"},
	}, toD)

}

func TestSearch_bfs_depth(t *testing.T) {
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

	route := bfs(g, "a")
	_, err := route.depthTo("foobar")
	assert.Error(t, errNodeNotFound, err)

	d, err := route.depthTo("d")
	assert.NoError(t, err)
	assert.Equal(t, 3, d)

}
