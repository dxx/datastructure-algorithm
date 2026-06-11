import unittest
from main import IntQueue, MIN_VALUE


class Test(unittest.TestCase):
    
    def test_int_queue(self):
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
