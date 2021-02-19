package search

import (
    "fmt"
    "testing"
)

func TestInsertValSearch(t *testing.T) {
    var nums []int
    // 填充 1 - 100
    for i := 1; i <= 100; i++ {
        nums = append(nums, i)
    }
    value := 58
    index := insertValSearch(nums, 0, len(nums) - 1, value)
    if index != -1 {
        fmt.Printf("找到 %d, 下标为 %d\n", value, index)
    } else {
        fmt.Printf("未找到 %d\n", value)
    }
}

