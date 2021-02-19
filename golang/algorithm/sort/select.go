package sort

// 选择排序
// 1.假定第一元素为最大或最小的元素
// 2.找出最大或最小的元素的小标，循环 length - 1 次
// 3.每次循环完成后将最大值或最小值和本次循环的第一个元素交换
func selectSort(nums []int) {
    if nums == nil {
        return
    }
    length := len(nums)
    for i := 0; i < length - 1; i++ {
        // 记录最小值的下标
        minIndex := i
        for j := i + 1; j < length; j++ {
            if nums[minIndex] > nums[j] {
                // 修改最小值下标
                minIndex = j
            }
        }
        // 优化：判断是否需要交换
        if minIndex != i {
            nums[i], nums[minIndex] = nums[minIndex], nums[i]
        }
    }
}
