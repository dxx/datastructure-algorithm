package main

import "fmt"

// 线性查找
func sequenceSearch(nums[] int, num int) int {
    if nums == nil {
        return -1
    }
    for i, n := range nums {
        if n == num {
            return i
        }
    }
    return -1
}

func main() {
    value := 8
    nums := []int{2, 5, 1, 7, 8, 16}
    index := sequenceSearch(nums, value)
    if index != -1 {
        fmt.Printf("%d 在 nums 中的下标为: %d", value, index)
    }
}
