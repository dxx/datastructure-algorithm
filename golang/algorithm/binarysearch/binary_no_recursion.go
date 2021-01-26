package main

import "fmt"

// 二分法查找(非递归)
func binarySearchNoRecursion(nums []int, findVal int) int {
    start := 0
    end := len(nums) - 1
    for start <= end {
        mid := (start + end) / 2
        if findVal < nums[mid] { // 查找的值在左边
            end = mid - 1
        } else if findVal > nums[mid] { // 查找的值在右边
            start = mid + 1
        } else {
            // 找到目标值的下标
            return mid
        }
    }
    return -1
}

func main() {
    value := 100
    nums := []int{1, 8, 10, 89, 100, 100, 123}
    index := binarySearchNoRecursion(nums, value)
    if index != -1 {
        fmt.Printf("找到 %d, 下标为 %d\n", value, index)
    } else {
        fmt.Printf("未找到 %d\n", value)
    }
}
