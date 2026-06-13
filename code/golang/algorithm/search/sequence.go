package search

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
