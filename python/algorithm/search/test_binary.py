import unittest
from binary import binary_search


class Test(unittest.TestCase):
    
    def test_binary_search(self):
        value = 100
        nums = [1, 8, 10, 89, 100, 100, 123]
        index = binary_search(nums, 0, len(nums) - 1, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
