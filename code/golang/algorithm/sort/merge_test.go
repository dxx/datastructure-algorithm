package sort

import (
    "fmt"
    "testing"
)

func TestMergeSort(t *testing.T) {
    nums := []int{5, 0, 1, 7, 3, 2, 4, 9, 6, 8}
    fmt.Printf("排序前: %v\n", nums)
    mergeSort(nums)
    fmt.Printf("排序后: %v\n", nums)
}
