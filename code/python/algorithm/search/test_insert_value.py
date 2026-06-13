import unittest
from insert_value import insert_val_search


class Test(unittest.TestCase):
    
    def test_insert_val_search(self):
        nums = []
        # 填充 1 - 100
        for i in range(1, 101):
            nums.append(i)
        value = 58
        index = insert_val_search(nums, 0, len(nums) - 1, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
