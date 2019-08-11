package graph

import "log"

type directedDfsInfo struct {
	graph     graph
	origin    node
	wayPoints map[node]node
}

func directedDFS(g *directed, x node) *directedDfsInfo {
	d := &directedDfsInfo{
		graph:     g,
		origin:    x,
		wayPoints: map[node]node{x: x},
	}
	d.traverse(x, visited{})
	log.Printf("way-points: %v", d.wayPoints)
	return d
}

func (d *directedDfsInfo) to(x node) path {
	log.Printf("Find path from: %v -> %v | path: %v", x, d.origin, d.wayPoints)
	route := path{x}
	if x == d.origin {
		return route
	}

	if !d.isReachable(x) {
		log.Printf("node %v not in dfs: %v", x, d.wayPoints)
		return nil
	}

	for {
		log.Printf("   > %v |  path: %v", x, route)
		x = d.wayPoints[x]
		route = append(path{x}, route...)
		if x == d.origin {
			break
		}
	}
	return route
}

func (d *directedDfsInfo) isReachable(x node) bool {
	_, ok := d.wayPoints[x]
	return ok
}

// traverse  builds dfs information for node x
// internal
func (d *directedDfsInfo) traverse(x node, v visited) {
	log.Printf("path: %v", d.wayPoints)
	log.Printf("... mark %s visited", x)
	v[x] = true

	for _, adj := range d.graph.adjacents(x) {
		if v[adj] {
			log.Printf("   >> %v SKIPPED", adj)
			continue
		}
		log.Printf("   >> %v -> %v", adj, x)
		d.wayPoints[adj] = x
		d.traverse(adj, v)
	}
}
