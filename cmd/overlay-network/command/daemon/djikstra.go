package daemon

import (
	"container/heap"
	"errors"
)

// Node ...
type Node interface {
	Edges() []Edge
}

// Edge ...
type Edge interface {
	Destination() Node
	Weight() int
}

// ShortestPath ...
func ShortestPath(start, end Node) ([]Node, error) {
	visited := make(map[Node]struct{})
	dists := make(map[Node]int)
	prev := make(map[Node]Node)

	dists[start] = 0
	queue := &queue{&queueItem{value: start, weight: 0, index: 0}}
	heap.Init(queue)

	for queue.Len() > 0 {
		// Done.
		if _, ok := visited[end]; ok {
			break
		}

		item := heap.Pop(queue).(*queueItem)
		n := item.value
		for _, edge := range n.Edges() {
			dest := edge.Destination()
			dist := dists[n] + edge.Weight()
			if tentativeDist, ok := dists[dest]; !ok || dist < tentativeDist {
				dists[dest] = dist
				prev[dest] = n
				heap.Push(queue, &queueItem{value: dest, weight: dist})
			}
		}
		visited[n] = struct{}{}
	}

	if _, ok := visited[end]; !ok {
		return nil, errors.New("no shortest path exists")
	}

	path := []Node{end}
	for next := prev[end]; next != nil; next = prev[next] {
		path = append(path, next)
	}

	// Reverse path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, nil
}

type queueItem struct {
	value  Node
	weight int
	index  int
}

type queue []*queueItem

// Len ...
func (q queue) Len() int {
	return len(q)
}

// Less ...
func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

// Swap ...
func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Push ...
func (q *queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*queueItem)
	item.index = n
	*q = append(*q, item)
}

// Pop ...
func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

// Node implementation
type vertex struct {
	id    string
	edges []edge
}

// Edges ...
func (v *vertex) Edges() []Edge {
	edges := make([]Edge, len(v.edges))
	for i := range v.edges {
		edges[i] = v.edges[i]
	}
	return edges
}

// Edge implementation
type edge struct {
	destination *vertex
	weight      int
}

// Destination ...
func (e edge) Destination() Node {
	return e.destination
}

// Weight ...
func (e edge) Weight() int {
	return e.weight
}

// Connect two vertices both ways
func link(a, b *vertex, dist int) {
	a.edges = append(a.edges, edge{destination: b, weight: dist})
	b.edges = append(b.edges, edge{destination: a, weight: dist})
}
