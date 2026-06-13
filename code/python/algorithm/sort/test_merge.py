import unittest
from merge import merge_sort


class Test(unittest.TestCase):
    
    def test_merge_sort(self):
        nums = [5, 0, 1, 7, 3, 2, 4, 9, 6, 8]
        print("排序前: " + str(nums))
        merge_sort(nums)
        print("排序后: " + str(nums))
