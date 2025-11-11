package structures

import (
	"bytes"
	"context"
	"fmt"
	"slices"

	"github.com/goccy/go-graphviz"
)

type GraphNode[T any] struct {
	Val  T
	Name string
	Adj  []*GraphNode[T]
}

func NewGraphNode[T any](val T, name string) *GraphNode[T] {
	adj := make([]*GraphNode[T], 0)
	return &GraphNode[T]{val, name, adj}
}

func (n *GraphNode[T]) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (n *GraphNode[T]) AddAdj(adj *GraphNode[T]) {
	n.Adj = append(n.Adj, adj)
}

func (n *GraphNode[T]) RemoveAdj(adj *GraphNode[T]) {
	for i, a := range n.Adj {
		if a == adj {
			n.Adj = append(n.Adj[:i], n.Adj[i+1:]...)
			break
		}
	}
}

func (n *GraphNode[T]) BFS(onVisit func(GraphNode[T])) {
	visited := []*GraphNode[T]{}
	queue := NewQueue[*GraphNode[T]]()

	queue.Enqueue(n)

	for !queue.IsEmpty() {
		cur, _ := queue.Dequeue()
		if slices.Contains(visited, cur) {
			continue
		}

		onVisit(*cur)
		visited = append(visited, cur)

		for _, adj := range cur.Adj {
			queue.Enqueue(adj)
		}

	}
}

func (n *GraphNode[T]) DFS(onVisit func(GraphNode[T])) {
	visited := []*GraphNode[T]{}
	stack := NewStack[*GraphNode[T]]()

	stack.Push(n)

	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		if slices.Contains(visited, cur) {
			continue
		}

		onVisit(*cur)
		visited = append(visited, cur)

		// NOTE(daniel): push in reverse order, so we visit them in the order the nodes got appended to the next array
		nextLength := len(cur.Adj)
		for i := nextLength - 1; i >= 0; i-- {
			stack.Push(cur.Adj[i])
		}

	}
}

func (n *GraphNode[T]) ToDot() (out string, err error) {
	ctx := context.Background()
	g, err := graphviz.New(ctx)
	if err != nil {
		return
	}

	graph, err := g.Graph()
	if err != nil {
		return
	}

	OnVisit := func(n GraphNode[T]) {
		start, _ := graph.CreateNodeByName(n.Name)
		for _, adj := range n.Adj {
			end, _ := graph.CreateNodeByName(adj.Name)
			graph.CreateEdgeByName("", start, end)
		}
	}

	n.BFS(OnVisit)
	var buf bytes.Buffer
	g.Render(ctx, graph, "dot", &buf)
	out = buf.String()
	return
}
