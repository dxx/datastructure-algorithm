"""
队列
队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
所以队列又称为先进先出（FIFO—first in first out）

顺序队列
顺序队列类似数组，它需要一块连续的内存，并有两个指针，一个是队头指针 front，它指向队头元素
另一个是队尾指针 rear，它指向下一个入队的位置。
不断的进行插入和删除操作，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元
素，没有元素的队列称为空队列。对于已经出队列的元素所占用的空间，顺序队列无法再次利用。

数组实现顺序队列
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
        # 队尾指针不能超过最大队列元素大小
        if self.rear >= self.max_size:
            print("queue is full")
            return False
        # 把元素放入队尾，然后队尾指针加一
        self.array[self.rear] = elem
        self.rear += 1
        return True

    def take(self):
        """取出队列元素"""
        # 队头指针等于队尾指针表示队列为空
        if self.front == self.rear:
            print("queue is empty")
            return MIN_VALUE
        # 取出当前队头指向的元素，然后队头指针加一
        elem = self.array[self.front]
        self.front += 1
        return elem

    def show(self):
        result = "["
        for i in range(self.front, self.rear):
            result += str(self.array[i]) + " "
        result += "]"
        print(result)


def main():
    int_queue = IntQueue(3)
    int_queue.put(1)
    int_queue.put(2)
    int_queue.put(3)
    int_queue.put(4)  # 队列已满，无法放入数据

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
    if num == MIN_VALUE:
        print("出队失败!!!")

    # 此时队列已经用完，无法放数据
    is_success = int_queue.put(4)
    if not is_success:
        print("入队失败!!!")
    print("intQueue: ")
    int_queue.show()


if __name__ == "__main__":
    main()
