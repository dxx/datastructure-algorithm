import unittest
from calculator import Calculator


class Test(unittest.TestCase):
    
    def test_calculate(self):
        calculator = Calculator()

        calculator.calculate("3+5*3-6")
        calculator.calculate("30+5*3-6")
        calculator.calculate("130+5*3-6")
