package sort

import (
    "fmt"
)

// 冒泡排序
// 1.从当前元素起，向前依次比较每一对相邻元素，若逆序则交换
// 2.对所有元素均重复以上步骤，直至最后一个元素
func bubbleSort(nums []int) {
    if nums == nil {
        return
    }
    length := len(nums)
    // 外循环为排序趟数，length 个数进行 length - 1 趟
    for i := 0; i < length - 1; i++ {
        // 内循环为每趟比较的次数，第 i 趟比较 length - i 次
        for j := length - 1; j > i; j-- {
            // 相邻元素比较比较大小，然后交换位置
            if nums[j] < nums[j - 1] {
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
            }
        }
        fmt.Printf("第 %d 趟排序结果:%v\n", i + 1, nums)
    }
}

// 优化
// 在某次循环中，如果发现没有发生交换，则终止循环
func optimizeBubbleSort(nums []int) {
    if nums == nil {
        return
    }
    length := len(nums)
    isChange := false // 标记是否发生交换
    for i := 0; i < length - 1; i++ {
        for j := length - 1; j > i; j-- {
            if nums[j] < nums[j - 1] {
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
                isChange = true // 发生交换
            }
        }
        fmt.Printf("第 %d 趟排序结果:%v\n", i + 1, nums)
        if !isChange {
            break // 跳出循环，终止比较
        } else {
            isChange = false // 重置
        }
    }
}
