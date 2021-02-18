package sequence

import (
    "fmt"
    "testing"
)

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
