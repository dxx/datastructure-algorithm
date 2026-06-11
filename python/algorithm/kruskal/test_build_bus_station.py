import unittest
from build_bus_station import MAX_VALUE, MinTree


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
