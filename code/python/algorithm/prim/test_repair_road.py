import unittest
from repair_road import MAX_VALUE, MinTree


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
