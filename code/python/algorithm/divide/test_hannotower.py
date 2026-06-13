import unittest
from hannotower import hannotower


class Test(unittest.TestCase):
    
    def test_hannotower(self):
        hannotower(3, "A", "B", "C")
