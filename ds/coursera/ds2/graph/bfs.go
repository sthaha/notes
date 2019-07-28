package graph

import (
	"log"

	"github.com/sthaha/ds/queue"
)

type bfsInfo struct {
	graph  graph
	target node
	path   map[node]node
	depth  map[node]int
}

func bfs(g graph, x node) *bfsInfo {
	p := &bfsInfo{
		graph:  g,
		target: x,
		path:   map[node]node{x: x},
		depth:  map[node]int{},
	}

	p.traverse()
	log.Printf("Path: %v", p.path)
	return p
}

func (p *bfsInfo) depthTo(x node) (int, error) {
	d, ok := p.depth[x]
	if !ok {
		log.Printf("node %v not in bfs: %v", x, p.path)
		return -1, errNodeNotFound
	}
	return d, nil
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

func (p *bfsInfo) traverse() {

	q := &queue.Queue{}
	q.Push(p.target)
	p.depth[p.target] = 0

	for !q.IsEmpty() {
		x, _ := q.Pop()
		depth := p.depth[x] + 1

		log.Printf("   >> %v  traversing at %d", x, depth)

		for _, adj := range p.graph.adjacents(x) {
			if _, ok := p.path[adj]; ok {
				log.Printf("     >> %v SKIPPED: already in path", adj)
				continue
			}
			p.path[adj] = x
			p.depth[adj] = depth
			log.Printf("     >> %v -> %v: %d | %v", adj, x, depth, p.depth)
			q.Push(adj)
		}

	}
}

// traverse  builds bfs information for node x
// internal
func (p *bfsInfo) traverseRecursive(x node) {
	log.Printf("path: %v", p.path)
	log.Printf("... mark %s visited", x)

	next := []node{}
	for _, adj := range p.graph.adjacents(x) {
		if _, ok := p.path[adj]; ok {
			log.Printf("   >> %v SKIPPED: already in path", adj)
			continue
		}
		log.Printf("   >> %v -> %v", adj, x)
		p.path[adj] = x
		next = append(next, adj)
	}

	for _, n := range next {
		p.traverseRecursive(n)
	}
}
