## 队列(queue)

> 各种语言实现代码：[Go](../code/golang/datastructure/queue)   [Java](../code/java/datastructure/src/com/dxx/queue)   [JavaScript](../code/javascript/datastructure/queue)   [TypeScript](../code/typescript/datastructure/queue)   [Python](../code/python/datastructure/queue)   [Rust](../code/rust/datastructure/src/queue)
>
> 默认使用 **Python** 语言实现。

### 简介

队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作，所以队列又称为先进先出（FIFO—first in first out）的线性表。进行插入操作的一端叫做队尾，进行删除操作的一端叫做队头。队列的数据元素叫做队列元素，在队列中插入一个队列元素称为入队，从队列中删除一个队列元素称为出队。

![data_structure_queue_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_queue_01.png)

### 顺序队列

顺序队列类似数组，它需要一块连续的内存，并有两个指针，一个是队头指针 front，它指向队头元素，另一个是队尾指针 rear，它指向下一个入队的位置。在队尾插入一个元素时，队尾指针加一，在队头删除一个元素时，队头指针加一。不断的进行插入和删除操作，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元素，没有元素的队列称为空队列。对于已经出队列的元素所占用的空间，顺序队列无法再次利用。

队列可以使用数组结构或者链表结构来实现，这里使用数组来实现队列。

用数组来实现顺序队列的思路：

* 定义数组，存储队列元素
* 定义队列最大大小 maxSize
* 定义队头指针 front，初始化为 0
* 队尾指针 rear，初始化为 0
* 入队方法 put
* 出队方法 take

定义顺序队列类：

```python
MIN_VALUE = float("-inf")


class IntQueue:
    def __init__(self, size: int) -> None:
        self.array: list[int | None] = [None] * size  # 存放队列元素的数组
        self.max_size = size  # 最大队列元素大小
        self.front = 0  # 队头指针
        self.rear = 0  # 队尾指针
```

入队方法：

```python
def put(self, elem: int) -> bool:
    """放入队列元素"""
    # 队尾指针不能超过最大队列元素大小
    if self.rear >= self.max_size:
        print("queue is full")
        return False
    # 把元素放入队尾，然后队尾指针加一
    self.array[self.rear] = elem
    self.rear += 1
    return True
```

出队方法：

```python
def take(self) -> int | float | None:
    """取出队列元素"""
    # 队头指针等于队尾指针表示队列为空
    if self.front == self.rear:
        print("queue is empty")
        return MIN_VALUE
    # 取出当前队头指向的元素，然后队头指针加一
    elem = self.array[self.front]
    self.front += 1
    return elem
```

为了方便查看输出结果，定义 `show` 方法：

```python
def show(self) -> None:
    result = "["
    for i in range(self.front, self.rear):
        result += str(self.array[i]) + " "
    result += "]"
    print(result)
```

测试代码如下：

```python
class Test(unittest.TestCase):
    
    def test_int_queue(self):
        int_queue = IntQueue(3)
        int_queue.put(1)
        int_queue.put(2)
        int_queue.put(3)
        int_queue.put(4)  # 队列已满，无法放入数据

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
        if num == MIN_VALUE:
            print("出队失败!!!")

        # 此时队列已经用完，无法放数据
        is_success = int_queue.put(4)
        if not is_success:
            print("入队失败!!!")
        print("intQueue: ", end="")
        int_queue.show()
```

测试以上代码，运行：

```shell
❯ python -m unittest test_main.Test.test_int_queue
queue is full
intQueue: [1 2 3 ]
取出一个元素: 1
取出一个元素: 2
取出一个元素: 3
queue is empty
取出一个元素: -inf
出队失败!!!
queue is full
入队失败!!!
intQueue: []
```

### 循环队列

在实际使用队列时，顺序队列空间不能重复使用，需要对顺序队列进行改进。不管是插入或删除，一旦 rear 指针增 1 或 front 指针增 1 时超出了队列所分配的空间，就让它指向起始位置。当 MaxSize - 1增 1变到 0，可用取余运算获取队头或者队尾指针增加 1 后的位置，队尾指针计算方法为 rear % MaxSize，队尾指针计算方法为 front % MaxSize。这种循环使用队列空间的队列称为循环队列。除了一些简单应用之外，真正实用的队列是循环队列。

使用数组实现循环队列思路：

* 定义数组，存储队列元素
* 定义队列最大大小 maxSize
* 定义队头指针 front，初始化为 0，删除元素后，重新计算值，计算公式为：(front + 1) % maxSize
* 队尾指针 rear，初始化为 0，插入元素后，重新计算值，计算公式为：(q.rear + 1) % q.maxSize
* 判断队列是否为空的方法，判断 front 等于 rear 判断即可
* 判断队列是否已满的方法，判断 (rear + 1) % maxSize 是否等于 front
* 获取队列元素大小方法，计算公式：(rear + maxSize - front) % maxSize
* 入队方法 put
* 出队方法 take

定义循环队列类：

```python
MIN_VALUE = float("-inf")


class IntQueue:
    def __init__(self, size: int) -> None:
        self.array: list[int | None] = [None] * size  # 存放队列元素的数组
        self.max_size = size  # 最大队列元素大小
        self.front = 0  # 队头指针
        self.rear = 0  # 队尾指针
```

判断队列是否为空：

```python
def is_empty(self) -> bool:
    # 队头指针等于队尾指针表示队列为空
    return self.front == self.rear
```

判断队列是否已满：

```python
def is_full(self) -> bool:
    # 空出一个位置，判断是否等于队头指针
    # 队尾指针指向的位置不能存放队列元素，实际上会比 maxSize 指定的大小少一
    return (self.rear + 1) % self.max_size == self.front
```

获取队列元素大小：

```python
def size(self) -> int:
    return (self.rear + self.max_size - self.front) % self.max_size
```

入队方法：

```python
def put(self, elem: int) -> bool:
    """放入队列元素"""
    if self.is_full():
        print("queue is full")
        return False
    # 把元素放入队尾
    self.array[self.rear] = elem
    # 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
    self.rear = (self.rear + 1) % self.max_size
    return True
```

出队方法：

```python
def take(self) -> int | float | None:
    """取出队列元素"""
    if self.is_empty():
        print("queue is empty")
        return MIN_VALUE
    # 取出当前队头指向的元素
    elem = self.array[self.front]
    self.front = (self.front + 1) % self.max_size
    return elem
```

定义 `show` 方法：

```python
def show(self) -> None:
    result = "["
    temp_front = self.front
    for _ in range(self.size()):
        result += str(self.array[temp_front]) + " "
        temp_front = (temp_front + 1) % self.max_size
    result += "]"
    print(result)
```

测试代码如下：

```python
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
```

测试以上代码，运行：

```shell
❯ python -m unittest test_main.Test.test_int_queue
queue is full
intQueue: [1 2 3 4 ]
取出一个元素: 1
取出一个元素: 2
取出一个元素: 3
取出一个元素: 4
queue is empty
取出一个元素: -inf
出队失败!!!
intQueue: [5 ]
```
