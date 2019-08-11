package graph

// path is a list of nodes with edge between each of them representing
// a route from the first element to the last
type Path []Node

func (p Path) Reverse() {
	for left, right := 0, len(p)-1; left < right; left, right = left+1, right-1 {
		p[left], p[right] = p[right], p[left]
	}
}
