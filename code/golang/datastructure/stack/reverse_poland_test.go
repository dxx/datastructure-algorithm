package stack

import (
    "fmt"
    "strings"
    "testing"
)

func TestReversePoland(t *testing.T) {
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
