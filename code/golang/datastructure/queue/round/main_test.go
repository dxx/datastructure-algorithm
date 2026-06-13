package round

import (
    "fmt"
    "testing"
)

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
