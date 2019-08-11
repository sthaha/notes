package dfs

import (
	"log"

	"github.com/sthaha/ds/graph"
)

type Directed struct {
	graph     graph.Graph
	origin    graph.Node
	wayPoints map[graph.Node]graph.Node
}

func ForDirected(g *graph.Directed, x graph.Node) *Directed {
	d := &Directed{
		graph:     g,
		origin:    x,
		wayPoints: map[graph.Node]graph.Node{x: x},
	}
	d.traverse(x, visited{})
	log.Printf("way-points: %v", d.wayPoints)
	return d
}

func (d *Directed) To(x graph.Node) graph.Path {
	log.Printf("Find path from: %v -> %v | path: %v", x, d.origin, d.wayPoints)
	route := graph.Path{x}
	if x == d.origin {
		return route
	}

	if !d.IsReachable(x) {
		log.Printf("graph.Node %v not in dfs: %v", x, d.wayPoints)
		return nil
	}

	for {
		log.Printf("   > %v |  path: %v", x, route)
		x = d.wayPoints[x]
		route = append(graph.Path{x}, route...)
		if x == d.origin {
			break
		}
	}
	return route
}

func (d *Directed) IsReachable(x graph.Node) bool {
	_, ok := d.wayPoints[x]
	return ok
}

// traverse  builds dfs information for graph.Node x
// internal
func (d *Directed) traverse(x graph.Node, v visited) {
	log.Printf("path: %v", d.wayPoints)
	log.Printf("... mark %s visited", x)
	v[x] = true

	for _, adj := range d.graph.Adjacents(x) {
		if v[adj] {
			log.Printf("   >> %v SKIPPED", adj)
			continue
		}
		log.Printf("   >> %v -> %v", adj, x)
		d.wayPoints[adj] = x
		d.traverse(adj, v)
	}
}
