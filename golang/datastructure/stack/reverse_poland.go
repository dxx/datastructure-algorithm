package main

import (
    "datastructure/stack/array"
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type Opt struct {
    operation string
    priority  int
    optFunc   func(int, int) int
}

// 定义相关操作符对应的优先级和计算方法
var opts = map[string]Opt{
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

// 计算结果
func calculateNum(num1, num2 int, opt string) int {
    optFunc := opts[opt].optFunc
    if opt == "-" || opt == "/" {
        // 因为出栈后两数的位置颠倒，需交换两个数的位置
        num1, num2 = num2, num1
    }
    return optFunc(num1, num2)
}

func isNum(s string) bool {
    matched, _ := regexp.MatchString("\\d+", s)
    return matched
}

// 后缀表达式计算
// 1. 循环读取每个字符，判断是否是数字
// 2. 如果是数字直接入栈
// 3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中
func calSuffixExpression(expr []string) int {
    stack := array.NewStack(len(expr))
    for _, str := range expr {
        if isNum(str) {
            _ = stack.Push(str)
            continue
        }
        _, ok := opts[str]
        if !ok {
            panic("无效的运算符: " + str)
        }
        // 计算
        numStr1, _ := stack.Pop()
        numStr2, _ := stack.Pop()
        num1, _ := strconv.Atoi(numStr1)
        num2, _ := strconv.Atoi(numStr2)
        result := calculateNum(num1, num2, str)
        // 入栈
        _ = stack.Push(strconv.Itoa(result))
    }
    elem, _ := stack.Pop()
    result, _ := strconv.Atoi(elem)
    return result
}

// 将表达式转换成切片
func exprToSlice(expr string) []string {
    var expressions []string
    for i := 0; i < len(expr); i++ {
        char := expr[i : i+1]
        if isNum(char) {
            // 向后面继续判断是否为数字
            for i + 1 < len(expr) && isNum(expr[i+1:i+2]) {
                char += expr[i+1:i+2]
                i++
            }
        }
        expressions = append(expressions, char)
    }
    return expressions
}

// 判断两个操作符的优先级
func priority(opt1, opt2 string) int {
    operation1, ok1 := opts[opt1]
    operation2, ok2 := opts[opt2]
    if ok1 && ok2 {
        return operation1.priority - operation2.priority
    } else {
        panic(fmt.Sprintf("请检查运算符: %s, %s\n", opt1, opt2))
    }
}

// 中缀表达式转后缀表达式
func infixToSuffix(infix []string) []string {
    // 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2
    stack := array.NewStack(len(infix))
    var suffixes []string // 由于中间结果栈不需要弹出元素，可以使用数组来保存
    // 循环表达式
    for _, str := range infix {
        // 遇到数字时，将其放入 suffixes
        if isNum(str) {
            suffixes = append(suffixes, str)
            continue
        }
        if str == "(" {
            // 如果是 ( 直接入栈
            _ = stack.Push(str)
            continue
        }
        if str == ")" {
            for stack.Peek() != "(" {
                // 弹出 stack 中栈顶的元素，并追加到 suffixes
                elem, _ := stack.Pop()
                suffixes = append(suffixes, elem)
            }
            // 弹出 (，消除一对 ( )
            _, _ = stack.Pop()
            continue
        }
        // 如果是运算符
        opt, ok := opts[str]
        if ok {
            if stack.IsEmpty() || stack.Peek() == "(" {
                // 如果 stack 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈
                _ = stack.Push(str)
                continue
            }
            // 栈不为空，并且当前字符串的优先级小于等于栈顶的元素
            for !stack.IsEmpty() && priority(opt.operation, stack.Peek()) <= 0 {
                elem, _ := stack.Pop()
                // 将栈顶的元素追加到 suffixes
                suffixes = append(suffixes, elem)
            }
            // 直接入栈
            _ = stack.Push(str)
        } else {
            panic("无法识别的字符: " + str)
        }
    }
    for !stack.IsEmpty() {
        // 将 stack 中剩余的运算符依次追加到 suffixes
        elem, _ := stack.Pop()
        suffixes = append(suffixes, elem)
    }
    // 因为这里用的是数组，它里面元素的顺序就是栈元素出栈后逆序排列的顺序
    return suffixes
}

func main() {
    expr := "3 5 3 * + 2 -"
    // 假设数和数或符号之间有空格
    expressions := strings.Split(expr, " ")
    result := calSuffixExpression(expressions)
    fmt.Printf("后缀表达式 %s 的计算结果为: %d\n", expr, result)

    expr = "1+((2+3)*4)-5"
    expressions = exprToSlice(expr)
    fmt.Printf("将中缀表达式放入切片, 结果为: %v\n", expressions)

    expressions = infixToSuffix(expressions)
    fmt.Printf("中缀表达式转换成后缀表达式, 结果为: %v\n", expressions)

    result = calSuffixExpression(expressions)
    fmt.Printf("计算表达式 %s, 结果为: %v\n", expr, result)
}
