"""
队列
队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
所以队列又称为先进先出（FIFO—first in first out）

循环队列
在实际使用队列时，顺序队列空间不能重复使用，需要对顺序队列进行改进。
不管是插入或删除，一旦 rear 指针增 1 或 front 指针增 1 时超出了队列所分配的空间，就
让它指向起始位置。当 MaxSize - 1增 1变到 0，可用取余运算获取队头或者队尾指针增加 1 后
的位置，队尾指针计算方法为 rear % MaxSize，队尾指针计算方法为 front % MaxSize。
这种循环使用队列空间的队列称为循环队列。

数组实现循环队列
"""

MIN_VALUE = float("-inf")


class IntQueue:
    def __init__(self, size):
        self.array = [None] * size  # 存放队列元素的数组
        self.max_size = size  # 最大队列元素大小
        self.front = 0  # 队头指针
        self.rear = 0  # 队尾指针

    def put(self, elem):
        """放入队列元素"""
        if self.is_full():
            print("queue is full")
            return False
        # 把元素放入队尾
        self.array[self.rear] = elem
        # 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
        self.rear = (self.rear + 1) % self.max_size
        return True

    def take(self):
        """取出队列元素"""
        if self.is_empty():
            print("queue is empty")
            return MIN_VALUE
        # 取出当前队头指向的元素
        elem = self.array[self.front]
        self.front = (self.front + 1) % self.max_size
        return elem

    def is_empty(self):
        # 队头指针等于队尾指针表示队列为空
        return self.front == self.rear

    def is_full(self):
        # 空出一个位置，判断是否等于队头指针
        # 队尾指针指向的位置不能存放队列元素，实际上会比 maxSize 指定的大小少一
        return (self.rear + 1) % self.max_size == self.front

    def size(self):
        return (self.rear + self.max_size - self.front) % self.max_size

    def show(self):
        result = "["
        temp_front = self.front
        for _ in range(self.size()):
            result += str(self.array[temp_front]) + " "
            temp_front = (temp_front + 1) % self.max_size
        result += "]"
        print(result)


def main():
    int_queue = IntQueue(5)
    int_queue.put(1)
    int_queue.put(2)
    int_queue.put(3)
    int_queue.put(4)
    int_queue.put(5)  # 队列已满，无法放入数据，实际上只能放 4 个元素

    print("intQueue: ")
    int_queue.show()

    num = int_queue.take()
    print("取出一个元素: " + str(num))
    num = int_queue.take()
    print("取出一个元素: " + str(num))
    num = int_queue.take()
    print("取出一个元素: " + str(num))
    num = int_queue.take()
    print("取出一个元素: " + str(num))
    num = int_queue.take()
    print("取出一个元素: " + str(num))
    if num == MIN_VALUE:
        print("出队失败!!!")

    # 取出数据后可以继续放入数据
    int_queue.put(5)
    print("intQueue: ")
    int_queue.show()


if __name__ == "__main__":
    main()
