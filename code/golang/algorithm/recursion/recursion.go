package recursion

// 递归
// 递归是指在程序运行过程中调用本身的编程技巧

// 计算阶乘
func factorial(n int) int {
    if n > 0 {
        return n * factorial(n - 1)
    }
    return 1
}
