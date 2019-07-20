package graph

import "log"

func degree(g graph, x node) int {
	return len(g.adjacents(x))
}

func maxDegree(g graph) int {
	max := 0
	for _, n := range g.nodes() {
		if d := degree(g, n); d > max {
			max = d
		}
	}
	return max
}

type visited map[node]bool
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

func (p *dfsInfo) pathFrom(x node) path {
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

func (p *dfsInfo) pathTo(x node) path {
	log.Printf("Find path to: %v from %v", x, p.target)
	path := p.pathFrom(x)
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
			log.Printf("   >> %s SKIPPED", adj)
			continue
		}
		log.Printf("   >> %s", adj)
		p.path[adj] = x
		p.traverse(adj, v)
	}
}
