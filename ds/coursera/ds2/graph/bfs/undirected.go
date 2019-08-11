package bfs

import (
	"log"

	"github.com/sthaha/ds/graph"
	"github.com/sthaha/ds/queue"
)

type bfsInfo struct {
	graph  graph.Graph
	target graph.Node
	path   map[graph.Node]graph.Node
	depth  map[graph.Node]int
}

func ForUndirected(g *graph.Undirected, x graph.Node) *bfsInfo {
	p := &bfsInfo{
		graph:  g,
		target: x,
		path:   map[graph.Node]graph.Node{x: x},
		depth:  map[graph.Node]int{},
	}

	p.traverse()
	log.Printf("Path: %v", p.path)
	return p
}

func (p *bfsInfo) depthTo(x graph.Node) (int, error) {
	d, ok := p.depth[x]
	if !ok {
		log.Printf("graph.Node %v not in bfs: %v", x, p.path)
		return -1, graph.ErrNodeNotFound
	}
	return d, nil
}

// TODO: from and to is the same between dfs and bfs .. make interface
func (p *bfsInfo) from(x graph.Node) graph.Path {
	log.Printf("Find path from: %v -> %v | path: %v", x, p.target, p.path)
	path := graph.Path{x}
	if x == p.target {
		return path
	}

	if _, ok := p.path[x]; !ok {
		log.Printf("graph.Node %v not in bfs: %v", x, p.path)
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

func (p *bfsInfo) to(x graph.Node) graph.Path {
	log.Printf("Find path to: %v from %v", x, p.target)
	path := p.from(x)
	if path != nil {
		path.Reverse()
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

		for _, adj := range p.graph.Adjacents(x) {
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

// traverse  builds bfs information for graph.Node x
// internal
func (p *bfsInfo) traverseRecursive(x graph.Node) {
	log.Printf("path: %v", p.path)
	log.Printf("... mark %s visited", x)

	next := []graph.Node{}
	for _, adj := range p.graph.Adjacents(x) {
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
