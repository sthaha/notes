package dfs

import (
	"log"

	"github.com/sthaha/ds/graph"
)

type visited map[graph.Node]bool

type Undirected struct {
	graph     graph.Graph
	origin    graph.Node
	wayPoints map[graph.Node]graph.Node
}

func ForUndirected(g *graph.Undirected, x graph.Node) *Undirected {
	d := &Undirected{
		graph:     g,
		origin:    x,
		wayPoints: map[graph.Node]graph.Node{x: x},
	}
	d.traverse(x, visited{})
	log.Printf("way-points: %v", d.wayPoints)
	return d
}

func (d *Undirected) From(x graph.Node) graph.Path {
	log.Printf("Find path from: %v -> %v | path: %v", x, d.origin, d.wayPoints)
	path := graph.Path{x}
	if x == d.origin {
		return path
	}

	if _, ok := d.wayPoints[x]; !ok {
		log.Printf("graph.Node %v not in dfs: %v", x, d.wayPoints)
		return nil
	}

	for {
		log.Printf("   > %v |  path: %v", x, path)
		x = d.wayPoints[x]
		path = append(path, x)
		if x == d.origin {
			return path
		}
	}

}

func (d *Undirected) To(x graph.Node) graph.Path {
	log.Printf("Find path to: %v from %v", x, d.origin)
	path := d.From(x)
	if path != nil {
		path.Reverse()
	}
	return path
}

// traverse  builds dfs information for graph.Node x
// internal
func (d *Undirected) traverse(x graph.Node, v visited) {
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
