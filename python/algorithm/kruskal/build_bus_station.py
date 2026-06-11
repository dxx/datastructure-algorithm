"""
克鲁斯卡尔(Kruskal)算法解决建设公交站问题
Kruskal 算法适用于稀疏图
思想：按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路
"""

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
