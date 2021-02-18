package sequence

import (
    "errors"
    "strconv"
)

// 队列
// 队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
// 所以队列又称为先进先出（FIFO—first in first out）

// 顺序队列
// 顺序队列类似数组，它需要一块连续的内存，并有两个指针，一个是队头指针 front，它指向队头元素
// 另一个是队尾指针 rear，它指向下一个入队的位置。
// 不断的进行插入和删除操作，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元
// 素，没有元素的队列称为空队列。对于已经出队列的元素所占用的空间，顺序队列无法再次利用。

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
