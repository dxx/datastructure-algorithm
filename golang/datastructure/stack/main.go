package main

import (
    "datastructure/stack/array"
    "fmt"
)

func main() {
    // 创建一个栈
    stack := array.NewStack(3)
    // 入栈
    _ = stack.Push("one")
    _ = stack.Push("two")
    _ = stack.Push("three")

    // 栈满，如法入栈
    err := stack.Push("four")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(stack)

    elem1, _ := stack.Pop()
    elem2, _ := stack.Pop()
    elem3, _ := stack.Pop()

    fmt.Println("出栈:", elem1)
    fmt.Println("出栈:", elem2)
    fmt.Println("出栈:", elem3)

    // 栈空无法出栈
    _, err = stack.Pop()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(stack)
}
