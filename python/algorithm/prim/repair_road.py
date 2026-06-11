"""
普里姆(Prim)算法解决村庄修路问题。
Prim 算法适用于稠密图
算法思路：
1.首先随便选一个点加入集合。
2.用该点的所有边去刷新到其它点的最短路。
3.找出最短路中最短的一条连接（且该点未被加入集合）。
4.用该点去刷新到其他点的最短路。
5.重复以上操作 n-1 次。
"""

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
