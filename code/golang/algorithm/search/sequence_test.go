package search

import (
    "fmt"
    "testing"
)

func TestSequenceSearch(t *testing.T) {
    value := 8
    nums := []int{2, 5, 1, 7, 8, 16}
    index := sequenceSearch(nums, value)
    if index != -1 {
        fmt.Printf("%d 在 nums 中的下标为: %d\n", value, index)
    }
}
