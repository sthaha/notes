package graph

import "log"

type lookup map[Node]Node

func (lookup lookup) from(x, target Node) Path {
	log.Printf("Find path from: %v -> %v | lookup: %v", x, target, lookup)
	if _, ok := lookup[x]; !ok {
		log.Printf("node %v not in lookup: %v", x, lookup)
		return nil
	}

	path := Path{x}
	if x == target {
		return path
	}

	for {
		log.Printf("   > %v |  path: %v", x, path)
		x = lookup[x]
		path = append(path, x)
		if x == target {
			return path
		}
	}

}
