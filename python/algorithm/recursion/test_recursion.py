import unittest
from recursion import factorial


class Test(unittest.TestCase):
    
    def test_factorial(self):
        res = factorial(5)
        print(res)  # 120
