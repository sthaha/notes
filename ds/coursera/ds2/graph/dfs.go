package graph

import "log"

type dfsInfo struct {
	graph  graph
	target node
	path   map[node]node
}

func dfs(g graph, x node) *dfsInfo {
	p := &dfsInfo{
		graph:  g,
		target: x,
		path:   map[node]node{},
	}
	p.traverse(x, visited{})
	log.Printf("Path: %v", p.path)
	return p
}

func (p *dfsInfo) from(x node) path {
	log.Printf("Find path from: %v -> %v | path: %v", x, p.target, p.path)
	if _, ok := p.path[x]; !ok {
		log.Printf("node %v not in dfs: %v", x, p.path)
		return nil
	}

	path := path{x}
	for {
		log.Printf("   > %v |  path: %v", x, path)
		x = p.path[x]
		path = append(path, x)
		if x == p.target {
			return path
		}
	}

}

func (p *dfsInfo) to(x node) path {
	log.Printf("Find path to: %v from %v", x, p.target)
	path := p.from(x)
	if path != nil {
		path.reverse()
	}
	return path
}

// traverse  builds dfs information for node x
// internal
func (p *dfsInfo) traverse(x node, v visited) {
	log.Printf("path: %v", p.path)
	log.Printf("... mark %s visited", x)
	v[x] = true

	for _, adj := range p.graph.adjacents(x) {
		if v[adj] {
			log.Printf("   >> %v SKIPPED", adj)
			continue
		}
		log.Printf("   >> %v -> %v", adj, x)
		p.path[adj] = x
		p.traverse(adj, v)
	}
}
