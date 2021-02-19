package avl_tree

import (
    "fmt"
    "testing"
)

func TestHeight(t *testing.T) {
    nos := []int{3, 2, 5, 4, 6, 7}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Printf("左子树的高度为: %d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为: %d\n", avlTree.rightHeight())
}

func TestLeftRotate(t *testing.T) {
    nos := []int{3, 2, 5, 4, 6, 7}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("左旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点 = %v\n", avlTree.root)

    fmt.Printf("左子树的高度为: %d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为: %d\n", avlTree.rightHeight())
}

func TestRightRotate(t *testing.T) {
    nos := []int{6, 4, 7, 3, 5, 2}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("右旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点 = %v\n", avlTree.root)

    fmt.Printf("左子树的高度为: %d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为: %d\n", avlTree.rightHeight())
}

func TestDoubleRotate(t *testing.T) {
    nos := []int{6, 3, 7, 2, 4, 5}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("双旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点 = %v\n", avlTree.root)

    fmt.Printf("左子树的高度为: %d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为: %d\n", avlTree.rightHeight())
}
