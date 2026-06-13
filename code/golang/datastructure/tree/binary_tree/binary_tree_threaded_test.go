package binary_tree

import (
    "fmt"
    "testing"
)

func initThreadedNode() *ThreadedBinaryTreeNode {
    root := &ThreadedBinaryTreeNode{no: 1}
    node2 := &ThreadedBinaryTreeNode{no: 2}
    node3 := &ThreadedBinaryTreeNode{no: 6}
    node4 := &ThreadedBinaryTreeNode{no: 8}
    node5 := &ThreadedBinaryTreeNode{no: 10}
    node6 := &ThreadedBinaryTreeNode{no: 16}

    // 手动建立树的关系
    root.left = node2
    root.right = node3
    node2.left = node4
    node2.right = node5
    node3.left = node6

    return root
}

func TestInfixThreadedTree(t *testing.T) {
    root := initThreadedNode()

    infixThreadTree(root)

    // 获取 no = 10 的结点，输出前驱和后继节点
    node := root.search(10)

    fmt.Printf("no=%v的前驱节点为%v\n", node.no, node.left.no)
    fmt.Printf("no=%v的后继节点为%v\n", node.no, node.right.no)
}

func TestInfixOrderThreadedTree(t *testing.T) {
    root := initThreadedNode()

    infixThreadTree(root)

    infixOrderThreadedTree(root)
}
