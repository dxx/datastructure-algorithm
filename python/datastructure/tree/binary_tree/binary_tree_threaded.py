"""
二叉树线索化
对于 n 个结点的二叉树，在二叉链存储结构中有 n+1 个空链域，用这些
空链域存放该结点的前驱结点和后继结点的指针，这些指针称为线索，加上
线索的二叉树称为线索二叉树。对二叉树以某种遍历方式（如先序、中序、
后序或层次等）进行遍历，使其变为线索二叉树的过程称为对二叉树进行线索化
"""


class ThreadedBinaryTreeNode:
    def __init__(self, no: int) -> None:
        self.no = no  # 编号
        self.left: ThreadedBinaryTreeNode | None = None  # 左子节点
        self.right: ThreadedBinaryTreeNode | None = None  # 右子节点

        # 增加两个标记
        self.left_tag = 0  # 左节点标记。如果 leftTag = 0, left 表示左子节点, 如果为 leftTag = 1, left 表示前驱节点
        self.right_tag = 0  # 右节点标记。如果 rightTag = 0, right 表示右子节点, 如果为 rightTag = 1, right 表示后继节点

    def search(self, no: int) -> "ThreadedBinaryTreeNode | None":
        """从最左边至最右边查找指定节点（测试使用）"""
        left_child_node = self
        while left_child_node.left is not None:
            left_child_node = left_child_node.left

        while left_child_node is not None:
            if left_child_node.no == no:
                break
            left_child_node = left_child_node.right
        return left_child_node


previous: ThreadedBinaryTreeNode | None = None  # 记录遍历时的上一个结点


def infix_thread_tree(node: ThreadedBinaryTreeNode | None) -> None:
    """中序线索化二叉树"""
    global previous
    if not node:
        return
    # 线索化左子节点
    infix_thread_tree(node.left)

    # 线索化当前结点
    # 如果 left 为 null, 处理前驱节点
    if node.left is None:
        node.left = previous
        node.left_tag = 1  # 修改标记

    # 如果 right 为 null, 处理后继节点
    if previous is not None and previous.right is None:
        previous.right = node  # 将上一个节点的后继节点指向当前节点
        previous.right_tag = 1  # 修改标记

    # 修改 previous
    previous = node

    # 线索化右子节点
    infix_thread_tree(node.right)


def infix_order_threaded_tree(node: ThreadedBinaryTreeNode | None) -> None:
    """中序遍历线索化二叉树"""
    current_node = node
    while current_node is not None:
        # 循环找到 leftTag = 0 的节点, 第一个找到的就是最左边的叶子节点
        while current_node.left_tag == 0:
            if current_node.left is None:
                return
            current_node = current_node.left
        # 输出当前节点
        print("id:" + str(current_node.no))
        # 循环输出后继节点
        while current_node.right_tag == 1:
            if current_node.right is None:
                return
            current_node = current_node.right
            print("id:" + str(current_node.no))
        # 移动当前节点
        current_node = current_node.right


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


def test_infix_threaded_tree() -> None:
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


def test_infix_order_threaded_tree() -> None:
    global previous
    previous = None
    root = init_threaded_node()

    infix_thread_tree(root)

    infix_order_threaded_tree(root)


def main() -> None:
    # test_infix_threaded_tree()
    # test_infix_order_threaded_tree()
    pass


if __name__ == "__main__":
    main()
