package main

import "fmt"

// 堆排序
// 堆排序是利用堆这种数据结构而设计的一种排序算法，堆排序是一种选择排序
// 堆是一个具有特殊性质的完全二叉树，任意非叶子节点的值大于或等于左右子
// 节点的值，或者任意非叶子节点的值小于或等于左右子节点的值

func heapSort(nums []int) {
    if nums == nil {
        return
    }
    // 调整第 1 个节点 [1 7 5 2 8] => [1 8 5 2 7]
    // adjustHeap(nums, 1, len(nums))
    // 调整第 0 个节点 [1 8 5 2 7] => [8 7 5 2 1]
    // adjustHeap(nums, 0, len(nums))

    // 调整所有叶子节点, 构造成一个大顶堆
    // 堆顶的根节点就是序列的最大值
    for i := len(nums)/2 - 1; i >= 0; i-- {
        adjustHeap(nums, i, len(nums))
    }
    // 将堆顶的根节点和叶子节点进交换，此时叶子节点就是最大值
    for i := len(nums) - 1; i > 0; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        // 对于剩余的元素重新构造成大顶堆或小顶堆
        adjustHeap(nums, 0, i)
    }
}

// 调整堆, 使其成为大顶堆
// i: 当前需要调整的节点下标
// count: 调整次数
func adjustHeap(nums []int, i, count int) {
    temp := nums[i] // 当前节点
    for j := 2*i + 1; j < count; j = 2*j + 1 {
        // 左子节点小于右子节点
        if j+1 < count && nums[j] < nums[j+1] {
            j++ // 指向右子节点
        }
        // 子节点比父节点大
        if nums[j] > temp {
            // 将节点赋值给父节点
            nums[i] = nums[j]
            i = j // 修改成下一个子节点
        } else {
            // 跳出循环，因为调整顺序为从左至右，从下至上，子树是已经调整好的堆
            break
        }
    }
    // 放入到最终位置
    nums[i] = temp
}

func main() {
    nums := []int{1, 7, 5, 2, 8}

    fmt.Printf("排序前: %v\n", nums)

    heapSort(nums)

    fmt.Printf("排序后: %v\n", nums)
}
