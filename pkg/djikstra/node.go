package djikstra

import (
	"strings"
	"sync/atomic"
)

// Node ...
type Node struct {
	ID     string
	weight int
	index  int32
	// k -> dest ID
	// v -> weigth
	Edges map[string]int
}

// Path ...
type Path []Node

// String ...
func (p Path) String() string {
	result := make([]string, 0)
	for _, v := range p {
		result = append(result, v.ID)
	}
	return strings.Join(result, ",")
}

// NewNode ...
// ID string,
func (g Graph) NewNode(ID string) {
	result := Node{Edges: make(map[string]int)}
	result.ID = ID
	result.weight = 0
	result.index = atomic.AddInt32(&g.counter, 1)
	g.nodes[ID] = result
}

// SortByNode for to use heap , besides pop and push
// we need to implement sortable interface
type SortByNode []Node

// Push ...
func (s SortByNode) Push(x interface{}) {
	n := len(s)
	item := x.(Node)
	item.index = int32(n)
	s = append(s, item)
}

// Pop ...
func (s SortByNode) Pop() interface{} {
	old := s
	n := len(old)
	item := old[n-1]
	item.index = -1
	s = old[0 : n-1]
	return item
}

// Len ...
func (s SortByNode) Len() int {
	return len(s)
}

// Less ...
func (s SortByNode) Less(i, j int) bool {
	return s[i].weight < s[j].weight
}

// Swap ...
func (s SortByNode) Swap(i, j int) {
	s[i] = s[j]
	s[j] = s[i]
	s[i].index = int32(i)
	s[j].index = int32(j)
}
