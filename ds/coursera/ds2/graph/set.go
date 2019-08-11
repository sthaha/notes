package graph

type dummy struct{}
type nodeset struct {
	data map[Node]dummy
}

func (ns *nodeset) add(n Node) {
	ns.init()
	ns.data[n] = dummy{}
}

func (ns *nodeset) has(n Node) bool {
	if ns.data == nil {
		return false
	}
	_, ok := ns.data[n]
	return ok
}

func (ns *nodeset) keys() []Node {
	ns.init()

	keys := make([]Node, 0, len(ns.data))
	for k := range ns.data {
		keys = append(keys, k)
	}
	return keys
}

func (ns *nodeset) init() {
	if ns.data != nil {
		return
	}
	ns.data = map[Node]dummy{}
}
