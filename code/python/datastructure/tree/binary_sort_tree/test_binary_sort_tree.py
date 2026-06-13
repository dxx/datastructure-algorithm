import unittest
from binary_sort_tree import BinarySortTree, BinaryTreeNode


class Test(unittest.TestCase):
    
    def test_binary_sort_tree(self):
        nos = [8, 5, 10, 3, 6, 9, 12, 2]
        binary_sort_tree = BinarySortTree()
        for no in nos:
            binary_sort_tree.add(BinaryTreeNode(no))
        print("======中序遍历======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(6)

        print("======删除叶子节点 6======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(5)

        print("======删除只有一颗子树的节点 5======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(10)

        print("======删除有两颗子树的节点 10======")
        binary_sort_tree.infix_order()
