package main

import "fmt"

// 线索化二叉树
// 对于 n 个结点的二叉树，在二叉链存储结构中有 n+1 个空链域，用这些
// 空链域存放该结点的前驱结点和后继结点的指针，这些指针称为线索，加上
// 线索的二叉树称为线索二叉树。对二叉树以某种遍历方式（如先序、中序、
// 后序或层次等）进行遍历，使其变为线索二叉树的过程称为对二叉树进行线索化

type ThreadedBinaryTreeNode struct {
    no    int                     // 编号
    left  *ThreadedBinaryTreeNode // 左子节点
    right *ThreadedBinaryTreeNode // 右子节点

    // 增加两个标记
    leftTag  int // 左节点标记。如果 leftTag = 0, left 表示左子节点, 如果为 leftTag = 1, left 表示前驱节点
    rightTag int // 右节点标记。如果 rightTag = 0, right 表示右子节点, 如果为 rightTag = 1, right 表示后继节点
}

// 从最左边至最右边查找指定节点（测试使用）
func (threadedNode *ThreadedBinaryTreeNode) search(no int) *ThreadedBinaryTreeNode {
    leftChildNode := threadedNode
    for leftChildNode.left != nil {
        leftChildNode = leftChildNode.left
    }

    node := leftChildNode
    // 遍历右边
    for node != nil {
        if node.no == no {
            break
        }
        node = node.right
    }
    return node
}

var previous *ThreadedBinaryTreeNode // 记录遍历时的上一个结点

// 中序线索化二叉树
func infixThreadTree(node *ThreadedBinaryTreeNode) {
    if node == nil {
        return
    }
    // 线索化左子节点
    infixThreadTree(node.left)

    // 线索化当前结点
    // 如果 left 为 nil, 处理前驱节点
    if node.left == nil {
        node.left = previous
        node.leftTag = 1 // 修改标记
    }

    // 如果 right 为 nil, 处理后继节点
    if previous != nil && previous.right == nil {
        previous.right = node // 将上一个节点的后继节点指向当前节点
        previous.rightTag = 1 // 修改标记
    }

    // 修改 previous
    previous = node

    // 线索化右子节点
    infixThreadTree(node.right)
}

// 中序遍历线索化二叉树
func infixOrderThreadedTree(node *ThreadedBinaryTreeNode) {
    currentNode := node
    for currentNode != nil {
        // 循环找到 leftTag = 0 的节点, 第一个找到的就是最左边的叶子节点
        for currentNode.leftTag == 0 {
            currentNode = currentNode.left
        }
        // 输出当前节点
        fmt.Printf("id:%v\n", currentNode.no)
        // 循环输出后继节点
        for currentNode.rightTag == 1 {
            currentNode = currentNode.right
            fmt.Printf("id:%v\n", currentNode.no)
        }
        // 移动当前节点
        currentNode = currentNode.right
    }
}

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

func testInfixThreadedTree() {
    root := initThreadedNode()

    infixThreadTree(root)

    // 获取 no = 10 的结点，输出前驱和后继节点
    node := root.search(10)

    fmt.Printf("no=%v的前驱节点为%v\n", node.no, node.left.no)
    fmt.Printf("no=%v的后继节点为%v\n", node.no, node.right.no)
}

func testInfixOrderThreadedTree() {
    root := initThreadedNode()

    infixThreadTree(root)

    infixOrderThreadedTree(root)
}

func main() {
    // testInfixThreadedTree()
    // testInfixOrderThreadedTree()
}
