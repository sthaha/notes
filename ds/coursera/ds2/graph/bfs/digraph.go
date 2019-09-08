package bfs

import (
	"log"

	"github.com/sthaha/ds/graph"
	"github.com/sthaha/ds/queue"
)

type Directed struct {
	graph     graph.Graph
	origin    graph.Node
	wayPoints map[graph.Node]graph.Node
	depth     map[graph.Node]int
}

func ForDirected(g graph.Graph, x graph.Node) *Directed {
	d := &Directed{
		graph:     g,
		origin:    x,
		wayPoints: map[graph.Node]graph.Node{},
		depth:     map[graph.Node]int{},
	}
	d.traverse()
	return d
}

func (d *Directed) To(target graph.Node) graph.Path {
	log.Printf("Find path from: %v -> %v | path: %v", d.origin, target, d.wayPoints)

	if !d.IsReachable(target) {
		log.Printf("graph.Node %v not in bfs: %v", target, d.wayPoints)
		return nil
	}

	route := graph.Path{}
	for x := target; x != d.origin; {
		log.Printf("   > %v |  path: %v", x, route)
		route = append(graph.Path{x}, route...)
		x = d.wayPoints[x]
	}
	route = append(graph.Path{d.origin}, route...)
	return route
}

func (d *Directed) IsReachable(target graph.Node) bool {
	_, ok := d.wayPoints[target]
	return ok
}

func (d *Directed) traverse() {

	q := &queue.Queue{}
	q.Push(d.origin)
	d.depth[d.origin] = 0
	d.wayPoints[d.origin] = d.origin

	visited := map[graph.Node]bool{}
	for !q.IsEmpty() {
		parent, _ := q.Pop() // ignore error since !IsEmpty() ensures it
		depth := d.depth[parent] + 1
		visited[parent] = true

		log.Printf("   >> %v  traversing at %d", parent, depth)

		for _, adj := range d.graph.Adjacents(parent) {
			if visited[adj] {
				continue
			}

			d.wayPoints[adj] = parent
			d.depth[adj] = depth

			q.Push(adj)
		}
	}
}
