import unittest
from sequence import sequence_search


class Test(unittest.TestCase):
    
    def test_sequence_search(self):
        value = 8
        nums = [2, 5, 1, 7, 8, 16]
        index = sequence_search(nums, value)
        if index != -1:
            print(f"{value} 在 nums 中的下标为: {index}")
