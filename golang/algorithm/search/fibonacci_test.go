package search

import (
    "fmt"
    "testing"
)

func TestFibonacciSearch(t *testing.T) {
    value := 100
    nums := []int{1, 8, 10, 89, 100, 100, 123}
    index := fibonacciSearch(nums, value)
    if index != -1 {
        fmt.Printf("找到 %d, 下标为 %d\n", value, index)
    } else {
        fmt.Printf("未找到 %d\n", value)
    }
}
