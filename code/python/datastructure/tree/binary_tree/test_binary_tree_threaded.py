import unittest
from binary_tree_threaded import (
    ThreadedBinaryTreeNode,
    infix_thread_tree,
    infix_order_threaded_tree
)

def init_threaded_node() -> ThreadedBinaryTreeNode:
    root = ThreadedBinaryTreeNode(1)
    node2 = ThreadedBinaryTreeNode(2)
    node3 = ThreadedBinaryTreeNode(6)
    node4 = ThreadedBinaryTreeNode(8)
    node5 = ThreadedBinaryTreeNode(10)
    node6 = ThreadedBinaryTreeNode(16)

    # 手动建立树的关系
    root.left = node2
    root.right = node3
    node2.left = node4
    node2.right = node5
    node3.left = node6

    return root


class Test(unittest.TestCase):
    
    def test_infix_threaded_tree(self):
        global previous
        previous = None
        root = init_threaded_node()

        infix_thread_tree(root)

        # 获取 no = 10 的结点，输出前驱和后继节点
        node = root.search(10)
        if node is None or node.left is None or node.right is None:
            return

        print("no=" + str(node.no) + "的前驱节点为" + str(node.left.no))
        print("no=" + str(node.no) + "的后继节点为" + str(node.right.no))


    def test_infix_order_threaded_tree(self):
        global previous
        previous = None
        root = init_threaded_node()

        infix_thread_tree(root)

        infix_order_threaded_tree(root)
