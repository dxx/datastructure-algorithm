import unittest
from binary_tree_seq_storage import SeqBinaryTree


class Test(unittest.TestCase):
    
    def test_pre_order(self):
        nos = [1, 2, 3, 4, 5, 6, 7]
        seq_binary_tree = SeqBinaryTree(nos)

        print("======前序遍历======")
        seq_binary_tree.pre_order()
