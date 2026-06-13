package sort

import (
    "strconv"
)

// 基数排序
// 1.创建一个二维数组，数组长度为 10，并初始化 10 个一维数组，每个一维数组的长度为待排序数组的长度
// 2.遍历待排序的数组，从最低位开始，求出每个元素的个位数作为二维数组的下标，将其放入到二维数组对应的数组中
// 3.从二维数组中依次取出所有元素放入原数组中
// 4.重复步骤 2，依次计算个位、十位、百位等，作为下标，直到待排序数组中的最大位数
func radixSort(nums[] int) {
    if nums == nil {
        return
    }

    max := nums[0] // 最大位的元素
    for _, num := range nums {
        if max < num {
            max = num
        }
    }
    maxLength := len(strconv.Itoa(max))

    // var bucket [10][len(nums)]int
    var bucket [10][]int // 桶数组，二维数组中的元素使用切片替代
    for i := 0; i < 10; i++ {
        bucket[i] = make([]int, len(nums)) // 初始化切片长度
    }
    var order [10]int // 存放每个桶真实存放数据的长度

    n := 1 // 控制元素的位数
    for i := 0 ; i < maxLength; i++ {
        for _, num := range nums {
            bucketIndex := num / n % 10 // 计算桶的下表
            bucket[bucketIndex][order[bucketIndex]] = num
            order[bucketIndex]++ // 尾下标 + 1
        }

        numIndex := 0
        // 从桶数组中依次取出所有元素放入原数组
        for i, bucketLength := range order {
            if bucketLength != 0 {
                for j := 0 ; j < bucketLength; j++ {
                    nums[numIndex] = bucket[i][j]
                    numIndex++
                }
                order[i] = 0 // 重置当前桶下标
            }
        }

        n *= 10
    }
}
