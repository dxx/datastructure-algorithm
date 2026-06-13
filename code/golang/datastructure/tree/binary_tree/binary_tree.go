package binary_tree

import "fmt"

// 二叉树
// 二叉树是每个结点最多有两个子树的树结构
// 通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）
// 二叉树常被用于实现二叉查找树和二叉堆

type BinaryTreeNode struct {
    no    int             // 编号
    left  *BinaryTreeNode // 左子节点
    right *BinaryTreeNode // 右子节点
}

// 前序遍历
func preOrder(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 当前节点
    fmt.Println(node)
    // 遍历左子树
    preOrder(node.left)
    // 遍历右子树
    preOrder(node.right)
}

// 中序遍历
func infixOrder(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 遍历左子树
    infixOrder(node.left)
    // 当前节点
    fmt.Println(node)
    // 遍历右子树
    infixOrder(node.right)
}

// 后序遍历
func postOrder(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 遍历左子树
    postOrder(node.left)
    // 遍历右子树
    postOrder(node.right)
    // 当前节点
    fmt.Println(node)
}

// 前序查找
func preOrderSearch(node *BinaryTreeNode, no int) *BinaryTreeNode {
    if node == nil {
        return nil
    }
    fmt.Println("进入查找")
    if node.no == no {
        return node
    }
    // 左边查找
    returnNode := preOrderSearch(node.left, no)
    if returnNode != nil {
        // 左边找到了节点，返回
        return returnNode
    }
    // 右边查找
    return preOrderSearch(node.right, no)
}

// 中序查找
func infixOrderSearch(node *BinaryTreeNode, no int) *BinaryTreeNode {
    if node == nil {
        return nil
    }
    // 左边查找
    returnNode := infixOrderSearch(node.left, no)
    if returnNode != nil {
        // 左边找到了节点，返回
        return returnNode
    }
    fmt.Println("进入查找")
    if node.no == no {
        return node
    }
    // 右边查找
    return infixOrderSearch(node.right, no)
}

// 后序查找
func postOrderSearch(node *BinaryTreeNode, no int) *BinaryTreeNode {
    if node == nil {
        return nil
    }
    var returnNode *BinaryTreeNode
    // 左边查找
    returnNode = postOrderSearch(node.left, no)
    if returnNode != nil {
        // 左边找到了节点，返回
        return returnNode
    }
    // 右边查找
    returnNode = postOrderSearch(node.right, no)
    if returnNode != nil {
        // 右边找到了节点，返回
        return returnNode
    }
    fmt.Println("进入查找")
    if node.no == no {
        returnNode = node
    }
    return returnNode
}

// 删除节点
// 如果被删除的节点是叶子节点，直接删除
// 如果被删除的节点是非叶子节点，删除该树
func deleteNode(node *BinaryTreeNode, no int) {
    // 当前节点为空直接返回
    if node == nil {
        return
    }
    // 如果左子节点就是要删除结点，就将 binaryTree.left = nil, 并且就返回(结束递归删除)
    if node.left.no == no {
        node.left = nil
        return
    }
    // 如果右子节点就是要删除结点，就将 binaryTree.right = nil, 并且就返回(结束递归删除)
    if node.right.no == no {
        node.right = nil
        return
    }
    // 左边查找删除
    deleteNode(node.left, no)
    // 右边查找删除
    deleteNode(node.right, no)
}

func (binaryTreeNode *BinaryTreeNode) String() string {
    return fmt.Sprintf("no:%d", binaryTreeNode.no)
}
