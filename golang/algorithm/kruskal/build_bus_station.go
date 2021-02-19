package kruskal

import "fmt"

// 克鲁斯卡尔(Kruskal)算法解决建设公交站问题
// Kruskal 算法适用于稀疏图
// 思想：按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路

// 最小生成树
type MinTree struct {
    graph *Graph
}

type Graph struct {
    vertexes  []string // 顶点
    matrix    [][]int  // 邻接矩阵
    numOfEdge int      // 边的条数
}

// 边
type Edge struct {
    start  string // 起始顶点
    end    string // 结束顶点
    weight int    // 边的权值
}

func (edge *Edge) String() string {
    return fmt.Sprintf("%s-%s:%d", edge.start, edge.end, edge.weight)
}

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
    var numOfEdge int
    for i := 0; i < numOfVertex; i++ {
        // 已经统计过的边不统计
        for j := i + 1; j < numOfVertex; j++ {
            numOfEdge++
        }
    }
    graph := Graph{
        vertexes:  initialVertexes,
        matrix:     initialMatrix,
        numOfEdge: numOfEdge,
    }
    return &MinTree{graph: &graph}
}

func (minTree *MinTree) Kruskal() {
    // 保存最小生成树的边
    edges := make([]*Edge, 0)

    // 保存已存在最小生成树中每个顶点对应的在树中的终点下标
    endPosits := make([]int, minTree.graph.numOfEdge)

    // 获取所有边
    allEdges := minTree.getEdges()
    fmt.Println("======边排序前======")
    fmt.Println(allEdges)

    minTree.sortEdges(allEdges)
    fmt.Println("======边排序后======")
    fmt.Println(allEdges)

    // 遍历所以的边
    for _, v := range allEdges {
        // 获取边的起始顶点下标
        startPosit := minTree.getVertexPosit(v.start)
        // 获取边的结束顶点下标
        endPosit := minTree.getVertexPosit(v.end)

        // 获取起始顶点的终点下标
        endPosit1 := minTree.getEndPosit(endPosits, startPosit)
        // 获取结束顶点的终点下标
        endPosit2 := minTree.getEndPosit(endPosits, endPosit)
        // 判断是否形成回路
        if endPosit1 != endPosit2 { // 没有构成回路
            // 设置 endPosit1 的终点下标 endPosit2
            // [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
            endPosits[endPosit1] = endPosit2
            edges = append(edges, v) // 将改变加入最小生成树
        }
    }

    // 输出最小生成树
    for _, edge := range edges {
        fmt.Printf("边: %s-%s => %d\n",
            edge.start,
            edge.end,
            edge.weight)
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

// 获取顶点在顶点集合中的位置
// vertex: 顶点
func (minTree *MinTree) getVertexPosit(vertex string) int {
    for i, v := range minTree.graph.vertexes {
        if vertex == v {
            return i
        }
    }
    return -1
}

// 获取指定下标顶点的终点下标
// posits: 存放顶点和对应终点下标，posits 的下标表示顶点下标，值表示对应顶点的终点下标
// i: 目标顶点下标
// posits = [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0], i = 4 返回 5
func (minTree *MinTree) getEndPosit(posits []int, i int) int {
    for posits[i] != 0 {
        i = posits[i]
    }
    return i
}

// int32 位最大值，使用最大值表示两个顶点不连通
const intMax = 1 << 31 - 1

// 获取所有边
func (minTree *MinTree) getEdges() []*Edge {
    edges := make([]*Edge, 0)
    numOfVertex := len(minTree.graph.vertexes)
    for i := 0; i < numOfVertex; i++ {
        for j := i + 1; j < numOfVertex; j++ {
            // 不连通的跳过
            if minTree.graph.matrix[i][j] == intMax {
                continue
            }
            // 创建边
            edges = append(edges, &Edge{
                start:  minTree.graph.vertexes[i],
                end:    minTree.graph.vertexes[j],
                weight: minTree.graph.matrix[i][j],
            })
        }
    }
    return edges
}

// 对边按照权值从小到大进行排序
func (minTree *MinTree) sortEdges(edges []*Edge) {
    if edges == nil {
        return
    }
    for i := 1; i < len(edges); i++ {
        insertIndex := i - 1
        insertValue := edges[i]
        for insertIndex >= 0 && edges[insertIndex].weight > insertValue.weight {
            edges[insertIndex + 1] = edges[insertIndex]
            insertIndex--
        }
        if insertIndex + 1 != i {
            edges[insertIndex + 1] = insertValue
        }
    }
}
