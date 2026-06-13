package prim

import "testing"

func TestMinTree(t *testing.T) {
    vertexes := []string{"A", "B", "C", "D", "E", "F", "G"}
    edges := [][]int{
        {intMax, 5, 7, intMax, intMax, intMax, 2},
        {5, intMax, intMax, 9, intMax, intMax, 3},
        {7, intMax, intMax, intMax, 8, intMax, intMax},
        {intMax, 9, intMax, intMax, intMax, 4, intMax},
        {intMax, intMax, 8, intMax, intMax, 5, 4},
        {intMax, intMax, intMax, 4, 5, intMax, 6},
        {2, 3, intMax, intMax, 4, 6, intMax},
    }
    minTree := NewMinTree(vertexes, edges)
    // minTree.ShowGraph()
    // 从 A 点开始
    minTree.Prim(0)
}
