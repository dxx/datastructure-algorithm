package sort

import (
    "fmt"
    "testing"
)

func TestRadixSort(t *testing.T) {
    nums := []int{5, 1, 7, 13, 21, 32, 9, 66, 8, 20}
    fmt.Printf("排序前: %v\n", nums)
    radixSort(nums)
    fmt.Printf("排序后: %v\n", nums)
}
