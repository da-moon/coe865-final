package djikstra_test

import (
	"container/heap"
	"log"
	"testing"

	"github.com/da-moon/coe865-final/pkg/djikstra"
	"github.com/stretchr/testify/assert"
)

func init() {
	var _ heap.Interface = djikstra.SortByNode{}
}

type Connection struct {
	src    string
	dst    string
	weight int
}
type TableTest struct {
	nodes    []string
	links    []Connection
	expected djikstra.Path
}

var test = TableTest{
	nodes: []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
	},
	links: []Connection{
		{src: "A", dst: "B", weight: 1},
		{src: "A", dst: "C", weight: 3},
		{src: "A", dst: "D", weight: 2},
		{src: "B", dst: "E", weight: 1},
		{src: "B", dst: "F", weight: 2},
		{src: "C", dst: "D", weight: 1},
		{src: "C", dst: "F", weight: 6},
		{src: "D", dst: "F", weight: 4},
		{src: "E", dst: "F", weight: 1},
	},
}

func TestDjikstra(t *testing.T) {
	graph := djikstra.NewGraph(log.New(testWriter{t}, "test", log.LstdFlags))
	for _, node := range test.nodes {
		graph.NewNode(node)
	}
	// setting up connections
	for _, v := range test.links {
		err := graph.Link(v.src, v.dst, v.weight)
		ok := assert.NoError(t, err)
		if !ok {
			t.Fail()
		}
	}
	path, err := graph.ShortestPath("A", "F")
	ok := assert.NoError(t, err)
	if !ok {
		t.Fail()
	}
	t.Log(path.String())
}

type testWriter struct {
	t *testing.T
}

func (tw testWriter) Write(p []byte) (n int, err error) {
	tw.t.Log(string(p))
	return len(p), nil
}
