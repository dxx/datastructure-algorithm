import unittest
from main import IntQueue, MIN_VALUE


class Test(unittest.TestCase):
    
    def test_int_queue(self):
        int_queue = IntQueue(5)
        int_queue.put(1)
        int_queue.put(2)
        int_queue.put(3)
        int_queue.put(4)
        int_queue.put(5)  # 队列已满，无法放入数据，实际上只能放 4 个元素

        print("intQueue: ", end="")
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
        print("intQueue: ", end="")
        int_queue.show()
