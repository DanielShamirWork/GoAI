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
	n1 := structures.NewGraphNode(1, "A")
	n2 := structures.NewGraphNode(2, "B")
	n3 := structures.NewGraphNode(3, "C")
	n4 := structures.NewGraphNode(4, "D")
	n5 := structures.NewGraphNode(5, "E")
	n6 := structures.NewGraphNode(6, "F")

	n1.AddAdj(n2)
	n1.AddAdj(n3)

	n2.AddAdj(n4)
	n2.AddAdj(n5)

	n3.AddAdj(n6)

	n1.BFS(OnVisit)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)

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
