## 栈(stack)

> 各种语言实现代码：[Go](../code/golang/datastructure/stack)   [Java](../code/java/datastructure/src/com/dxx/stack)   [JavaScript](../code/javascript/datastructure/stack)   [TypeScript](../code/typescript/datastructure/stack)   [Python](../code/python/datastructure/stack)   [Rust](../code/rust/datastructure/src/stack)
>
> 默认使用 **Python** 语言实现。

### 简介

栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有先进后出（FILO—first in last out）的特点。

![data_structure_stack_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_stack_01.png)

### 实现

使用数组来实现栈。

* 定义数组用来存储栈元素
* 定义栈中元素最大大小 maxSize
* 定义栈顶，初始值为 -1
* 入栈方法 push
* 出栈方法 pop

定义栈类：

```python
class Stack:
    def __init__(self, size: int) -> None:
        self.array: list[object | None] = [None] * size  # 存放栈元素
        self.max_size = size  # 最大栈元素大小
        self.top = -1  # 栈顶
```

入栈方法：

```python
def push(self, elem: object) -> bool:
    # 判栈是否已满
    if self.top == self.max_size - 1:
        print("stack is full")
        return False
    # 栈顶加 1，将元素放入栈顶
    self.top += 1
    self.array[self.top] = elem
    return True
```

出栈方法：

```python
def pop(self) -> object:
    if self.top == -1:
        print("stack is empty")
        return ""
    # 取出栈顶元素，然后加 1
    elem = self.array[self.top]
    self.top -= 1
    return elem
```

为了方便查看输出结果，定义 `show` 方法：

```python
def show(self) -> None:
    text = "["
    for i in range(self.top, -1, -1):
        text += str(self.array[i]) + " "
    text += "]"
    print(text)
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_stack(self):
        # 创建一个栈
        stack = Stack(3)
        # 入栈
        stack.push("one")
        stack.push("two")
        stack.push("three")

        # 栈满，无法入栈
        is_success = stack.push("four")
        if not is_success:
            print("入栈失败!!!")
        stack.show()

        elem1 = stack.pop()
        elem2 = stack.pop()
        elem3 = stack.pop()

        print("出栈: " + str(elem1))
        print("出栈: " + str(elem2))
        print("出栈: " + str(elem3))

        elem = stack.pop()
        if elem is None or elem == "":
            print("出栈失败!!!")
        stack.show()
```

运行：

```shell
❯ python -m unittest test_stack.Test.test_stack
stack is full
入栈失败!!!
[three two one ]
出栈: three
出栈: two
出栈: one
stack is empty
出栈失败!!!
[]
```

### 综合计算器

使用栈实现一个加减乘除的计算器。假设一个字符串为 `3+5*3-6`，计算该表达式的值。

#### 思路分析

1. 定义两个栈，一个为数栈，一个为符号栈。
2. 截取字符串（需要考虑多位数），判断是否为数字，如果为数字，将字符串压入数栈，如果为运算符压入符号栈。
3. 压入符号栈前判断符号栈是否为空。
4. 如果符号栈为空，直接入栈。
5. 如果符号栈不为空，从符号栈中试探出一个符号，判断优先级。
6. 如果当前将要入栈的符号优先级小于或等于从符号栈中取出来的优先级，从数栈中弹出两个数，再从符号栈中弹出一个符号，进行运算，将运算结果压入数栈，再将要入栈的符号入栈。否则将当前符号直接入栈。
7. 不断的从操作符栈中取出一个符号，从数栈中取出两个数，进行计算，将计算结果压入数栈，当符号栈为空时跳出循环，此时数栈中的最后一个元素就是最终的计算结果。

#### 画图分析

![stack_calculator_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/stack_calculator_01.png)

![stack_calculator_02](https://dxx.github.io/static-resource/datastructure-algorithm/images/stack_calculator_02.png)

![stack_calculator_03](https://dxx.github.io/static-resource/datastructure-algorithm/images/stack_calculator_03.png)

![stack_calculator_04](https://dxx.github.io/static-resource/datastructure-algorithm/images/stack_calculator_04.png)

#### 代码实现

在之前实现的栈类中增加两个方法：

```python
def is_empty(self) -> bool:
    return self.top == -1


def peek(self) -> object | None:
    if self.is_empty():
        return None
    return self.array[self.top]
```

定义操作符类，定义相关操作符对应的优先级和计算方法：

```python
class Operation:
    def __init__(self, opt: str, priority: int, opt_fun) -> None:
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
```

定义计算器类和相关方法：

```python
class Calculator:
    def __init__(self) -> None:
        self.num_stack = Stack(10)
        self.operation_stack = Stack(10)

    def _is_operation(self, opt: str) -> bool:
        return opt in operations

    def _is_num(self, text: str) -> bool:
        return text.isdigit()

    def _priority(self, opt1: str, opt2: str) -> int | None:
        operation1 = operations.get(opt1)
        operation2 = operations.get(opt2)
        if not operation1 or not operation2:
            print("请检查运算符: " + opt1 + "," + opt2)
            return None
        return operation1.priority - operation2.priority

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

    def _calculate_num_from_stack(self) -> float:
        # 从数栈中弹出两个数，从符号栈中弹出一个符号
        num_str1 = self.num_stack.pop()
        num_str2 = self.num_stack.pop()
        opt = self.operation_stack.pop()
        # 计算值
        return self._calculate_num(float(str(num_str1)), float(str(num_str2)), str(opt))
```

计算表达式的核心方法如下：

```python
def calculate(self, expression: str) -> None:
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
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_calculate(self):
        calculator = Calculator()

        calculator.calculate("3+5*3-6")
        calculator.calculate("30+5*3-6")
        calculator.calculate("130+5*3-6")
```

运行：

```shell
❯ python -m unittest test_calculator.Test.test_calculate
表达式执行结果:3+5*3-6=12.0
表达式执行结果:30+5*3-6=39.0
表达式执行结果:130+5*3-6=139.0
```

### 逆波兰计算器

#### 简介

逆波兰计算器使用逆波兰表达式来计算表达式的值。逆波兰表达式也叫后缀表达式，后缀表达式指的是运算符写在操作数之后，比如 `12+`，它是计算机比较容易计算的一种表达式，因为计算机采用栈结构，执行先进后出的顺序。与之对应的有前缀表达式，中缀表达式，我们人一般识比较容易理解的是中缀表达式，比如 `3+5*3-6` 就属于中缀表达式。

#### 后缀表达式计算

假设一个后缀表达式为 `353*+2-`，计算出该后缀表达式的值。

思路分析：

1. 循环读取每个字符，判断是否是数字。
2. 如果是数字直接入栈。
3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中。

步骤：

1. 将 3，5，3 压入栈中。
2. 读取到 * 时，从栈中弹出两个数，栈顶弹出一个数 3，次栈顶弹出一个数 5。
3. 计算 3 * 5，结果等于 15，将结果压入栈中。
4. 读取到 + 时，从栈中弹出 15 和 3。
5. 计算 15 + 3，结果等于 18，将 18 压入栈中。
6. 将 2 压入栈中。
7. 读取到 - ，从栈中弹出 2 和 18。后一个数减去前一个数即 18 - 2。
8. 计算 18 - 2，结果等于 16，将结果压入栈中。
9. 循环结束后，栈中的 16 就是表达式的值。

代码实现

定义逆波兰计算器类，计算两个数，判断是否为数值等方法：

```python
class ReversePoland:
    def _is_num(self, text: str) -> bool:
        return text.isdigit()

    def _priority(self, opt1: str, opt2: str) -> int | None:
        operation1 = operations.get(opt1)
        operation2 = operations.get(opt2)
        if not operation1 or not operation2:
            print("请检查运算符: " + opt1 + "," + opt2)
            return None
        return operation1.priority - operation2.priority

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
```

计算后缀表达式的方法：

```python
def cal_suffix_expression(self, expr: list[str] | None) -> float | None:
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
```

测试代码如下：

```python
class Test(unittest.TestCase):
    
    def test_reverse_poland(self):
        reverse_poland = ReversePoland()
        expr = "3 5 3 * + 2 -"
        # 假设数和数或符号之间有空格
        expressions = expr.split(" ")
        result = reverse_poland.cal_suffix_expression(expressions)
        print("后缀表达式 " + expr + " 的计算结果为: " + str(result))
```

运行：

```shell
❯ python -m unittest test_reverse_poland.Test.test_reverse_poland
后缀表达式 3 5 3 * + 2 - 的计算结果为: 16.0
```

#### 中缀转后缀表达式

将中缀表达式转换成后缀表达式，步骤如下：

1. 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2。
2. 从左至右扫描中缀表达式。
3. 遇到数字时，将其压入 stack2。
4. 如果是 "(" 号直接压入 stack1。
5. 如果是 ")" 号，依次弹出 stack1 中栈顶的元素，并压入 stack2 中，直到遇到 "(" 将这一对括号丢弃。
6. 遇到运算符时，比较其与 stack1 栈顶运算符的优先级
   * 如果 stack1 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈。
   * 如果优先级比栈顶运算符低或者相等，将 stack1 栈顶的运算符弹出并压入到 stack2 中，再次转到 6-1 步，与 stack1 中新的栈顶运算符相比较，最后将当前运算符压入 stack1。

7. 重复 2 - 6，直到表达式末尾。
8. 将 stack1 中剩余的运算符依次弹出并压入 stack2 中。
9. 依次弹出 stack2 中的元素，将结果逆序就是转换后的后缀表达式。

**代码实现**

定义一个将字符串表达式转换成数组的函数：

```python
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
```

接下来就是关键的中缀表达式转后缀表达式函数：

```python
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
```

测试代码如下：

```python
class Test(unittest.TestCase):
    
    def test_reverse_poland(self):
        reverse_poland = ReversePoland()
        expr = "1+((2+3)*4)-5"
        expressions = reverse_poland.expr_to_array(expr)
        print("将中缀表达式放入数组, 结果为: " + str(expressions))

        expressions = reverse_poland.infix_to_suffix(expressions)
        print("中缀表达式转换成后缀表达式, 结果为: " + str(expressions))

        result = reverse_poland.cal_suffix_expression(expressions)
        print("计算表达式 " + expr + ", 结果为: " + str(result))
```

运行：

```shell
❯ python -m unittest test_reverse_poland.Test.test_reverse_poland
将中缀表达式放入数组, 结果为: ['1', '+', '(', '(', '2', '+', '3', ')', '*', '4', ')', '-', '5']
中缀表达式转换成后缀表达式, 结果为: ['1', '2', '3', '+', '4', '*', '+', '5', '-']
计算表达式 1+((2+3)*4)-5, 结果为: 16.0
```

> 注意：这里只能计算整数，且表达式前后不能有空格。
