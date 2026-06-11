import unittest
from main import test_delete_node, test_insert_at_tail, test_sort_insert_by_no


class Test(unittest.TestCase):
    
    def test_insert_at_tail(self):
        test_insert_at_tail()

    def test_sort_insert_by_no(self):
        test_sort_insert_by_no()

    def test_delete_node(self):
        test_delete_node()
