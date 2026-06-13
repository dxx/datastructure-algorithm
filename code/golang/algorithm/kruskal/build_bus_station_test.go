package kruskal

import "testing"

func TestMinTree(t *testing.T) {
    vertexes := []string{"A", "B", "C", "D", "E", "F", "G"}
    // 0-表示自己跟自己不连通，intMax-表示跟其它顶点不连通
    edges := [][]int{
        {0, 12, intMax, intMax, intMax, 16, 14},
        {12, 0, 10, intMax, intMax, 7, intMax},
        {intMax, 10, 0, 3, 5, 6, intMax},
        {intMax, intMax, 3, 0, 4, intMax, intMax},
        {intMax, intMax, 5, 4, 0, 2, 8},
        {16, 7, 6, intMax, 2, 0, 9},
        {14, intMax, intMax, intMax, 8, 9, 0},
    }
    minTree := NewMinTree(vertexes, edges)

    // minTree.ShowGraph()

    minTree.Kruskal()
}
