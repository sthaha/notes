package bfs

import (
	"testing"

	"github.com/sthaha/ds/graph"
	"github.com/stretchr/testify/assert"
)

func TestDigraph_bfs(t *testing.T) {
	//         -----------------------> x
	//        /                       ^
	//       /                       /
	// a -> b -> c -> d ------------'
	//  ^      /
	//   -  m
	g := &graph.Directed{}

	g.Connect("a", "b")
	g.Connect("b", "c")
	g.Connect("b", "x")

	g.Connect("c", "d")
	g.Connect("c", "m")
	g.Connect("m", "a")

	route := ForDirected(g, "a")
	assert.NotNil(t, route)

	assert.False(t, route.IsReachable("foobar"))

	assert.Equal(t, graph.Path{"a"}, route.To("a"))

	assert.Equal(t, graph.Path{"a", "b", "c", "d"}, route.To("d"))
	assert.Equal(t, graph.Path{"a", "b", "c", "m"}, route.To("m"))

	// takes shortest path
	assert.Equal(t, graph.Path{"a", "b", "x"}, route.To("x"))

}
