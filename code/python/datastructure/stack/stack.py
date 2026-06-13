"""
栈
栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。
在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入
新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个
新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有
先进后出（FILO—first in last out）的特点。
"""


class Stack:
    def __init__(self, size: int) -> None:
        self.array: list[object | None] = [None] * size  # 存放栈元素
        self.max_size = size  # 最大栈元素大小
        self.top = -1  # 栈顶

    """
    入栈
    """
    def push(self, elem: object) -> bool:
        # 判栈是否已满
        if self.top == self.max_size - 1:
            print("stack is full")
            return False
        # 栈顶加 1，将元素放入栈顶
        self.top += 1
        self.array[self.top] = elem
        return True

    """
    出栈
    """
    def pop(self) -> object:
        if self.top == -1:
            print("stack is empty")
            return ""
        # 取出栈顶元素，然后加 1
        elem = self.array[self.top]
        self.top -= 1
        return elem

    """
    判断栈是否为空
    """
    def is_empty(self) -> bool:
        return self.top == -1

    """
    窥视栈顶元素
    """
    def peek(self) -> object | None:
        if self.is_empty():
            return None
        return self.array[self.top]

    def show(self) -> None:
        text = "["
        for i in range(self.top, -1, -1):
            text += str(self.array[i]) + " "
        text += "]"
        print(text)
