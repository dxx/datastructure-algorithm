package main

import "fmt"

// 普里姆(Prim)算法解决村庄修路问题。
// Prim 算法适用于稠密图
// 算法思路：
// 1.首先随便选一个点加入集合。
// 2.用该点的所有边去刷新到其它点的最短路。
// 3.找出最短路中最短的一条连接（且该点未被加入集合）。
// 4.用该点去刷新到其他点的最短路。
// 5.重复以上操作 n-1 次。

// 最小生成树
type MinTree struct {
    graph *Graph
}

type Graph struct {
    vertexes []string // 顶点
    matrix    [][]int  // 领接矩阵，代表边的值
}

// 创建图
func NewMinTree(vertexes []string, edges [][]int) *MinTree {
    numOfVertex := len(vertexes)
    initialVertexes := make([]string, numOfVertex)
    initialMatrix := make([][]int, numOfVertex)
    for i := 0; i < numOfVertex; i++ {
        initialVertexes[i] = vertexes[i]
        initialMatrix[i] = make([]int, numOfVertex)
        for j := 0; j < numOfVertex; j++ {
            initialMatrix[i][j] = edges[i][j]
        }
    }
    graph := Graph{vertexes: initialVertexes, matrix: initialMatrix}
    return &MinTree{graph: &graph}
}

func (minTree *MinTree) Prim(v int) {
    numOfVertex := len(minTree.graph.vertexes)
    // 存放已经连通的顶点集合
    vertexMap := make(map[string]string, numOfVertex)
    // 将当前顶点加入集合
    vertexMap[minTree.graph.vertexes[v]] = minTree.graph.vertexes[v]

    // 记录顶点下标
    v1 := -1
    v2 := -1
    // 记录最小边的权值，初始化成一个最大数，后续遍历中会被替换
    minWeight := intMax
    // n 个顶点就有 n-1 条边
    for k := 1; k < numOfVertex; k++ {
        // 查找已经加入集合中的顶点，和这些顶点中最近的一个顶点
        for i := 0; i < numOfVertex; i++ {
            for j := 0; j < numOfVertex; j++ {
                weight := minTree.graph.matrix[i][j]
                if vertexMap[minTree.graph.vertexes[i]] == minTree.graph.vertexes[i] && // 表示已经加入集合的顶点
                    vertexMap[minTree.graph.vertexes[j]] == "" && // 表示未被加入集合的顶点
                    weight < minWeight {
                    v1 = i
                    v2 = j
                    minWeight = weight
                }
            }
        }
        // 将最小的顶点加入到集合中
        vertexMap[minTree.graph.vertexes[v2]] = minTree.graph.vertexes[v2]
        // 修改最小的权值
        minWeight = intMax

        fmt.Printf("边:%s-%s => %d\n",
            minTree.graph.vertexes[v1],
            minTree.graph.vertexes[v2],
            minTree.graph.matrix[v1][v2])
    }
}

func (minTree *MinTree) ShowGraph() {
    for _, edges := range minTree.graph.matrix {
        fmt.Printf("[ ")
        for _, val := range edges {
            fmt.Printf("%10d ", val)
        }
        fmt.Printf("]\n")
    }
}

// int32 位最大值，使用最大值表示两个顶点不连通
const intMax = 1 << 31 - 1

func main() {
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
