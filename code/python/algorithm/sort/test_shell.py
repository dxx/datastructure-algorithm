import unittest
from shell import shell_sort


class Test(unittest.TestCase):
    
    def test_shell_sort(self):
        nums = [5, 1, 7, 3, 2, 4, 9, 6, 8]
        print("排序前: " + str(nums))
        shell_sort(nums)
        print("排序后: " + str(nums))
