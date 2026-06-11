import unittest
from heap_sort import heap_sort


class Test(unittest.TestCase):
    
    def test_heap_sort(self):
        nums = [1, 7, 5, 2, 8]

        print("排序前: " + str(nums))

        heap_sort(nums)

        print("排序后: " + str(nums))
