import unittest
from stack import Stack


class Test(unittest.TestCase):
    
    def test_stack(self):
        # 创建一个栈
        stack = Stack(3)
        stack.push("one")
        stack.push("two")
        stack.push("three")
        stack.show()
