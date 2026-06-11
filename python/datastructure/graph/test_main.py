import unittest
from main import Graph


class Test(unittest.TestCase):
    
    def test_graph(self):
        vertexes = ["A", "B", "C", "D", "E"]
        graph = Graph(5)

        for vertex in vertexes:
            graph.add_vertex(vertex)

        # A-B
        graph.add_edge(0, 1, 1)
        # A-C
        graph.add_edge(0, 2, 1)
        # B-C
        graph.add_edge(1, 2, 1)
        # B-E
        graph.add_edge(1, 4, 1)
        # C-D
        graph.add_edge(2, 3, 1)

        graph.show_edges()

        print("======深度优先遍历======")
        graph.dfs()

        print()

        print("======广度优先遍历======")
        graph.bfs()
