package graph

import "log"

type dfsInfo struct {
	graph     graph
	origin    node
	wayPoints map[node]node
}

func undirectedDFS(g *undirected, x node) *dfsInfo {
	d := &dfsInfo{
		graph:     g,
		origin:    x,
		wayPoints: map[node]node{x: x},
	}
	d.traverse(x, visited{})
	log.Printf("way-points: %v", d.wayPoints)
	return d
}

func (d *dfsInfo) from(x node) path {
	log.Printf("Find path from: %v -> %v | path: %v", x, d.origin, d.wayPoints)
	path := path{x}
	if x == d.origin {
		return path
	}

	if _, ok := d.wayPoints[x]; !ok {
		log.Printf("node %v not in dfs: %v", x, d.wayPoints)
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

func (d *dfsInfo) to(x node) path {
	log.Printf("Find path to: %v from %v", x, d.origin)
	path := d.from(x)
	if path != nil {
		path.reverse()
	}
	return path
}

// traverse  builds dfs information for node x
// internal
func (d *dfsInfo) traverse(x node, v visited) {
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
