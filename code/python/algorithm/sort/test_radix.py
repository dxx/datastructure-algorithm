import unittest
from radix import radix_sort


class Test(unittest.TestCase):
    
    def test_radix_sort(self):
        nums = [5, 1, 7, 13, 21, 32, 9, 66, 8, 20]
        print("排序前: " + str(nums))
        radix_sort(nums)
        print("排序后: " + str(nums))
