import unittest
from chessboard import Chessboard


class Test(unittest.TestCase):
    
    def test_chessboard(self):
        chessboard = Chessboard(8, 8)
        chessboard.move(4, 4)
