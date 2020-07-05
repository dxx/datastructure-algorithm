package main

import (
    "datastructure/stack/array"
    "fmt"
    "regexp"
    "strconv"
)

type Operation struct {
    operation string
    priority  int
    optFunc   func(num1, num2 int) int
}

// 定义相关操作符对应的优先级和计算方法
var operations = map[string]Operation{
    "+": {"+", 1, func(num1, num2 int) int {
        return num1 + num2
    }},
    "-": {"-", 1, func(num1, num2 int) int {
        return num1 - num2
    }},
    "*": {"*", 2, func(num1, num2 int) int {
        return num1 * num2
    }},
    "/": {"/", 2, func(num1, num2 int) int {
        return num1 / num2
    }},
}

// 计算器结构体
type Calculator struct {
    numStack       *array.Stack // 数栈
    operationStack *array.Stack // 符号栈
}

func NewCalculator() *Calculator {
    numStack := array.NewStack(10)
    operationStack := array.NewStack(10)
    return &Calculator{numStack: numStack, operationStack: operationStack}
}

// 判断是否是操作符号
func (cal *Calculator) isOperation(opt string) bool {
    _, ok := operations[opt]
    return ok
}

// 判断是否为数字
func (cal *Calculator) isNum(char string) bool {
    matched, _ := regexp.MatchString("\\d+", char)
    return matched
}

// 计算操作符的优先级
func (cal *Calculator) priority(opt1, opt2 string) int {
    operation1, ok1 := operations[opt1]
    operation2, ok2 := operations[opt2]
    if ok1 && ok2 {
        return operation1.priority - operation2.priority
    } else {
        panic(fmt.Sprintf("请检查运算符: %s, %s\n", opt1, opt2))
    }
}

// 计算结果
func (cal *Calculator) calculateNum(num1, num2 int, opt string) int {
    optFunc := operations[opt].optFunc
    if opt == "-" || opt == "/" {
        // 因为出栈后两数的位置颠倒，需交换两个数的位置
        num1, num2 = num2, num1
    }
    return optFunc(num1, num2)
}

func (cal *Calculator) calculateNumFromStack() int {
    // 从数栈中弹出两个数，从符号栈中弹出一个符号
    numStr1, _ := cal.numStack.Pop()
    numStr2, _ := cal.numStack.Pop()
    opt, _ := cal.operationStack.Pop()
    num1, _ := strconv.Atoi(numStr1)
    num2, _ := strconv.Atoi(numStr2)
    // 计算值
    return cal.calculateNum(num1, num2, opt)
}

// 计算表达式的值
func (cal *Calculator) Calculate(expression string) {
    if expression == "" {
        return
    }
    var index int
    var number string
    for index < len(expression) {
        char := expression[index : index+1]
        // 判断是否为符号
        if cal.isOperation(char) {
            // 判断符号栈是否为空
            if cal.operationStack.IsEmpty() {
                // 压入符号栈
                _ = cal.operationStack.Push(char)
            } else {
                // 符号栈不为空，判断优先级
                opt := cal.operationStack.Peek()
                // char 优先级小于等于 elem
                if cal.priority(char, opt) <= 0 {
                    // 计算值
                    result := cal.calculateNumFromStack()
                    // 将计算结果入数栈
                    _ = cal.numStack.Push(strconv.Itoa(result))
                }
                // 将当前操作符入符号栈
                _ = cal.operationStack.Push(char)
            }
        } else if cal.isNum(char) {
            // 向后面再取一位判断是否为数字
            if index < len(expression)-1 && cal.isNum(expression[index+1:index+2]) {
                number += char
                index++
                continue
            }
            // 压入数栈
            _ = cal.numStack.Push(number + char)
            number = ""
        } else {
            panic("无法识别的字符:" + char)
        }

        index++
    }

    // 全部数和符号都压入对应的栈后，取出计算
    // 符号栈为空，跳出循环
    for !cal.operationStack.IsEmpty() {
        // 计算值
        result := cal.calculateNumFromStack()
        // 将计算结果入数栈
        _ = cal.numStack.Push(strconv.Itoa(result))
    }
    // 弹出最终结果
    result, _ := cal.numStack.Pop()
    fmt.Printf("表达式执行结果: %s=%s\n", expression, result)
}

func main() {
    calculator := NewCalculator()

    calculator.Calculate("3+5*3-6")
    calculator.Calculate("30+5*3-6")
    calculator.Calculate("130+5*3-6")
}
