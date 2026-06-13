## 克鲁斯卡尔算法

>各种语言实现代码：[Go](./golang/algorithm/kruskal)   [Java](java/algorithm/src/com/dxx/kruskal)   [JavaScript](javascript/algorithm/kruskal)   [TypeScript](./typescript/algorithm/kruskal)   [Python](./python/algorithm/kruskal)   [Rust](./rust/algorithm/src/kruskal)
>
>默认使用 **Python** 语言实现。

### 简介

Kruskal 算法是一种用来查找最小生成树的算法。用来解决同样问题的还有Prim 算法和 Boruvka 算法等。三种算法都是贪心算法的应用。和 Boruvka 算法不同的地方是，Kruskal 算法在图中存在相同权值的边时也有效。**Kruskal 算法适用于稀疏图**。

思想：

按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路。

### 建设公交站

某城市需要建设 7 个公交站点，现在需要修路把这 7 个站点连通，每个站点的距离用边的权值表示，如 A 到 B 距离 12 公里。如何修路保证每个站点连通，并且总的修建里程最短？

![algorithm_kruskal_1](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_1.png)

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

![algorithm_kruskal_2](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_2.png)

上述分析过程中有以下关键点：

1. 对图的所有边按照权值从小到大排序。
2. 当边添加到集合中时，如何判断是否构成回路。

按照权值从小到大排序可以使用排序算法来解决。

对于判断是否构成回路分析如下：

在将 E-F、C-D、D-E 加入到 V 中之后，这几条边的顶点就都有了终点 F。

![algorithm_kruskal_3](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_kruskal_3.png)

终点就是将所有顶点按照从小到大的顺序排列好之后，某个顶点的终点就是与它连通的最大顶点。也就是说边的两个顶点不能都指向同一个终点，否则将构成回路。

代码实现：

```python
MAX_VALUE = float("inf")


class Graph:
    def __init__(self, vertexes: list[str], matrix: list[list[float]], num_of_edge: int) -> None:
        self.vertexes = vertexes  # 顶点
        self.matrix = matrix  # 邻接矩阵，代表边的值
        self.num_of_edge = num_of_edge  # 边的条数


class Edge:
    """边"""

    def __init__(self, start: str, end: str, weight: float) -> None:
        self.start = start  # 起始顶点
        self.end = end  # 结束顶点
        self.weight = weight  # 边的权值

    def __repr__(self) -> str:
        return f"Edge(start={self.start!r}, end={self.end!r}, weight={self.weight!r})"


class MinTree:
    """最小生成树"""

    def __init__(self, vertexes: list[str], edges: list[list[float]]) -> None:
        num_of_vertex = len(vertexes)
        initial_vertexes = vertexes[:]
        initial_matrix = [row[:] for row in edges]
        num_of_edge = 0
        for i in range(num_of_vertex):
            # 已经统计过的边不统计
            for _ in range(i + 1, num_of_vertex):
                num_of_edge += 1
        self.graph = Graph(initial_vertexes, initial_matrix, num_of_edge)

    def kruskal(self) -> None:
        # 保存最小生成树的边
        edges = []
        # 保存已存在最小生成树中每个顶点对应的在树中的终点下标
        end_posits = [0 for _ in range(self.graph.num_of_edge)]
        # 获取所有边
        all_edges = self.get_edges()
        print("======边排序前======")
        print(all_edges)
        # 对边按照权值从小到大进行排序
        all_edges.sort(key=lambda edge: edge.weight)
        print("======边排序后======")
        print(all_edges)

        # 遍历所有的边
        for edge in all_edges:
            # 获取边的起始顶点下标
            start_posit = self.get_vertex_posit(edge.start)
            # 获取边的结束顶点下标
            end_posit = self.get_vertex_posit(edge.end)
            # 获取起始顶点的终点下标
            end_posit1 = self.get_end_posit(end_posits, start_posit)
            # 获取结束顶点的终点下标
            end_posit2 = self.get_end_posit(end_posits, end_posit)
            # 判断是否形成回路
            if end_posit1 != end_posit2:  # 没有构成回路
                # 设置 endPosit1 的终点下标 endPosit2
                end_posits[end_posit1] = end_posit2
                # [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
                edges.append(edge)  # 将改变加入最小生成树

        # 输出最小生成树
        for edge in edges:
            print(f"边: {edge.start}-{edge.end} => {edge.weight:g}")

    def get_vertex_posit(self, vertex: str) -> int:
        """
        获取顶点在顶点集合中的位置
        vertex 顶点
        """
        for i, item in enumerate(self.graph.vertexes):
            if vertex == item:
                return i
        return -1

    def get_end_posit(self, posits: list[int], i: int) -> int:
        """
        获取指定下标顶点的终点下标
        posits: 存放顶点和对应终点下标，posits 的下标表示顶点下标，值表示对应顶点的终点下标
        i: 目标顶点下标
        posits = [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0], i = 4 返回 5
        """
        while posits[i]:
            i = posits[i]
        return i

    def get_edges(self) -> list[Edge]:
        """获取所有的边"""
        edges = []
        num_of_vertex = len(self.graph.vertexes)
        for i in range(num_of_vertex):
            for j in range(i + 1, num_of_vertex):
                # 不连通的跳过
                if self.graph.matrix[i][j] == MAX_VALUE:
                    continue
                # 创建边
                edges.append(Edge(self.graph.vertexes[i], self.graph.vertexes[j], self.graph.matrix[i][j]))
        return edges

    def show_graph(self) -> None:
        for edges in self.graph.matrix:
            print("[ " + " ".join(str(edge) for edge in edges) + " ]")
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_kruskal(self):
        vertexes = ["A", "B", "C", "D", "E", "F", "G"]
        # 0-表示自己跟自己不连通，intMax-表示跟其它顶点不连通
        edges = [
            [0, 12, MAX_VALUE, MAX_VALUE, MAX_VALUE, 16, 14],
            [12, 0, 10, MAX_VALUE, MAX_VALUE, 7, MAX_VALUE],
            [MAX_VALUE, 10, 0, 3, 5, 6, MAX_VALUE],
            [MAX_VALUE, MAX_VALUE, 3, 0, 4, MAX_VALUE, MAX_VALUE],
            [MAX_VALUE, MAX_VALUE, 5, 4, 0, 2, 8],
            [16, 7, 6, MAX_VALUE, 2, 0, 9],
            [14, MAX_VALUE, MAX_VALUE, MAX_VALUE, 8, 9, 0],
        ]
        min_tree = MinTree(vertexes, edges)
        min_tree.kruskal()
```

运行：

```shell
❯ python -m unittest test_build_bus_station.Test.test_kruskal
======边排序前======
[A-B:12, A-F:16, A-G:14, B-C:10, B-F:7, C-D:3, C-E:5, C-F:6, D-E:4, E-F:2, E-G:8, F-G:9]
======边排序后======
[E-F:2, C-D:3, D-E:4, C-E:5, C-F:6, B-F:7, E-G:8, F-G:9, B-C:10, A-B:12, A-G:14, A-F:16]
边: E-F => 2
边: C-D => 3
边: D-E => 4
边: B-F => 7
边: E-G => 8
边: A-B => 12
```
