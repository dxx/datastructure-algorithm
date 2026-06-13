package sort

import (
    "fmt"
    "testing"
)

func TestSelectSort(t *testing.T) {
    nums := []int{3, 5, 7, 1, 2, 4, 9, 6, 8}
    fmt.Printf("排序前: %v\n", nums)
    selectSort(nums)
    fmt.Printf("排序后: %v\n", nums)
}
