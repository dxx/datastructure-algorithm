## 图

### 简介

在数据的逻辑结构，如果结构中的某一个节点的前驱和后继的个数不加限制，则称这种数据结构为图形结构。图形结构是一种比树形结构更复杂的非线性结构。在树形结构中，结点间具有分支层次关系，每一层上的结点只能和上一层中的至多一个结点相关，但可能和下一层的多个结点相关。而在图形结构中，任意两个结点之间都可能相关，即结点之间的邻接关系可以是任意的。

### 常用概念

顶点：A、B、C、...

边：顶点和顶点之间的连线

![data_structure_graph_01_1](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_1.png)

无向图：边没有方向的图称为无向图

![data_structure_graph_01_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_2.png)

有向图：顶点和顶点之间的边有方向

![data_structure_graph_01_3](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_3.png)

带权图：边带有权值，权值指的是边的值

![data_structure_graph_01_4](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_4.png)

### 存储方式

**邻接矩阵**

使用二维数组来存储图信息，一维数组存放图中顶点，二维数组存储各顶点之间的关系。

![data_structure_graph_02](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_02.png)

**邻接表**

给图中的每个顶点独自建立一个链表，用节点存储该顶点，用链表中其他节点存储各自的临界点。

![data_structure_graph_03](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_03.png)

### 实现

定义结构体：

```go
type Graph struct {
    vertexes  []string // 顶点
    matrix     [][]int  // 邻接矩阵。0-不通，1-通
    numOfEdge int      // 边的数目
}

// 创建图
func NewGraph(num int) *Graph {
    var vertexes []string
    matrix := make([][]int, num)
    for i := 0; i < len(matrix); i++ {
        matrix[i] = make([]int, num)
    }
    return &Graph{vertexes: vertexes, matrix: matrix}
}
```

添加顶点和边的方法：

```go
// 添加顶点
func (graph *Graph) AddVertex(vertex string) {
    graph.vertexes = append(graph.vertexes, vertex)
}

// 添加边
// i1: 第一个顶点下标
// i2: 第二个顶点下标
// weight: 权值. 0-表示不通, 1-表示通
func (graph *Graph) AddEdge(i1, i2, weight int) {
    // 在二维数组中设置权值，因为无方向图，所以两个位置都需要设置
    graph.matrix[i1][i2] = weight
    graph.matrix[i2][i1] = weight
    graph.numOfEdge++
}

// 获取顶点数量
func (graph *Graph) GetNumOfVertex() int {
    return len(graph.vertexes)
}

// 获取边的数量
func (graph *Graph) GetNumOfEdge() int {
    return graph.numOfEdge
}
```

显示邻接矩阵方法：

```go
// 显示邻接矩阵
func (graph *Graph) ShowEdges() {
    for i := 0; i < len(graph.matrix); i++ {
        fmt.Print("[")
        for j := 0; j < len(graph.matrix[i]); j++ {
            fmt.Printf(" %d ", graph.matrix[i][j])
        }
        fmt.Println("]")
    }
}
```

测试代码：

```go
func main() {
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
}
```

输出：

```
[ 0  1  1  0  0 ]
[ 1  0  1  0  1 ]
[ 1  1  0  1  0 ]
[ 0  0  1  0  0 ]
[ 0  1  0  0  0 ]
```

### 深度优先遍历

深度优先遍历从某个顶点出发，访问此顶点 v，然后从 v 的未被访问的邻接点触发深度优先遍历图，直至所有和 v 有路径想通的顶点都被访问到。

![data_structure_graph_01_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_2.png)

上图中添加顶点的顺序为 A B C D E。假设 A 为初始顶点，对上图进行深度优先遍历后就是 A -> B -> C -> D -> E。

步骤：

1. 访问初始顶点 v，并标记顶点 v 为已访问。
2. 查找顶点 v 的第一个邻接顶点 w。
3. 如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问。
4. 如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3。
5. 查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3。

代码实现：

```go
// 深度优先遍历
func (graph *Graph) DFS() {
    isVisited := make([]bool, graph.GetNumOfVertex())
    // 遍历所有的顶点，进行深度优先遍历
    for i := 0; i < graph.GetNumOfVertex(); i++ {
        if isVisited[i] == false {
            graph.dfsRecursion(isVisited, i)
        }
    }
}

// 递归遍历
func (graph *Graph) dfsRecursion(isVisited []bool, v int) {
    fmt.Printf("%s->", graph.vertexes[v])

    // 标记已被访问
    isVisited[v] = true

    // 获取第一个邻接顶点
    w := graph.getFirstVertex(v)
    // 存在则继续调用
    for w != -1 {
        // 未被访问
        if isVisited[w] == false {
            // 继续遍历
            graph.dfsRecursion(isVisited, w)
        }
        // 查找顶点 v 的 w 邻接顶点的下一个邻接顶点
        w = graph.getNextVertex(v, w)
    }
}

// 获取第一个邻接顶点下标
func (graph *Graph) getFirstVertex(i int) int {
    for j := 0; j < len(graph.vertexes); j++ {
        if graph.matrix[i][j] > 0 {
            return j
        }
    }
    return -1
}

// 获取下一个邻接顶点下标
func (graph *Graph) getNextVertex(i1, i2 int) int {
    for j := i2 + 1; j < len(graph.vertexes); j++ {
        if graph.matrix[i1][j] > 0 {
            return j
        }
    }
    return -1
}
```

测试代码：

```go
func main() {
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

    fmt.Println("======深度优先遍历======")
    graph.DFS()
}
```

输出：

```
======深度优先遍历======
A->B->C->D->E->
```

### 广度优先遍历

从图中某顶点v出发，在访问了 v 之后依次访问 v 的各个未曾访问过的邻接点，然后分别从这些邻接点出发依次访问它们的邻接点，并使得先被访问的顶点的邻接点先于后被访问的顶点的邻接点被访问，直至图中所有已被访问的顶点的邻接点都被访问到。如果此时图中尚有顶点未被访问，则需要另选一个未曾被访问过的顶点作为新的起始点，重复上述过程，直至图中所有顶点都被访问到为止。

![data_structure_graph_01_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_graph_01_2.png)

上图中添加顶点的顺序为 A B C D E。假设 A 为初始顶点，对上图进行深度优先遍历后就是 A -> B -> C -> E -> D。

步骤：

1. 访问初始顶点 v 并标记顶点 v 为已访问。
2. 顶点 v 入队列。
3. 当队列非空时，继续执行，否则结束。
4. 出队列，取得队头结点 u。
5. 查找结点 u 的第一个邻接顶点 w。
6. 若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
7. 若顶点点 w 尚未被访问，则访问顶点点 w 并标记为已访问。
8. 顶点 w 入队列。
9. 查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。

代码实现：

```go
// 广度优先遍历
func (graph *Graph) BFS() {
    isVisited := make([]bool, graph.GetNumOfVertex())
    // 遍历所有的顶点，进行广度优先遍历
    for i := 0; i < graph.GetNumOfVertex(); i++ {
        if isVisited[i] == false {
            graph.bfs(isVisited, i)
        }
    }
}

func (graph *Graph) bfs(isVisited []bool, v int) {
    fmt.Printf("%s->", graph.vertexes[v])

    queue := NewQueue(graph.GetNumOfVertex())
    // 标记已被访问
    isVisited[v] = true
    // 将顶点入队列
    _ = queue.Put(v)
    for !queue.IsEmpty() {
        // 取出头结点下标
        u, _ := queue.Take()
        // 获取第一个邻接节点的小标
        w := graph.getFirstVertex(u)
        for w != -1 {
            // 未被访问
            if isVisited[w] == false {
                fmt.Printf("%s->", graph.vertexes[w])
                // 标记已被访问
                isVisited[w] = true
                // 入队列
                _ = queue.Put(w)
            }
            // 获取顶点 u 的继 w 邻接顶点后的下一个邻接顶点
            w = graph.getNextVertex(u, w)
        }
    }
}
```

测试代码：

```go
func main() {
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

    fmt.Println("======广度优先遍历======")
    graph.BFS()
}
```

输出：

```
======广度优先遍历======
A->B->C->E->D->
```
