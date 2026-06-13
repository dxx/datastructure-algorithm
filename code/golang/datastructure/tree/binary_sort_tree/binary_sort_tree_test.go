package binary_sort_tree

import (
    "fmt"
    "testing"
)

func TestBinarySortTree(t *testing.T) {
    nos := []int{8, 5, 10, 3, 6, 9, 12, 2}
    binarySortTree := NewBinarySortTree()
    for _, no := range nos {
        binarySortTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("======中序遍历======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(6)

    fmt.Println("======删除叶子节点 6======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(5)

    fmt.Println("======删除只有一颗子树的节点 5======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(10)

    fmt.Println("======删除有两颗子树的节点 10======")
    binarySortTree.InfixOrder()
}
