package models

import (
	"sync"
)

type Node struct {
	URL string `json:"url"`
}

type NodesRepository struct {
	Nodes map[string]*Node
	mutex sync.RWMutex
}

func (n *NodesRepository) Get(key string) *Node {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	node := n.Nodes[key]

	return node
}
func (n *NodesRepository) Put(node Node) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.Nodes[node.URL] = &node
}

func (nr *NodesRepository) GetAllNodes() []Node {
	nr.mutex.RLock()
	defer nr.mutex.RUnlock()

	nodes := make([]Node, 0, len(nr.Nodes))
	for _, node := range nr.Nodes {
		nodes = append(nodes, *node)
	}

	return nodes
}
