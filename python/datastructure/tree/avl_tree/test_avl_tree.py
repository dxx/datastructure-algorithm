import unittest
from avl_tree import test_double_rotate, test_left_rotate, test_right_rotate


class Test(unittest.TestCase):
    
    def test_left_rotate(self):
        test_left_rotate()

    def test_right_rotate(self):
        test_right_rotate()

    def test_double_rotate(self):
        test_double_rotate()
