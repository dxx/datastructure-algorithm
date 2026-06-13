import unittest
from avl_tree import AVLTree, BinaryTreeNode


class Test(unittest.TestCase):
    
    def test_left_rotate(self):
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

    def test_right_rotate(self):
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

    def test_double_rotate(self):
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
