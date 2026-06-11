import unittest
from eight_queen import EightQueen


class Test(unittest.TestCase):
    
    def test_eight_queen(self):
        eight_queen = EightQueen()
        eight_queen.put_queen(0)
