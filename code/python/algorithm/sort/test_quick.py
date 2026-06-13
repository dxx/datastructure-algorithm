import unittest
from quick import quick_sort


class Test(unittest.TestCase):
    
    def test_quick_sort(self):
        nums = [5, 1, 8, 3, 7, 2, 9, 4, 6]
        print("排序前: " + str(nums))
        quick_sort(nums, 0, len(nums) - 1)
        print("排序后: %s\n" % nums)
