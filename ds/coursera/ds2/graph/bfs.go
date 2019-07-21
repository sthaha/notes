package graph

import "log"

type bfsInfo struct {
	graph  graph
	target node
	path   map[node]node
}

func bfs(g graph, x node) *bfsInfo {
	p := &bfsInfo{
		graph:  g,
		target: x,
		path:   map[node]node{x: x},
	}
	p.traverse(x, visited{})
	log.Printf("Path: %v", p.path)
	return p
}

// TODO: from and to is the same between dfs and bfs .. make interface
func (p *bfsInfo) from(x node) path {
	log.Printf("Find path from: %v -> %v | path: %v", x, p.target, p.path)
	path := path{x}
	if x == p.target {
		return path
	}

	if _, ok := p.path[x]; !ok {
		log.Printf("node %v not in bfs: %v", x, p.path)
		return nil
	}

	for {
		log.Printf("   > %v |  path: %v", x, path)
		x = p.path[x]
		path = append(path, x)
		if x == p.target {
			return path
		}
	}

}

func (p *bfsInfo) to(x node) path {
	log.Printf("Find path to: %v from %v", x, p.target)
	path := p.from(x)
	if path != nil {
		path.reverse()
	}
	return path
}

// traverse  builds bfs information for node x
// internal
func (p *bfsInfo) traverse(x node, v visited) {
	log.Printf("path: %v", p.path)
	//log.Printf("... mark %s visited", x)
	//v[x] = true

	next := []node{}
	for _, adj := range p.graph.adjacents(x) {
		//if v[adj] {
		//log.Printf("   >> %v SKIPPED", adj)
		//continue
		//}
		if _, ok := p.path[adj]; ok {
			log.Printf("   >> %v SKIPPED: already in path", adj)
			continue
		}
		log.Printf("   >> %v -> %v", adj, x)
		p.path[adj] = x
		//v[adj] = true
		next = append(next, adj)
	}

	for _, n := range next {
		p.traverse(n, v)
	}
}
