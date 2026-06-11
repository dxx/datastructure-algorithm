"""
马踏棋盘算法
国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，
按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最
终使得“马”走遍棋盘 64 个方格
"""

import time


class Point:
    """位置"""

    def __init__(self, x: int, y: int) -> None:
        self.x = x  # X 下标
        self.y = y  # Y 下标


class Chessboard:
    """棋盘"""

    def __init__(self, row: int, col: int) -> None:
        length = row * col
        self.row = row  # 表示棋盘的行数
        self.col = col  # 表示棋盘的列数
        self.visited = [False for _ in range(length)]  # 标记点是否被访问过
        self.steps = [[0 for _ in range(col)] for _ in range(row)]  # 存放步数
        self.finished = False  # 表示是否已经走完

    def move(self, start_x: int, start_y: int) -> None:
        """
        移动
        startX 起始横坐标。从 0 开始
        startY 起始纵坐标。从 0 开始
        """
        print("开始走马踏棋")
        start_time = time.time()
        self.traversal(start_x, start_y, 1)
        print("马踏棋结束")
        end_time = time.time()
        print("耗时:%fs" % (end_time - start_time))

        for i in range(self.row):
            print(" ".join(str(self.steps[i][j]) for j in range(self.col)))

    def traversal(self, x: int, y: int, step: int) -> None:
        """骑士周游算法"""
        # 将当前位置标记已访问。y = 4, col = 8, x = 4 => 4 * 8 + 4 = 36
        self.visited[y * self.col + x] = True
        # 记录步数
        self.steps[y][x] = step
        # 获取下一步可以走的所有位置
        points = self.next_points(Point(x, y))
        # 排序优化，优先选择下一步最少可走数目的位置，减少回溯，体现出贪心算法的特点
        # 去掉此方法，算法的耗时会很久，根据走法不同，结果也会不同
        points.sort(key=lambda point: len(self.next_points(point)))

        while len(points) != 0:
            # 取出第一个点
            p = points.pop(0)
            # 该点未被访问过
            if not self.visited[p.y * self.col + p.x]:
                # 继续往下走
                self.traversal(p.x, p.y, step + 1)
        # 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置
        if step < self.row * self.col and not self.finished:
            self.visited[y * self.col + x] = False
            self.steps[y][x] = 0
        else:
            self.finished = True

    def next_points(self, point: Point) -> list[Point]:
        """获取当位置的下一步可走位置的集合，最多可有 8 个位置"""
        points = []
        # 判断是否可以走 1 的位置
        if point.x - 1 >= 0 and point.y - 2 >= 0:
            points.append(Point(point.x - 1, point.y - 2))
        # 判断是否可以走 2 的位置
        if point.x + 1 < self.col and point.y - 2 >= 0:
            points.append(Point(point.x + 1, point.y - 2))
        # 判断是否可以走 3 的位置
        if point.x + 2 < self.col and point.y - 1 >= 0:
            points.append(Point(point.x + 2, point.y - 1))
        # 判断是否可以走 4 的位置
        if point.x + 2 < self.col and point.y + 1 < self.row:
            points.append(Point(point.x + 2, point.y + 1))
        # 判断是否可以走 5 的位置
        if point.x + 1 < self.col and point.y + 2 < self.row:
            points.append(Point(point.x + 1, point.y + 2))
        # 判断是否可以走 6 的位置
        if point.x - 1 >= 0 and point.y + 2 < self.row:
            points.append(Point(point.x - 1, point.y + 2))
        # 判断是否可以走 7 的位置
        if point.x - 2 >= 0 and point.y + 1 < self.row:
            points.append(Point(point.x - 2, point.y + 1))
        # 判断是否可以走 8 的位置
        if point.x - 2 >= 0 and point.y - 1 >= 0:
            points.append(Point(point.x - 2, point.y - 1))
        return points
