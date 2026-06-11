import unittest
from mi_gong import walk, walk2


def print_map(mi_gong_map: list[list[int]]) -> None:
    result = ""
    for row in mi_gong_map:
        for value in row:
            result += " " + str(value)
        result += "\n"
    print(result)


class Test(unittest.TestCase):
    
    def test_walk(self):
        # 初始化地图，0 表示通道，1 表示墙
        mi_gong_map = [
            [1, 1, 1, 1, 1, 1, 1, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 1, 1, 1, 1, 1],
        ]

        print("探路之前:")
        print_map(mi_gong_map)
        # 开始探路, 起点为1,1
        walk(mi_gong_map, 1, 1)
        print("探路之后:")
        print_map(mi_gong_map)

    def test_walk2(self):
        # 初始化地图，0 表示通道，1 表示墙
        mi_gong_map = [
            [1, 1, 1, 1, 1, 1, 1, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 1, 1, 1, 1, 1],
        ]

        print("探路之前:")
        print_map(mi_gong_map)

        # 开始探路, 起点为1,1
        walk2(mi_gong_map, 1, 1)

        print("探路之后:")
        print_map(mi_gong_map)
