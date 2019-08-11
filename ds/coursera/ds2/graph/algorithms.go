package graph

func degree(g Graph, x Node) int {
	return len(g.Adjacents(x))
}

func maxDegree(g Graph) int {
	max := 0
	for _, n := range g.Nodes() {
		if d := degree(g, n); d > max {
			max = d
		}
	}
	return max
}
