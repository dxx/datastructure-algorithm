package graph

import (
    "fmt"
    "testing"
)

func TestGraph(t *testing.T) {
    vertexes := []string{"A", "B", "C", "D", "E"}
    graph := NewGraph(5)

    for _, v := range vertexes {
        graph.AddVertex(v)
    }

    // A-B
    graph.AddEdge(0, 1, 1)
    // A-C
    graph.AddEdge(0, 2, 1)
    // B-C
    graph.AddEdge(1, 2, 1)
    // B-E
    graph.AddEdge(1, 4, 1)
    // C-D
    graph.AddEdge(2, 3, 1)

    graph.ShowEdges()

    fmt.Println("======深度优先遍历======")
    graph.DFS()

    fmt.Println()

    fmt.Println("======广度优先遍历======")
    graph.BFS()

    fmt.Println()
}
