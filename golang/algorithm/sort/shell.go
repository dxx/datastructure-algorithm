package sort

// 希尔排序
// 1.计算出步长 step，step = length / 2
// 2.从 step 开始，循环到 length
// 3.将循环开始时的元素和比当前大 step 的元素进行比较
// 4.发现逆序则进行交换
// 交换法
func shellSort(nums []int) {
    if nums == nil {
        return
    }
    length := len(nums)
    // 控制步长
    for step := length / 2; step > 0; step /= 2 {
        for i := step; i < length; i++ {
            for j := i - step; j >= 0 && nums[j] > nums[j + step]; j -= step {
                // 前面的数比后面的数大，进行交换
                nums[j], nums[j + step] = nums[j + step], nums[j]
            }
        }
    }
}
