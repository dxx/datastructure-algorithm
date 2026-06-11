import unittest
from select import select_sort


class Test(unittest.TestCase):
    
    def test_select_sort(self):
        nums = [3, 5, 7, 1, 2, 4, 9, 6, 8]
        print("排序前: " + str(nums))
        select_sort(nums)
        print("排序后: " + str(nums))
