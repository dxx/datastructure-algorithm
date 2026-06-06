from collections.abc import Sequence
from typing import Callable

from stack import Stack


"""
逆波兰计算器
"""


class Operation:
    def __init__(self, opt: str, priority: int, opt_fun: Callable[[float, float], float]) -> None:
        self.opt = opt
        self.priority = priority
        self.opt_fun = opt_fun


# 定义相关操作符对应的优先级和计算方法
operations = {
    "+": Operation("+", 1, lambda num1, num2: num1 + num2),
    "-": Operation("-", 1, lambda num1, num2: num1 - num2),
    "*": Operation("*", 2, lambda num1, num2: num1 * num2),
    "/": Operation("/", 2, lambda num1, num2: num1 / num2),
}


class ReversePoland:
    """
    判断是否为数字
    """
    def _is_num(self, text: str) -> bool:
        return text.isdigit()

    """
    计算操作符的优先级
    """
    def _priority(self, opt1: str, opt2: str) -> int | None:
        operation1 = operations.get(opt1)
        operation2 = operations.get(opt2)
        if not operation1 or not operation2:
            print("请检查运算符: " + opt1 + "," + opt2)
            return None
        return operation1.priority - operation2.priority

    """
    计算结果
    """
    def _calculate_num(self, num1: float, num2: float, opt: str) -> float:
        operation = operations.get(opt)
        if operation:
            if operation.opt == "-" or operation.opt == "/":
                # 因为出栈后两数的位置颠倒，需交换两个数的位置
                num1 = num1 + num2
                num2 = num1 - num2
                num1 = num1 - num2
            return operation.opt_fun(num1, num2)
        return 0

    """
    后缀表达式计算
    1. 循环读取每个字符，判断是否是数字
    2. 如果是数字直接入栈
    3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中
    """
    def cal_suffix_expression(self, expr: Sequence[str] | None) -> float | None:
        if expr is None:
            return None
        stack = Stack(len(expr))
        for text in expr:
            if self._is_num(text):
                stack.push(text)
                continue
            operation = operations.get(text)
            if not operation:
                print("无效的运算符: " + text)
                return None
            # 计算
            num_str1 = stack.pop()
            num_str2 = stack.pop()
            result = self._calculate_num(float(str(num_str1)), float(str(num_str2)), text)
            stack.push(result)
        # 弹出最后结果
        return float(str(stack.pop()))

    """
    将表达式转换成数组
    """
    def expr_to_array(self, expr: str) -> list[str]:
        if not expr:
            return []
        expressions = []
        i = 0
        while i < len(expr):
            s = expr[i]
            if self._is_num(s):
                # 向后面继续判断是否为数字
                while i + 1 < len(expr) and self._is_num(expr[i + 1:i + 2]):
                    s += expr[i + 1:i + 2]
                    i += 1
            expressions.append(s)
            i += 1
        return expressions

    """
    中缀表达式转后缀表达式
    """
    def infix_to_suffix(self, infix: list[str]) -> list[str] | None:
        if not infix:
            return None
        # 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2
        stack = Stack(len(infix))
        # 由于中间结果栈不需要弹出元素，可以使用集合来保存
        suffixes: list[str] = []
        # 循环表达式
        for text in infix:
            # 遇到数字时，将其放入 suffixes
            if self._is_num(text):
                suffixes.append(text)
                continue
            # 如果是 ( 直接入栈
            if text == "(":
                stack.push(text)
                continue
            if text == ")":
                while not stack.is_empty() and stack.peek() != "(":
                    # 弹出 stack 中栈顶的元素，并添加到 suffixes
                    suffixes.append(str(stack.pop()))
                # 弹出 (，消除一对 ( )
                stack.pop()
                continue
            operation = operations.get(text)
            if operation:
                if stack.is_empty() or stack.peek() == "(":
                    # 如果 stack 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈
                    stack.push(text)
                    continue
                # 栈不为空，并且当前字符串的优先级小于等于栈顶的元素
                while not stack.is_empty():
                    priority = self._priority(text, str(stack.peek()))
                    if priority is None or priority > 0:
                        break
                    # 将栈顶的元素添加到 suffixes
                    suffixes.append(str(stack.pop()))
                # 入栈
                stack.push(text)
            else:
                print("无法识别的字符: " + text)
                return None
        # 将 stack 中剩余的运算符依次添加到 suffixes
        while not stack.is_empty():
            suffixes.append(str(stack.pop()))
        # 因为这里用的是集合，它里面元素的顺序就是栈元素出栈后逆序排列的顺序
        return suffixes


def main() -> None:
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


if __name__ == "__main__":
    main()
