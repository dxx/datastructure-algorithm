import unittest
from bubble import bubble_sort, optimize_bubble_sort


class Test(unittest.TestCase):
    
    def test_bubble_sort(self):
        nums = [1, 5, 7, 3, 2, 4, 9, 6, 8]
        print("交换前: " + str(nums))
        bubble_sort(nums)
        print("交换后: " + str(nums))

    def test_optimize_bubble_sort(self):
        nums = [1, 5, 7, 3, 2, 4, 9, 6, 8]
        print("优化前: " + str(nums))
        optimize_bubble_sort(nums)
        print("优化后: " + str(nums))
