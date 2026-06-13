import unittest
from reverse_poland import ReversePoland


class Test(unittest.TestCase):
    
    def test_reverse_poland(self):
        reverse_poland = ReversePoland()
        expr = "3 5 3 * + 2 -"
        # 假设数和数或符号之间有空格
        expressions = expr.split(" ")
        result = reverse_poland.cal_suffix_expression(expressions)
        print("后缀表达式 " + expr + " 的计算结果为: " + str(result))

        expr = "1+((2+3)*4)-5"
        expressions = reverse_poland.expr_to_array(expr)
        print("将中缀表达式放入数组, 结果为: " + str(expressions))

        expressions = reverse_poland.infix_to_suffix(expressions)
        print("中缀表达式转换成后缀表达式, 结果为: " + str(expressions))

        result = reverse_poland.cal_suffix_expression(expressions)
        print("计算表达式 " + expr + ", 结果为: " + str(result))
