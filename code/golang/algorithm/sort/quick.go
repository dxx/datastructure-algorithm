package sort

// 快速排序
// 1.获取最中间的元素，然后分别将左边和右边的元素分别和最中间的元素进行比较
// 2.左边找出比中间大的元素，右边找出比中间小的元素，交换左右的位置
// 3.一次比较完成后，将左边和右边分别递归重复以上操作
func quickSort(nums []int, start, end int) {
    l := start
    r := end
    // 获取最中间的元素
    centerValue := nums[(start + end) / 2]
    for l < r {
        // 从左边找出比中间元素大的元素值
        for nums[l] < centerValue {
            l++
        }
        // 从右边找出比中间元素小的值
        for nums[r] > centerValue {
            r--
        }
        if l >= r {
            break
        }
        // 交换左边和右边的值
        nums[l], nums[r] = nums[r], nums[l]

        // 交换后，nums[l] 的值等于 centerValue，r 前移
        if nums[l] == centerValue {
            r--
        }
        // 交换后，nums[r] 的值等于 centerValue，l 后移
        if nums[r] == centerValue {
            l++
        }
    }
    // 避免死递归
    if l == r {
        l++
        r--
    }
    if start < r {
        quickSort(nums, start, r)
    }
    if end > l {
        quickSort(nums, l, end)
    }
}
