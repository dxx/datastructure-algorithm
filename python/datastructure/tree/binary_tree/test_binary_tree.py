import unittest
from binary_tree import (
    BinaryTreeNode,
    pre_order,
    infix_order,
    post_order,
    pre_order_search,
    infix_order_search,
    post_order_search,
)

def init_node() -> BinaryTreeNode:
    root = BinaryTreeNode(1)
    node2 = BinaryTreeNode(2)
    node3 = BinaryTreeNode(3)
    node4 = BinaryTreeNode(4)
    node5 = BinaryTreeNode(5)

    # 手动建立树的关系
    root.left = node2
    root.right = node5
    node2.left = node3
    node2.right = node4

    return root

class Test(unittest.TestCase):
    
    def test_order(self):
        root = init_node()

        print("======前序遍历======")
        pre_order(root)

        print("======中续遍历======")
        infix_order(root)

        print("======后续遍历======")
        post_order(root)

    def test_search(self):
        root = init_node()

        no = 4

        print("======前序查找======")
        print("查找no=" + str(no))
        node = pre_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))

        no = 4
        print("======中序查找======")
        print("查找no=" + str(no))
        node = infix_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))

        no = 4
        print("======后序查找======")
        print("查找no=" + str(no))
        node = post_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))

