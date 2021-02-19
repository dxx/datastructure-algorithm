package search

import "fmt"

// 插值查找
// 基于二分法查找，步骤和二分法查找一样
// 中间下标计算公式: mid = start + (end - start) * (value - array[start]) / (array[end] - array[start])
func insertValSearch(nums[] int, start, end int, findVal int) int {
    if start > end {
        return -1
    }
    // 根据 findVal 自适应计算中间下标
    mid := start + (end - start) * (findVal - nums[start]) / (nums[end] - nums[start])
    fmt.Printf("mid: %d\n", mid)
    if findVal < nums[mid] {
        // 向左递归
        return insertValSearch(nums, start, mid - 1, findVal)
    } else if findVal > nums[mid] {
        // 向右递归
        return insertValSearch(nums, mid + 1, end, findVal)
    } else {
        // 查找值和中间值相等，返回下标
        return mid
    }
}
