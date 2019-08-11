package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigraph_dfs(t *testing.T) {
	// a -> b -> c -> d
	//  ^      /
	//   -  m
	g := &directed{}

	g.connect("a", "b")
	g.connect("b", "c")
	g.connect("c", "d")

	g.connect("c", "m")
	g.connect("m", "a")

	route := directedDFS(g, "a")
	assert.NotNil(t, route)

	assert.False(t, route.isReachable("foobar"))

	assert.Equal(t, path{"a"}, route.to("a"))

	// can take one of the path
	assert.Equal(t, path{"a", "b", "c", "d"}, route.to("d"))
	assert.Equal(t, path{"a", "b", "c", "m"}, route.to("m"))

}
