package main

import (
    "errors"
    "fmt"
)

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
// 1.访问初始顶点 v，并标记顶点 v 为已访问
// 2.查找顶点 v 的第一个邻接顶点 w
// 3.如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问
// 4.如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3
// 5.查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3
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

// 1.访问初始顶点 v 并标记顶点 v 为已访问。
// 2.顶点 v 入队列。
// 3.当队列非空时，继续执行，否则结束。
// 4.出队列，取得队头结点 u。
// 5.查找结点 u 的第一个邻接顶点 w。
// 6.若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
// 7.若顶点 w 尚未被访问，则访问顶点 w 并标记为已访问。
// 8.将顶点 w 入队列。
// 9.查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。
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
        // 获取第一个邻接节点的下标
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

type IntQueue struct {
    array   []int // 存放队列元素的切片（数组无法使用变量来定义长度）
    maxSize int   // 最大队列元素大小
    front   int   // 队头指针
    rear    int   // 队尾指针
}

func NewQueue(size int) *IntQueue {
    return &IntQueue{
        array:   make([]int, size),
        maxSize: size,
        front:   0,
        rear:    0,
    }
}

// 放入队列元素
func (q *IntQueue) Put(elem int) error {
    // 队尾指针不能超过最大队列元素大小
    if q.rear >= q.maxSize {
        return errors.New("queue is full")
    }
    q.array[q.rear] = elem
    q.rear++ // 队尾指针加一
    return nil
}

// 取队列元素
func (q *IntQueue) Take() (int, error) {
    // 队头指针等于队尾指针表示队列为空
    if q.front == q.rear {
        return 0, errors.New("queue is empty")
    }
    elem := q.array[q.front]
    q.front++ // 队头指针加一
    return elem, nil
}

func (q *IntQueue) IsEmpty() bool {
    return q.front == q.rear
}

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

    fmt.Println("======深度优先遍历======")
    graph.DFS()

    fmt.Println()

    fmt.Println("======广度优先遍历======")
    graph.BFS()
}
