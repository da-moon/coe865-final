package djikstra

import (
	"container/heap"
	"log"
	"sync"

	"github.com/palantir/stacktrace"
)

// Graph ...
type Graph struct {
	lock    *sync.Mutex
	logger  *log.Logger
	nodes   map[string]Node
	counter int32
}

// NewGraph ...
func NewGraph(logger *log.Logger) Graph {
	result := Graph{
		logger:  logger,
		nodes:   make(map[string]Node),
		counter: 0,
		lock:    new(sync.Mutex),
	}
	return result
}

// Link ...
func (g Graph) Link(srcID string, neighbourID string, distance int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	src, ok := g.nodes[srcID]
	if !ok {
		err := stacktrace.NewError("source with ID '%s' does not exist", srcID)
		g.logger.Printf("[DEBUG] djikstra: err : %v", err)
		return err

	}
	neighbour, ok := g.nodes[neighbourID]
	if !ok {
		err := stacktrace.NewError("neighbour with ID '%s' does not exist", neighbourID)
		g.logger.Printf("[DEBUG] djikstra: err : %v", err)
		return err
	}
	src.Edges[neighbourID] = distance
	neighbour.Edges[srcID] = distance
	// a.edges = append(a.edges, edge{dstinatione: b, weight: distance})
	// b.edges = append(b.edges, edge{destination: a, weight: distance})
	return nil
}

// ShortestPath ...
func (g Graph) ShortestPath(startID string, endID string) (Path, error) {
	// g.lock.Lock()
	// defer g.lock.Unlock()

	start, ok := g.nodes[startID]
	if !ok {
		err := stacktrace.NewError("start node with ID '%s' does not exist", startID)
		g.logger.Printf("[DEBUG] djikstra: err : %v", err)
		return nil, err
	}
	end, ok := g.nodes[endID]
	if !ok {
		err := stacktrace.NewError("end node with ID '%s' does not exist", endID)
		g.logger.Printf("[DEBUG] djikstra: err : %v", err)
		return nil, err

	}
	g.logger.Printf("[INFO] djikstra: initializing shortest path")
	// empty struct to use least memory ... better than bool
	visited := make(map[string]struct{})
	// cumulative weight/cost
	cumulative := make(map[string]int)
	// keeping track of attachments
	// BUG maybe pointer ?
	prev := make(map[string]Node)
	g.logger.Printf("[INFO] djikstra: initializing starting node with a distance of zero")
	// start node has a distance of zero from itself
	cumulative[start.ID] = 0
	// force first node cost to zero
	start.weight = 0
	start.index = 0
	g.logger.Printf("[TRACE] djikstra: initializing queue")

	queue := &SortByNode{start}
	// inint heap
	g.logger.Printf("[TRACE] djikstra: initializing MinHeap tree")
	heap.Init(queue)
	// as long as there is room in the que
	for queue.Len() > 0 {
		g.logger.Printf("[DEBUG] djikstra: exploring with queue len on %d ...", queue.Len())
		// check break condition first
		_, ok := visited[end.ID]
		if ok {
			g.logger.Printf("[INFO] djikstra: reached destination ...")
			break
		}
		// the traversing prosess hasen
		item := heap.Pop(queue).(Node)
		g.logger.Printf("[INFO] djikstra: poped a node with index '%v' with weight '%v'out of queue ...", item.index, item.weight)

		g.logger.Printf("[INFO] djikstra: about to start iterating through edges of node '%s' to get cumulative weight...", item.ID)

		for destID, weight := range item.Edges {
			dest, ok := g.nodes[destID]
			if !ok {
				// not fail ... gently warn
				g.logger.Printf("[WARN] djikstra: Node '%s' tried to reach Node '%s' but it doens't exist. continuing ...", item.ID, destID)
				continue
			}
			dist := cumulative[item.ID] + weight
			g.logger.Printf("[DEBUG] djikstra: node '%s' dest '%s' edge weight '%d' cumulative '%d'", item.ID, dest.ID, weight, dist)
			tentativeDist, ok := cumulative[dest.ID]
			// in case it is not last node or total updated weight
			// at this instance is lower than previous path weight
			// update it and puuts it bach to queue
			if !ok || dist < tentativeDist {
				cumulative[dest.ID] = dist
				prev[dest.ID] = item
				dest.weight = dist
				heap.Push(queue, dest)
			}
		}
		// checking of the node as visited
		visited[item.ID] = struct{}{}
	}
	_, ok = visited[end.ID]
	if !ok {
		err := stacktrace.NewError("could not find shortest path from '%s' to '%s'", startID, endID)
		g.logger.Printf("[DEBUG] djikstra: err %v", err)
		return nil, err
	}
	g.logger.Printf("[DEBUG] djikstra: forming path arrau from node '%s' dest '%s' ", start.ID, end.ID)
	path := []Node{{
		ID: end.ID,
	},
	}

	for {
		next, ok := prev[end.ID]
		if !ok {
			break
		}
		path = append(path, next)
		next = prev[next.ID]
	}
	// for next := prev[end.ID]; next != nil; next = prev[next.ID] {
	// 	path = append(path, next)
	// }
	// Reverse path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return Path(path), nil
}
