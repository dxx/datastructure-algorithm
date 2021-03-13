## 队列(queue)

> 各种语言实现代码：[Go](./golang/datastructure/queue)   [Java](./java/datastructure/src/com/dxx/queue)   [JavaScript](./javascript/datastructure/queue)   [Rust](./rust/datastructure/src/queue)
>
> 默认使用 **Go** 语言实现。

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

定义顺序队列结构体和创建结构体实例函数：

```go
type IntQueue struct {
    array []int // 存放队列元素的切片（数组无法使用变量来定义长度）
    maxSize int // 最大队列元素大小
    front int // 队头指针
    rear int // 队尾指针
}

func NewQueue(size int) *IntQueue {
    return &IntQueue{
        array:   make([]int, size),
        maxSize: size,
        front:   0,
        rear:    0,
    }
}
```

入队方法：

```go
func (q *IntQueue) Put(elem int) error {
    // 队尾指针不能超过最大队列元素大小
    if q.rear >= q.maxSize {
        return errors.New("queue is full")
    }
    q.array[q.rear] = elem
    q.rear++ // 队尾指针加一
    return nil
}
```

出队方法：

```go
func (q *IntQueue) Take() (int, error) {
    // 队头指针等于队尾指针表示队列为空
    if q.front == q.rear {
        return 0, errors.New("queue is empty")
    }
    elem := q.array[q.front]
    q.front++ // 队头指针加一
    return elem, nil
}
```

为了方便查看输出结果，重新定义 String 方法：

```go
// 重新定义 String 方法，方便输出
func (q *IntQueue) String() string {
    str := "["
    for i := q.front; i < q.rear; i++ {
        str += strconv.Itoa(q.array[i]) + " "
    }
    str += "]"
    return str
}
```

测试代码如下：

```go
func TestSequenceQueue(t *testing.T) {
    intQueue := NewQueue(3)
    _ = intQueue.Put(1)
    _ = intQueue.Put(2)
    _ = intQueue.Put(3)
    _ = intQueue.Put(4) // 队列已满，无法放入数据

    fmt.Println("intQueue:", intQueue)

    num, _ := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, takeErr := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    if takeErr != nil {
        fmt.Println("出队失败:", takeErr)
    }

    // 此时队列已经用完，无法放数据
    putErr := intQueue.Put(4)
    if putErr != nil {
        fmt.Println("入队失败:", putErr)
    }
    fmt.Println("intQueue:", intQueue)
}
```

测试以上代码，运行：

```shell
golang/datastructure>go test -v -run ^TestSequenceQueue$ ./queue/sequence
=== RUN   TestSequenceQueue
intQueue: [1 2 3 ]
取出一个元素: 1
取出一个元素: 2
取出一个元素: 3
取出一个元素: 0
出队失败: queue is empty
入队失败: queue is full
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

定义循环队列结构体和创建结构体实例函数：

```go
type IntQueue struct {
    array []int // 存放队列元素的切片（数组无法使用变量来定义长度）
    maxSize int // 最大队列元素大小
    front int // 队头指针
    rear int // 队尾指针
}

func NewQueue(size int) *IntQueue {
    return &IntQueue{
        array:   make([]int, size),
        maxSize: size,
        front:   0,
        rear:    0,
    }
}
```

判断队列是否为空：

```go
func (q *IntQueue) isEmpty() bool {
    // 队头指针等于队尾指针表示队列为空
    return q.front == q.rear
}
```

判断队列是否已满：

```go
func (q *IntQueue) isFull() bool {
    // 空出一个位置，判断是否等于队头指针
    // 队尾指针指向的位置不能存放队列元素，实际上会比 maxSize 指定的大小少一
    return (q.rear + 1) % q.maxSize == q.front
}
```

获取队列元素大小：

```go
func (q *IntQueue) size() int {
    return (q.rear + q.maxSize - q.front) % q.maxSize
}
```

入队方法：

```go
func (q *IntQueue) Put(elem int) error {
    if q.isFull() {
        return errors.New("queue is full")
    }
    q.array[q.rear] = elem
    // 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
    q.rear = (q.rear + 1) % q.maxSize
    return nil
}
```

出队方法：

```go
func (q *IntQueue) Take() (int, error) {
    if q.isEmpty() {
        return 0, errors.New("queue is empty")
    }
    elem := q.array[q.front]
    q.front = (q.front + 1) % q.maxSize
    return elem, nil
}
```

重新定义的 String 方法：

```go
func (q *IntQueue) String() string {
    str := "["
    tempFront := q.front
    for i := 0; i < q.size(); i++ {
        str += strconv.Itoa(q.array[tempFront]) + " "
        // 超过最大大小，从 0 开始
        tempFront = (tempFront + 1 ) % q.maxSize
    }
    str += "]"
    return str
}
```

测试代码如下：

```go
func TestRoundQueue(t *testing.T) {
    intQueue := NewQueue(5)
    _ = intQueue.Put(1)
    _ = intQueue.Put(2)
    _ = intQueue.Put(3)
    _ = intQueue.Put(4)
    _ = intQueue.Put(5) // 队列已满，无法放入数据，实际上只能放 4 个元素

    fmt.Println("intQueue:", intQueue)

    num, _ := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, takeErr := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    if takeErr != nil {
        fmt.Println("出队失败:", takeErr)
    }

    // 取出数据后可以继续放入数据
    _ = intQueue.Put(5)
    fmt.Println("intQueue:", intQueue)
}
```

测试以上代码，运行：

```shell
golang/datastructure>go test -v -run ^TestRoundQueue$ ./queue/round
=== RUN   TestRoundQueue
intQueue: [1 2 3 4 ]
取出一个元素: 1
取出一个元素: 2
取出一个元素: 3
取出一个元素: 4
取出一个元素: 0
出队失败: queue is empty
intQueue: [5 ]
```

