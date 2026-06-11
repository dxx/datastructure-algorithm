import unittest
from binary_no_recursion import binary_search_no_recursion


class Test(unittest.TestCase):
    
    def test_binary_search_no_recursion(self):
        value = 100
        nums = [1, 8, 10, 89, 100, 100, 123]
        index = binary_search_no_recursion(nums, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
