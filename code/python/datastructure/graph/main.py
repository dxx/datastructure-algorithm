"""
图
"""


class Graph:
    def __init__(self, num: int) -> None:
        self.vertexes: list[str | None] = [None] * num
        self.matrix = [[0] * num for _ in range(num)]
        self.num_of_vertex = 0
        self.num_of_edge = 0

    def add_vertex(self, vertex: str) -> None:
        """添加顶点"""
        self.vertexes[self.num_of_vertex] = vertex
        self.num_of_vertex += 1

    def add_edge(self, i1: int, i2: int, weight: int) -> None:
        """
        添加边
        i1: 第一个顶点下标
        i2: 第二个顶点下标
        weight: 权值. 0-表示不通, 1-表示通
        """
        # 在二维数组中设置权值，因为无方向图，所以两个位置都需要设置
        self.matrix[i1][i2] = weight
        self.matrix[i2][i1] = weight
        self.num_of_edge += 1

    def get_num_of_vertex(self) -> int:
        """获取顶点数量"""
        return self.num_of_vertex

    def get_num_of_edge(self) -> int:
        """获取边的数量"""
        return self.num_of_edge

    def dfs(self) -> None:
        """深度优先遍历"""
        is_visited = [False] * self.get_num_of_vertex()
        # 遍历所有的顶点，进行深度优先遍历
        for i in range(self.get_num_of_vertex()):
            if not is_visited[i]:
                self.dfs_recursion(is_visited, i)

    def dfs_recursion(self, is_visited: list[bool], v: int) -> None:
        """
        递归遍历
        1.访问初始顶点 v，并标记顶点 v 为已访问
        2.查找顶点 v 的第一个邻接顶点 w
        3.如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问
        4.如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3
        5.查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3
        """
        print(self.vertexes[v])

        # 标记已被访问
        is_visited[v] = True

        # 获取第一个邻接顶点
        w = self.get_first_vertex(v)
        # 存在则继续调用
        while w != -1:
            # 未被访问
            if not is_visited[w]:
                # 继续遍历
                self.dfs_recursion(is_visited, w)
            # 查找顶点 v 的 w 邻接顶点的下一个邻接顶点
            w = self.get_next_vertex(v, w)

    def bfs(self) -> None:
        """广度优先遍历"""
        is_visited = [False] * self.get_num_of_vertex()
        # 遍历所有的顶点，进行广度优先遍历
        for i in range(self.get_num_of_vertex()):
            if not is_visited[i]:
                self.bfs2(is_visited, i)

    def bfs2(self, is_visited: list[bool], v: int) -> None:
        """
        1.访问初始顶点 v 并标记顶点 v 为已访问。
        2.顶点 v 入队列。
        3.当队列非空时，继续执行，否则结束。
        4.出队列，取得队头结点 u。
        5.查找结点 u 的第一个邻接顶点 w。
        6.若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
        7.若顶点 w 尚未被访问，则访问顶点 w 并标记为已访问。
        8.将顶点 w 入队列。
        9.查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。
        """
        print(self.vertexes[v])

        queue: list[int] = []
        # 标记已被访问
        is_visited[v] = True
        # 将顶点入队列
        queue.insert(0, v)
        while len(queue) != 0:
            # 取出头结点下标
            u = queue.pop()
            # 获取第一个邻接节点的下标
            w = self.get_first_vertex(u)
            while w != -1:
                # 未被访问
                if not is_visited[w]:
                    print(self.vertexes[w])
                    # 标记已被访问
                    is_visited[w] = True
                    # 入队列
                    queue.insert(0, w)
                # 获取顶点 u 的继 w 邻接顶点后的下一个邻接顶点
                w = self.get_next_vertex(u, w)

    def get_first_vertex(self, i: int) -> int:
        """获取第一个邻接顶点下标"""
        for j in range(self.get_num_of_vertex()):
            if self.matrix[i][j] > 0:
                return j
        return -1

    def get_next_vertex(self, i1: int, i2: int) -> int:
        """获取下一个邻接顶点下标"""
        for j in range(i2 + 1, self.get_num_of_vertex()):
            if self.matrix[i1][j] > 0:
                return j
        return -1

    def show_edges(self) -> None:
        """显示邻接矩阵"""
        for row in self.matrix:
            print("[" + "".join(f" {value} " for value in row) + "]")
