package main

import (
    "errors"
    "fmt"
    "strconv"
)

// 队列
// 队列是一种特殊的线性表，特殊之处在于它只允许在表的前端进行删除操作，而在表的后端进行插入操作
// 所以队列又称为先进先出（FIFO—first in first out）

// 顺序队列
// 顺序队列必须分配一块连续的内存，并设置两个指针进行管理，一个是队头指针 front，它指向队头元素
// 另一个是队尾指针 rear，它指向下一个入队的位置。
// 随着插入和删除的进行，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元素，这时称为空队列
// 顺序队列不能有效的利用已经出队的元素占用的空间

// 数组实现顺序队列

type IntQueue struct {
    array   []int // 存放队列元素的切片（数组无法使用变量来定义长度）
    maxSize int   // 最大队列元素大小
    front   int   // 队头指针
    rear    int   // 队尾指针
}

func NewQueue(size int) *IntQueue {
    return &IntQueue{
        array:   make([]int, size),
        maxSize: size,
        front:   0,
        rear:    0,
    }
}

// 放入队列元素
func (q *IntQueue) Put(elem int) error {
    // 队尾指针不能超过最大队列元素大小
    if q.rear >= q.maxSize {
        return errors.New("queue is full")
    }
    q.array[q.rear] = elem
    q.rear++ // 队尾指针加一
    return nil
}

// 取出队列元素
func (q *IntQueue) Take() (int, error) {
    // 队头指针等于队尾指针表示队列为空
    if q.front == q.rear {
        return 0, errors.New("queue is empty")
    }
    elem := q.array[q.front]
    q.front++ // 队头指针加一
    return elem, nil
}

// 重新定义 String 方法，方便输出
func (q *IntQueue) String() string {
    str := "["
    for i := q.front; i < q.rear; i++ {
        str += strconv.Itoa(q.array[i]) + " "
    }
    str += "]"
    return str
}

func main() {
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
