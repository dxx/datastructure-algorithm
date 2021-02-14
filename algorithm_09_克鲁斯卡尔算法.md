## 克鲁斯卡尔算法

>各种语言实现代码：[Go](./golang/algorithm/kruskal)   [Java](java/algorithm/src/com/mcx/kruskal)   [JavaScript](javascript/algorithm/kruskal)   [Rust](./rust/algorithm/src/kruskal)
>
>默认使用 **Go** 语言实现。

### 简介

Kruskal 算法是一种用来查找最小生成树的算法。用来解决同样问题的还有Prim 算法和 Boruvka 算法等。三种算法都是贪心算法的应用。和 Boruvka 算法不同的地方是，Kruskal 算法在图中存在相同权值的边时也有效。**Kruskal 算法适用于稀疏图**。

思想：

按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路。

### 建设公交站

某城市需要建设 7 个公交站点，现在需要修路把这 7 个站点连通，每个站点的距离用边的权值表示，如 A 到 B 距离 12 公里。如何修路保证每个站点连通，并且总的修建里程最短？

![algorithm_kruskal_1](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_1.png)

分析：

1. 根据边的权值对所有边进行排序。
2. 创建一个数组或集合 V 保存最小生成树的顶点，选择最小权值为 2 的边 E-F，将 E、F 加入 V 中。
3. 上一步之后，选择最小权值为 3 的边 C-D，将 C 、D 加入 V 中。
4. 上一步之后，选择最小权值为 4 的边 D-E，将 D、E 加入 V 中。
5. 上一步之后，最小权值为 5，对应的边为 C-E，但是 C、D、E 三点构成回路，跳过 C-E。选择 C-F，C、D、E、F 四点同样构成回路，跳过 C-F。选择 B-F，权值为 7，将 B、F 加入 V 中。
6. 上一步之后，选择最小权值为 8 的边 E-G，将 E、G 加入 V 中。
7. 上一步之后，最小权值为 9，对应的边为 F-G，构成回路跳过。选择 B-C，构成回路跳过。选择 A-B，权值为 12。
8. 此时最小生成树构建完成。其树的边依次是 E-F、C-D、D-E、B-F、E-G、A-B，权值分别为 2、3、4、7、8、12。

最小生成树如下图：

![algorithm_kruskal_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_2.png)

上述分析过程中有以下关键点：

1. 对图的所有边按照权值从小到大排序。
2. 当边添加到集合中时，如何判断是否构成回路。

按照权值从小到大排序可以使用排序算法来解决。

对于判断是否构成回路分析如下：

在将 E-F、C-D、D-E 加入到 V 中之后，这几条边的顶点就都有了终点 F。

![algorithm_kruskal_3](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_3.png)

终点就是将所有顶点按照从小到大的顺序排列好之后，某个顶点的终点就是与它连通的最大顶点。也就是说边的两个顶点不能都指向同一个终点，否则将构成回路。

代码实现：

```go
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

// int32 位最大值，使用最大值表示两个顶点不连通
const intMax = 1 << 31 - 1

func main() {
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
```

输出：

```
======边排序前======
[A-B:12 A-F:16 A-G:14 B-C:10 B-F:7 C-D:3 C-E:5 C-F:6 D-E:4 E-F:2 E-G:8 F-G:9]
======边排序后======
[E-F:2 C-D:3 D-E:4 C-E:5 C-F:6 B-F:7 E-G:8 F-G:9 B-C:10 A-B:12 A-G:14 A-F:16]
边: E-F => 2
边: C-D => 3
边: D-E => 4
边: B-F => 7
边: E-G => 8
边: A-B => 12
```
