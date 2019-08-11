package graph

type dummy struct{}
type nodeset struct {
	data map[node]dummy
}

func (ns *nodeset) add(n node) {
	ns.init()
	ns.data[n] = dummy{}
}

func (ns *nodeset) has(n node) bool {
	if ns.data == nil {
		return false
	}
	_, ok := ns.data[n]
	return ok
}

func (ns *nodeset) keys() []node {
	ns.init()

	keys := make([]node, 0, len(ns.data))
	for k := range ns.data {
		keys = append(keys, k)
	}
	return keys
}

func (ns *nodeset) init() {
	if ns.data != nil {
		return
	}
	ns.data = map[node]dummy{}
}
