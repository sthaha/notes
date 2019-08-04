package graph

import (
	"log"

	"github.com/sthaha/ds/queue"
)

type groupInfo struct {
	graph  graph
	info   map[node]int
	groups int
}

func groups(g graph) *groupInfo {
	grp := &groupInfo{
		graph: g,
		info:  map[node]int{},
	}
	grp.traverse()
	return grp
}

func (g *groupInfo) count() int {
	return g.groups
}

func (g *groupInfo) same(first, second node, rest ...node) bool {
	group := g.info[first]
	if g.info[second] != group {
		return false
	}

	for _, n := range rest {
		if g.info[n] != group {
			return false
		}
	}
	return true
}

func (g *groupInfo) traverse() {
	for _, n := range g.graph.nodes() {
		if _, ok := g.info[n]; ok {
			log.Printf("Already traversed node: %v", n)
			continue
		}
		g.bfs(n)
		g.groups++
	}
}

func (g *groupInfo) bfs(n node) {
	q := queue.Queue{}
	g.info[n] = g.groups
	q.Push(n)

	for !q.IsEmpty() {
		x, _ := q.Pop()
		log.Printf("   >> %v  traversing group %d", x, g.groups)

		for _, adj := range g.graph.adjacents(x) {
			if _, ok := g.info[adj]; !ok {
				g.info[adj] = g.groups
				q.Push(adj)
			}
		}
	}
}
