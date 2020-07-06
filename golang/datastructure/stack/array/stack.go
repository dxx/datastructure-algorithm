package array

import (
    "errors"
)

// 栈
// 栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。
// 在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入
// 新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个
// 新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有
// 先进后出（FILO—first in last out）的特点。

type Stack struct {
    array   []string // 存放栈元素的切片（数组无法使用变量来定义长度）
    maxSize int      // 最大栈元素大小
    top     int      // 栈顶
}

func NewStack(size int) *Stack {
    return &Stack{
        array:   make([]string, size),
        maxSize: size,
        top:     -1, // 初始化为 -1
    }
}

// 入栈
func (s *Stack) Push(elem string) error {
    // 判断栈是否已满
    if s.top == s.maxSize-1 {
        return errors.New("stack is full")
    }
    s.top++ // 栈顶加 1
    s.array[s.top] = elem
    return nil
}

// 出栈
func (s *Stack) Pop() (string, error) {
    if s.top == -1 {
        return "", errors.New("stack is empty")
    }
    elem := s.array[s.top]
    s.top-- // 栈顶减 1
    return elem, nil
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
    return s.top == -1
}

// 窥视栈顶元素
func (s *Stack) Peek() string {
    if s.IsEmpty() {
        return ""
    }
    return s.array[s.top]
}

// 重新定义 String 方法，方便输出
func (s *Stack) String() string {
    str := "["
    for i := s.top; i >= 0; i-- {
        str += s.array[i] + " "
    }
    str += "]"
    return str
}
