package binary_tree

import (
    "fmt"
    "testing"
)

func TestHeapSort(t *testing.T) {
    nums := []int{1, 7, 5, 2, 8}

    fmt.Printf("排序前: %v\n", nums)

    heapSort(nums)

    fmt.Printf("排序后: %v\n", nums)
}
