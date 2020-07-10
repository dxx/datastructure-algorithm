package main

import (
    "fmt"
    "math"
)

// AVL 树
// 在 AVL 树中任何节点的两个子树的高度最大差别为 1 所以它也被称为高度平衡树
// 增加和删除可能需要通过一次或多次树旋转来重新平衡这个树。AVL 树本质上是带了
// 平衡功能的二叉排序树（二叉查找树，二叉搜索树）

// 树节点结构体
type BinaryTreeNode struct {
    no    int             // 编号
    left  *BinaryTreeNode // 左子节点
    right *BinaryTreeNode // 右子节点
}

func (sortTreeNode *BinaryTreeNode) String() string {
    return fmt.Sprintf("no:%d", sortTreeNode.no)
}

// AVL 树结构体
type AVLTree struct {
    root *BinaryTreeNode // 树的根节点
}

func NewAVLTree() *AVLTree {
    return &AVLTree{}
}

// 左子树高度
func (avlTree *AVLTree) leftHeight() int {
    if avlTree.root == nil {
        return 0
    }
    return avlTree.height(avlTree.root.left)
}

// 右子树高度
func (avlTree *AVLTree) rightHeight() int {
    if avlTree.root == nil {
        return 0
    }
    return avlTree.height(avlTree.root.right)
}

// 计算节点树的高度
func (avlTree *AVLTree) height(node *BinaryTreeNode) int {
    if node == nil {
        return 0
    }
    var lHeight, rHeight int
    // 递归左子节点
    lHeight = avlTree.height(node.left)
    // 递归右子节点
    rHeight = avlTree.height(node.right)
    // 取最大值，1 表示当前节点的高度值
    return int(math.Max(float64(lHeight), float64(rHeight))) + 1
}

// 左旋转
func (avlTree *AVLTree) leftRotate(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 以当前节点为基础，创建一个新的节点，新节点的值等于当前节点的值
    newNode := &BinaryTreeNode{no: node.no}
    // 让新节点的左子节点指向当前节点的左子节点，右子节点指向当前节点的右子节点的左子节点
    newNode.left = node.left
    newNode.right = node.right.left
    // 把当前节点的值替换为右子节点的值，并把当前节点右子节点指向其右子节点的右子节点
    node.no = node.right.no
    node.right = node.right.right
    // 让当前节点的左子节点指向新创建的节点
    node.left = newNode
}

// 右旋转
func (avlTree *AVLTree) rightRotate(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 以当前节点为基础，创建一个新的节点，新节点的值等于根节点的值
    newNode := &BinaryTreeNode{no: node.no}
    // 让新节点的右子节点指向当前节点的右子节点，左子节点指向当前节点的左子节点的右子节点
    newNode.right = node.right
    newNode.left = node.left.right
    // 把当前节点的值替换为左子节点的值，并把当前节点左子节点指向其左子节点的左子节点
    node.no = node.left.no
    node.left = node.left.left
    // 让当前节点的右子节点指向新创建的节点
    node.right = newNode
}

// 添加节点
func (avlTree *AVLTree) Add(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    // 根节点为 nil，要添加的节点作为根节点
    if avlTree.root == nil {
        avlTree.root = node
        return
    }
    avlTree.add(avlTree.root, node)

    // 添加节点后判断是否需要旋转

    // 右边高度超过左边 1 个高度以上，进行左旋转
    if avlTree.rightHeight()-avlTree.leftHeight() > 1 {

        rightNode := avlTree.root.right
        // 右子节点不为 nil，并且右子节点的左子树高度大于右子节点的右子树高度
        if rightNode != nil &&
            avlTree.height(rightNode.left) > avlTree.height(rightNode.right) {
            // 将右子节点右旋转
            avlTree.rightRotate(rightNode)
        }

        avlTree.leftRotate(avlTree.root)
        return
    }

    // 左边高度超过右边 1 个高度以上，进行右旋转
    if avlTree.leftHeight()-avlTree.rightHeight() > 1 {

        leftNode := avlTree.root.left
        // 左子节点不为 nil, 并且左子节点的右子树高度大于左子节点的左子树高度
        if leftNode != nil &&
            avlTree.height(leftNode.right) > avlTree.height(leftNode.left) {
            // 将左子节点左旋转
            avlTree.leftRotate(leftNode)
        }

        avlTree.rightRotate(avlTree.root)
    }
}

func (avlTree *AVLTree) add(root, node *BinaryTreeNode) {
    // 要添加的节点小于根节点
    if node.no < root.no {
        // 左子节点为 nil，直接添加为左子节点
        if root.left == nil {
            root.left = node
            return
        }
        // 左递归
        avlTree.add(root.left, node)
    } else {
        // 右子节点为 nil，直接添加为右子节点
        if root.right == nil {
            root.right = node
            return
        }
        // 右递归
        avlTree.add(root.right, node)
    }
}

func (avlTree *AVLTree) InfixOrder() {
    if avlTree.root == nil {
        return
    }
    avlTree.infixOrder(avlTree.root)
}

func (avlTree *AVLTree) infixOrder(node *BinaryTreeNode) {
    if node == nil {
        return
    }
    avlTree.infixOrder(node.left)
    fmt.Println(node)
    avlTree.infixOrder(node.right)
}

func testLeftRotate() {
    nos := []int{3, 2, 5, 4, 6, 7}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("左旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点=%v\n", avlTree.root)

    fmt.Printf("左子树的高度为:%d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为:%d\n", avlTree.rightHeight())
}

func testRightRotate() {
    nos := []int{6, 4, 7, 3, 5, 2}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("右旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点=%v\n", avlTree.root)

    fmt.Printf("左子树的高度为:%d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为:%d\n", avlTree.rightHeight())
}

func testDoubleRotate() {
    nos := []int{6, 3, 7, 2, 4, 5}
    avlTree := NewAVLTree()
    for _, no := range nos {
        avlTree.Add(&BinaryTreeNode{no: no})
    }

    fmt.Println("双旋转后")

    avlTree.InfixOrder()

    fmt.Printf("根节点=%v\n", avlTree.root)

    fmt.Printf("左子树的高度为:%d\n", avlTree.leftHeight())
    fmt.Printf("右子树的高度为:%d\n", avlTree.rightHeight())
}

func main() {
    // testLeftRotate()
    // testRightRotate()
    // testDoubleRotate()
}
