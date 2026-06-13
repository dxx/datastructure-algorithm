package sort

import (
    "fmt"
    "testing"
)

func TestQuickSort(t *testing.T) {
    nums := []int{5, 1, 8, 3, 7, 2, 9, 4, 6}
    fmt.Printf("排序前: %v\n", nums)
    quickSort(nums, 0, len(nums) - 1)
    fmt.Printf("排序后: %v\n", nums)
}
