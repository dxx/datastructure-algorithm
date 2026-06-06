from stack import Stack


"""
综合计算器
"""


class Operation:
    def __init__(self, opt, priority, opt_fun):
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


class Calculator:
    def __init__(self):
        self.num_stack = Stack(10)
        self.operation_stack = Stack(10)

    """
    判断是否是操作符号
    """
    def _is_operation(self, opt):
        return opt in operations

    """
    判断是否为数字
    """
    def _is_num(self, text):
        return text.isdigit()

    """
    计算操作符的优先级
    """
    def _priority(self, opt1, opt2):
        operation1 = operations.get(opt1)
        operation2 = operations.get(opt2)
        if not operation1 or not operation2:
            print("请检查运算符: " + opt1 + "," + opt2)
            return None
        return operation1.priority - operation2.priority

    """
    计算结果
    """
    def _calculate_num(self, num1, num2, opt):
        operation = operations.get(opt)
        if operation:
            if operation.opt == "-" or operation.opt == "/":
                # 因为出栈后两数的位置颠倒，需交换两个数的位置
                num1 = num1 + num2
                num2 = num1 - num2
                num1 = num1 - num2
            return operation.opt_fun(num1, num2)
        return 0

    def _calculate_num_from_stack(self):
        # 从数栈中弹出两个数，从符号栈中弹出一个符号
        num_str1 = self.num_stack.pop()
        num_str2 = self.num_stack.pop()
        opt = self.operation_stack.pop()
        # 计算值
        return self._calculate_num(float(str(num_str1)), float(str(num_str2)), str(opt))

    """
    计算表达式的值
    """
    def calculate(self, expression):
        if not expression:
            return
        index = 0
        number = ""
        while index < len(expression):
            char = expression[index:index + 1]
            # 判断是否为符号
            if self._is_operation(char):
                # 判断符号栈是否为空
                if self.operation_stack.is_empty():
                    # 压入符号栈
                    self.operation_stack.push(char)
                else:
                    # 符号栈不为空，判断优先级
                    opt = self.operation_stack.peek()
                    # char 优先级小于等于 elem
                    priority = self._priority(char, str(opt))
                    if priority is not None and priority <= 0:
                        # 计算值
                        result = self._calculate_num_from_stack()
                        # 将计算结果入数栈
                        self.num_stack.push(result)
                    # 将当前操作符入符号栈
                    self.operation_stack.push(char)
            elif self._is_num(char):
                # 向后面再取一位判断是否为数字
                if index + 1 < len(expression) and self._is_num(expression[index + 1:index + 2]):
                    number += char
                    index += 1
                    continue
                self.num_stack.push(number + char)
                number = ""
            else:
                print("无法识别的字符:" + char)
                return
            index += 1

        # 全部数和符号都压入对应的栈后，取出计算
        # 符号栈不为空，循环
        while not self.operation_stack.is_empty():
            # 计算值
            result = self._calculate_num_from_stack()
            # 将计算结果入数栈
            self.num_stack.push(result)
        result = self.num_stack.pop()
        print("表达式执行结果:" + expression + "=" + str(result))


def main():
    calculator = Calculator()

    calculator.calculate("3+5*3-6")
    calculator.calculate("30+5*3-6")
    calculator.calculate("130+5*3-6")


if __name__ == "__main__":
    main()
