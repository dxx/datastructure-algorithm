package binarysearch

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
