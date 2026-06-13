package binary_tree

import (
    "fmt"
    "testing"
)

func TestSeqStorage(t *testing.T) {
    nos := []int{1, 2, 3, 4, 5, 6, 7}
    arrayBinaryTree := NewArrayBinaryTree(nos)

    fmt.Println("======前序遍历======")
    arrayBinaryTree.PreOrder()
}
