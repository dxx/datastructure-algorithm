package main

import "fmt"

// 二叉排序树
// 二叉排序树（Binary Sort Tree），又称二叉查找树（Binary Search Tree），也叫二叉搜索树
// 在一般情况下，查询效率比链表结构要高。对于任何一个非叶子节点，它的左子节点的值小于自身的值，
// 它的右子节点的值大于自身的值，具有这样性质的二叉树称为二叉排序树

// 树节点结构体
type BinaryTreeNode struct {
    no    int             // 编号
    left  *BinaryTreeNode // 左子节点
    right *BinaryTreeNode // 右子节点
}

func (sortTreeNode *BinaryTreeNode) String() string {
    return fmt.Sprintf("no:%d", sortTreeNode.no)
}

// 二叉排序树结构体
type BinarySortTree struct {
    root *BinaryTreeNode // 树的根节点
}

func NewBinarySortTree() *BinarySortTree {
    return &BinarySortTree{}
}

// 添加结点
func (sortTree *BinarySortTree) Add(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 根节点为 nil，要添加的节点作为根节点
    if sortTree.root == nil {
        sortTree.root = node
        return
    }
    sortTree.add(sortTree.root, node)
}

func (sortTree *BinarySortTree) add(root, node *BinaryTreeNode) {
    // 要添加的节点小于根节点
    if node.no < root.no {
        // 左子节点为 nil，直接添加为左子节点
        if root.left == nil {
            root.left = node
            return
        }
        // 左递归
        sortTree.add(root.left, node)
    } else {
        // 右子节点为 nil，直接添加为右子节点
        if root.right == nil {
            root.right = node
            return
        }
        // 右递归
        sortTree.add(root.right, node)
    }
}

func (sortTree *BinarySortTree) search(no int) (*BinaryTreeNode, *BinaryTreeNode) {
    if sortTree.root == nil {
        return nil, nil
    }
    // 查找的节点就是根节点
    if sortTree.root.no == no {
        return nil, sortTree.root
    }
    return sortTree.recursionSearch(sortTree.root, no)
}

// 递归查找指定节点
// 返回查找到的父节点和查找到的节点
func (sortTree *BinarySortTree) recursionSearch(node *BinaryTreeNode, no int) (*BinaryTreeNode, *BinaryTreeNode) {
    if node == nil {
        return nil, nil
    }
    if node.left != nil && node.left.no == no {
        return node, node.left
    }
    if node.right != nil && node.right.no == no {
        return node, node.right
    }
    // 判断是往左边还是往右边查找
    if no < node.no {
        return sortTree.recursionSearch(node.left, no)
    } else {
        return sortTree.recursionSearch(node.right, no)
    }
}

// 删除节点
// 1.节点是叶子节点直接删除
// 2.节点是子节点且只有一颗子树，左子树或右子树。如果被删除的节点是父节点的
//   左子节点，将父节点的左子节点指向该删除节点的子树，如果是父节点的右子节
//   点，则将父节点的右子节点指向该删除节点的子树
// 3.节点是子节点且只有两颗子树。从被删除节点的右子节点的左子树中找到最小值的节点，将
//   其删除，然后将该节点的值赋值给被删除的节点
func (sortTree *BinarySortTree) Delete(no int) {
    parentNode, node := sortTree.search(no)
    // 没有找到要删除的节点
    if node == nil {
        return
    }

    // 当前节点为叶子节点
    if node.left == nil && node.right == nil {
        // 被删除的节点为根节点
        if parentNode == nil {
            sortTree.root = nil
            return
        }
        // 当前节点为父节点的左子节点
        if parentNode.left != nil && parentNode.left.no == no {
            parentNode.left = nil
        }
        // 当前节点为父节点的右子节点
        if parentNode.right != nil && parentNode.right.no == no {
            parentNode.right = nil
        }
        return
    }

    // 当前节点有两颗子树
    if node.left != nil && node.right != nil {
        // 把右子节点作为根节点，从左边开始遍历到最后一个叶子节点
        leftChildNode := node.right
        for leftChildNode.left != nil {
            leftChildNode = leftChildNode.left
        }
        // 删除最小的叶子节点
        sortTree.Delete(leftChildNode.no)

        // 替换掉被删除节点的值
        node.no = leftChildNode.no
    } else { // 当前节点只有一颗子树
        var replaceNode *BinaryTreeNode
        if node.left != nil {
            replaceNode = node.left
        } else if node.right != nil {
            replaceNode = node.right
        }

        // 父节点为 nil，表示根节点
        if parentNode == nil {
            sortTree.root = replaceNode
            return
        }

        // 当前节点为父节点的左子节点
        if parentNode.left != nil && parentNode.left.no == no {
            parentNode.left = replaceNode
        }
        // 当前节点为父节点的右子节点
        if parentNode.right != nil && parentNode.right.no == no {
            parentNode.right = replaceNode
        }
    }
}

func (sortTree *BinarySortTree) InfixOrder() {
    if sortTree.root == nil {
        return
    }
    sortTree.infixOrder(sortTree.root)
}

func (sortTree *BinarySortTree) infixOrder(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    sortTree.infixOrder(node.left)
    fmt.Println(node)
    sortTree.infixOrder(node.right)
}

func main() {
    nos := []int{8, 5, 10, 3, 6, 9, 12, 2}
    binarySortTree := NewBinarySortTree()
    for _, no := range nos {
        binarySortTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("======中序遍历======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(6)

    fmt.Println("======删除叶子节点6======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(5)

    fmt.Println("======删除只有一颗子树的节点5======")
    binarySortTree.InfixOrder()

    binarySortTree.Delete(10)

    fmt.Println("======删除有两颗子树的节点10======")
    binarySortTree.InfixOrder()
}
