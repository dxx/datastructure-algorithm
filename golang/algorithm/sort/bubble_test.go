package sort

import (
    "fmt"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    nums := []int{1, 5, 7, 3, 2, 4, 9, 6, 8}
    fmt.Printf("交换前: %v\n", nums)
    bubbleSort(nums)
    fmt.Printf("交换后: %v\n", nums)

    nums2 := []int{1, 5, 7, 3, 2, 4, 9, 6, 8}
    fmt.Printf("优化前: %v\n", nums2)
    optimizeBubbleSort(nums2)
    fmt.Printf("优化后: %v\n", nums2)
}
