package graph

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
