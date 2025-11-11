package main

import (
	"fmt"
	"go_ai/structures"
	"os"
	"path/filepath"
)

func OnVisit(n structures.GraphNode[int]) {
	fmt.Printf("Visited node %v with value %v\n", n.Name, n.Val)
}

func main() {
	n1 := structures.NewGraphNode[int](1, "A")
	n2 := structures.NewGraphNode[int](2, "B")
	n3 := structures.NewGraphNode[int](3, "C")
	n4 := structures.NewGraphNode[int](4, "D")
	n5 := structures.NewGraphNode[int](5, "E")
	n6 := structures.NewGraphNode[int](6, "F")

	n1.AddAdj(n2)
	n1.AddAdj(n3)

	n2.AddAdj(n4)
	n2.AddAdj(n5)

	n3.AddAdj(n6)

	n1.BFS(OnVisit)

	svg, err := n1.ToDot()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(".", "graph.dot")
	err = os.WriteFile(path, []byte(svg), 0644)
	if err != nil {
		panic(err)
	}
}
