package graph

import "log"

type lookup map[node]node

func (lookup lookup) from(x, target node) path {
	log.Printf("Find path from: %v -> %v | lookup: %v", x, target, lookup)
	if _, ok := lookup[x]; !ok {
		log.Printf("node %v not in lookup: %v", x, lookup)
		return nil
	}

	path := path{x}
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
