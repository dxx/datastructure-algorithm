package binary_tree

import (
    "fmt"
    "testing"
)

func initNode() *BinaryTreeNode {
    root := &BinaryTreeNode{no: 1}
    node2 := &BinaryTreeNode{no: 2}
    node3 := &BinaryTreeNode{no: 3}
    node4 := &BinaryTreeNode{no: 4}
    node5 := &BinaryTreeNode{no: 5}

    // 手动建立树的关系
    root.left = node2
    root.right = node5
    node2.left = node3
    node2.right = node4

    return root
}

func TestOrder(t *testing.T) {
    root := initNode()

    fmt.Println("======前序遍历======")
    preOrder(root)

    fmt.Println("======中续遍历======")
    infixOrder(root)

    fmt.Println("======后续遍历======")
    postOrder(root)
}

func TestSearch(t *testing.T) {
    root := initNode()

    no := 4

    fmt.Println("======前序查找======")
    fmt.Printf("查找no=%v\n", no)
    node := preOrderSearch(root, no)
    fmt.Printf("查找结果: no=%v\n", node.no)

    no = 4
    fmt.Println("======中序查找======")
    fmt.Printf("查找no=%v\n", no)
    node = infixOrderSearch(root, no)
    fmt.Printf("查找结果: no=%v\n", node.no)

    no = 4
    fmt.Println("======后序查找======")
    fmt.Printf("查找no=%v\n", no)
    node = postOrderSearch(root, no)
    fmt.Printf("查找结果: no=%v\n", node.no)
}

func TestDelete(t *testing.T) {
    root := initNode()
    no := 2

    fmt.Println("======删除前======")
    preOrder(root)

    deleteNode(root, no)
    // 被删除的节点恰好是 root 节点
    if root.no == no {
        root = nil
    }

    fmt.Println("======删除后======")
    if root == nil {
        fmt.Println("二叉树为nil")
    } else {
        preOrder(root)
    }
}
