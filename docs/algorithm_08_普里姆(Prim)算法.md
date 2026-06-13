## 普里姆(Prim)算法

>各种语言实现代码：[Go](./golang/algorithm/prim)   [Java](./java/algorithm/src/com/dxx/prim)   [JavaScript](./javascript/algorithm/prim)   [TypeScript](./typescript/algorithm/prim)   [Python](./python/algorithm/prim)   [Rust](./rust/algorithm/src/prim)
>
>默认使用 **Python** 语言实现。

### 简介

通过每次添加一个新节点加入集合，直到所有点加入停止的最小生成树的算法。一个有 n 个结点的连通图的生成树是原图的极小连通子图，且包含原图中的所有 n 个结点，并且有保持图连通的最少的边，也就是说所有边的权值之和最小，这样的树叫最小生成树。最小生成树可以用 prim（普里姆）算法或 kruskal（克鲁斯卡尔）算法求出。**Prim 算法适用于稠密图**。

普里姆算法思路：

1. 首先随便选一个点加入集合。

2. 用该点的所有边去刷新到其它点的最短路。

3. 找出最短路中最短的一条连接（且该点未被加入集合）。

4. 用该点去刷新到其他点的最短路。
5. 重复以上操作 n-1 次。

### 村庄修路

有 A、B、C、D、E 、F 和 G 这 7 个村庄，每个村庄之间的距离用边线的权表示，A-B 的距离为 5 公里。如何修路能保证各个村庄连通，总的修建总路程最短？

![algorithm_prim_1](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_prim_1.png)

分析：

1. 任取一个顶点，比如从顶点 **A** 开始，和 A 连接的顶点有 B、C 和 G，边的权值分别为：A-B => 5、A-C => 7、A-G => 2，最小边的权值为 2，将 **G** 加入集合。

2. 从 A、G 开始，除去已经加入到集合中的顶点，和 A 、G 连接的顶点分别有： A-C => 7、A-B => 5、G-B => 3、G-E => 4、G-F => 6，最小边的权值为 3，将 **B** 加入集合。

3. 从 A、G、B 开始，和 A 、G、B 连接且未被加入集合的顶点分别有：A-C => 7、G-E => 4、G-F => 6、B-D => 9，最小边的权值为 4，将 **E** 加入集合。

4. 从 A、G、B 、E 开始，A、G、B 、E 连接且未被加入集合的顶点分别有：A-C => 7、G-F => 6、B-D => 9、E-C => 8、E-F => 5，最小边的权值为 5，将 **F** 加入集合。

5. 从 A、G、B 、E、F 开始，A、G、B 、E 、F 连接且未被加入集合的顶点分别有：A-C => 7、B-D => 9、E-C => 8、F-D => 4，最小边的权值为 4，将 **D** 加入集合。

6. 从 A、G、B 、E、F、D 开始，A、G、B 、E 、F 、D 连接且未被加入集合的顶点分别有：A-C => 7、E-C => 8，最小边的权值为 7，将 **C** 加入集合。

7. 至此所有顶点全部加入到集合，将改集合中的顶点连接起来就是最小生成树。

8. 最后总的修路径长度为 `2 + 3 + 4 + 5 + 4 + 7 = 25`。

> 注意第一次选取的顶点不同，生成的顶点集合顺序也会不同，但最后的路径总长度相同。

最小生成树如下图：

![algorithm_prim_2](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_prim_2.png)

代码实现：

```python
MAX_VALUE = float("inf")


class Graph:
    def __init__(self, vertexes: list[str], matrix: list[list[float]]) -> None:
        self.vertexes = vertexes  # 顶点
        self.matrix = matrix  # 邻接矩阵，代表边的值


class MinTree:
    """最小生成树"""

    def __init__(self, vertexes: list[str], edges: list[list[float]]) -> None:
        initial_vertexes = vertexes[:]
        initial_matrix = [row[:] for row in edges]
        self.graph = Graph(initial_vertexes, initial_matrix)

    def prim(self, v: int) -> None:
        num_of_vertex = len(self.graph.vertexes)
        # 存放已经连通的顶点集合
        vertex_map = {}
        # 将当前顶点加入集合
        vertex_map[self.graph.vertexes[v]] = self.graph.vertexes[v]

        # n 个顶点就有 n-1 条边
        for _ in range(1, num_of_vertex):
            # 记录顶点下标
            v1 = -1
            v2 = -1
            # 记录最小边的权值，初始化成一个最大数，后续遍历中会被替换
            min_weight = MAX_VALUE
            # 查找已经加入集合中的顶点，和这些顶点中最近的一个顶点
            for i in range(num_of_vertex):
                for j in range(num_of_vertex):
                    weight = self.graph.matrix[i][j]
                    # 表示已经加入集合的顶点
                    # 表示未被加入集合的顶点
                    if (
                        vertex_map.get(self.graph.vertexes[i]) == self.graph.vertexes[i]
                        and vertex_map.get(self.graph.vertexes[j]) is None
                        and weight < min_weight
                    ):
                        v1 = i
                        v2 = j
                        # 修改最小的权值
                        min_weight = weight
            # 将最小的顶点加入到集合中
            vertex_map[self.graph.vertexes[v2]] = self.graph.vertexes[v2]
            print(f"边: {self.graph.vertexes[v1]}-{self.graph.vertexes[v2]} => {self.graph.matrix[v1][v2]:g}")

    def show_graph(self) -> None:
        for edges in self.graph.matrix:
            print("[ " + " ".join(str(edge) for edge in edges) + " ]")
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_prim(self):
        vertexes = ["A", "B", "C", "D", "E", "F", "G"]
        edges = [
            [MAX_VALUE, 5, 7, MAX_VALUE, MAX_VALUE, MAX_VALUE, 2],
            [5, MAX_VALUE, MAX_VALUE, 9, MAX_VALUE, MAX_VALUE, 3],
            [7, MAX_VALUE, MAX_VALUE, MAX_VALUE, 8, MAX_VALUE, MAX_VALUE],
            [MAX_VALUE, 9, MAX_VALUE, MAX_VALUE, MAX_VALUE, 4, MAX_VALUE],
            [MAX_VALUE, MAX_VALUE, 8, MAX_VALUE, MAX_VALUE, 5, 4],
            [MAX_VALUE, MAX_VALUE, MAX_VALUE, 4, 5, MAX_VALUE, 6],
            [2, 3, MAX_VALUE, MAX_VALUE, 4, 6, MAX_VALUE],
        ]
        min_tree = MinTree(vertexes, edges)
        # 从 A 点开始
        min_tree.prim(0)
```

运行：

```shell
❯ python -m unittest test_repair_road.Test.test_prim
边: A-G => 2
边: G-B => 3
边: G-E => 4
边: E-F => 5
边: F-D => 4
边: A-C => 7
```
