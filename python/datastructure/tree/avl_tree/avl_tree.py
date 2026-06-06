"""
AVL 树
在 AVL 树中任何节点的两个子树的高度最大差别为 1 所以它也被称为高度平衡树
增加和删除可能需要通过一次或多次树旋转来重新平衡这个树。AVL 树本质上是带了
平衡功能的二叉排序树（二叉查找树，二叉搜索树）
"""


class BinaryTreeNode:
    def __init__(self, no):
        self.no = no
        self.left: BinaryTreeNode | None = None
        self.right: BinaryTreeNode | None = None


class AVLTree:
    def __init__(self):
        self.root = None  # 树的根节点

    def add(self, node):
        """添加结点"""
        if node is None:
            return
        # 根节点为 nil，要添加的节点作为根节点
        if self.root is None:
            self.root = node
            return
        self._add(self.root, node)

        # 添加节点后判断是否需要旋转

        # 右边高度超过左边 1 个高度以上，进行左旋转
        if self._right_height() - self._left_height() > 1:
            right_node = self.root.right
            # 右子节点不为 null，并且右子节点的左子树高度大于右子节点的右子树高度
            if right_node is not None and self._height(right_node.left) > self._height(right_node.right):
                # 将右子节点右旋转
                self._right_rotate(right_node)

            self._left_rotate(self.root)
            return

        if self._left_height() - self._right_height() > 1:
            left_node = self.root.left
            # 左子节点不为 nil, 并且左子节点的右子树高度大于左子节点的左子树高度
            if left_node is not None and self._height(left_node.right) > self._height(left_node.left):
                # 将左子节点左旋转
                self._left_rotate(left_node)

            self._right_rotate(self.root)

    def _height(self, node):
        """计算节点的高度"""
        if not node:
            return 0
        return max(self._height(node.left), self._height(node.right)) + 1

    def _left_height(self):
        """左子树高度"""
        if self.root is None:
            return 0
        return self._height(self.root.left)

    def _right_height(self):
        """右子树高度"""
        if self.root is None:
            return 0
        return self._height(self.root.right)

    def _left_rotate(self, node):
        """左旋转"""
        if not node:
            return
        # 以当前节点为基础，创建一个新的节点，新节点的值等于当前节点的值
        new_node = BinaryTreeNode(node.no)
        # 让新节点的左子节点指向当前节点的左子节点，右子节点指向当前节点的右子节点的左子节点
        if node.right is None:
            return
        new_node.left = node.left
        new_node.right = node.right.left
        # 把当前节点的值替换为右子节点的值，并把当前节点右子节点指向其右子节点的右子节点
        node.no = node.right.no
        node.right = node.right.right
        # 让当前节点的左子节点指向新创建的节点
        node.left = new_node

    def _right_rotate(self, node):
        """右旋转"""
        if not node:
            return
        # 以当前节点为基础，创建一个新的节点，新节点的值等于根节点的值
        new_node = BinaryTreeNode(node.no)
        # 让新节点的右子节点指向当前节点的右子节点，左子节点指向当前节点的左子节点的右子节点
        if node.left is None:
            return
        new_node.right = node.right
        new_node.left = node.left.right
        # 把当前节点的值替换为左子节点的值，并把当前节点左子节点指向其左子节点的左子节点
        node.no = node.left.no
        node.left = node.left.left
        # 让当前节点的右子节点指向新创建的节点
        node.right = new_node

    def _add(self, root, node):
        # 要添加的节点小于根节点
        if node.no < root.no:
            # 左子节点为 null，直接添加为左子节点
            if root.left is None:
                root.left = node
                return
            # 左递归
            self._add(root.left, node)
        else:
            # 右子节点为 null，直接添加为左子节点
            if root.right is None:
                root.right = node
                return
            # 右递归
            self._add(root.right, node)

    def infix_order(self):
        self._infix_order(self.root)

    def _infix_order(self, node):
        if node is None:
            return
        self._infix_order(node.left)
        print("no:" + str(node.no))
        self._infix_order(node.right)


def test_left_rotate():
    nos = [3, 2, 5, 4, 6, 7]
    avl_tree = AVLTree()
    for no in nos:
        avl_tree.add(BinaryTreeNode(no))

    print("左旋转后")

    avl_tree.infix_order()

    if avl_tree.root is not None:
        print("根节点 = ", avl_tree.root.no)

    print("左子树的高度为: ", avl_tree._left_height())
    print("右子树的高度为: ", avl_tree._right_height())


def test_right_rotate():
    nos = [6, 4, 7, 3, 5, 2]
    avl_tree = AVLTree()
    for no in nos:
        avl_tree.add(BinaryTreeNode(no))

    print("右旋转后")

    avl_tree.infix_order()

    if avl_tree.root is not None:
        print("根节点 = ", avl_tree.root.no)

    print("左子树的高度为: ", avl_tree._left_height())
    print("右子树的高度为: ", avl_tree._right_height())


def test_double_rotate():
    nos = [6, 3, 7, 2, 4, 5]
    avl_tree = AVLTree()
    for no in nos:
        avl_tree.add(BinaryTreeNode(no))

    print("双旋转后")

    avl_tree.infix_order()

    if avl_tree.root is not None:
        print("根节点 = ", avl_tree.root.no)

    print("左子树的高度为: ", avl_tree._left_height())
    print("右子树的高度为: ", avl_tree._right_height())


def main():
    # test_left_rotate()
    # test_right_rotate()
    # test_double_rotate()
    pass


if __name__ == "__main__":
    main()
