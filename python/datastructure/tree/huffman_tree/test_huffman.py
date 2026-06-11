import unittest
from huffman import create_huffman_tree, pre_order


class Test(unittest.TestCase):
    
    def test_huffman(self):
        nums = [1, 7, 3, 8, 16]
        root = create_huffman_tree(nums)
        pre_order(root)
