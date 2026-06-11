"""
二叉树
二叉树是每个结点最多有两个子树的树结构
通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）
二叉树常被用于实现二叉查找树和二叉堆
"""


class BinaryTreeNode:
    def __init__(self, no: int) -> None:
        self.no = no
        self.left: BinaryTreeNode | None = None
        self.right: BinaryTreeNode | None = None


def pre_order(node: BinaryTreeNode | None) -> None:
    """前序遍历"""
    if not node:
        return
    # 当前节点
    print(node.no)
    # 遍历左子树
    pre_order(node.left)
    # 遍历右子树
    pre_order(node.right)


def infix_order(node: BinaryTreeNode | None) -> None:
    """中序遍历"""
    if not node:
        return
    # 遍历左子树
    infix_order(node.left)
    # 当前节点
    print(node.no)
    # 遍历右子树
    infix_order(node.right)


def post_order(node: BinaryTreeNode | None) -> None:
    """后序遍历"""
    if not node:
        return
    # 遍历左子树
    post_order(node.left)
    # 遍历右子树
    post_order(node.right)
    # 当前节点
    print(node.no)


def pre_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """前序查找"""
    if not node:
        return None
    print("进入查找")
    if node.no == no:
        return node
    # 左边查找
    return_node = pre_order_search(node.left, no)
    if return_node:
        # 左边找到了节点，返回
        return return_node
    # 右边查找
    return pre_order_search(node.right, no)


def infix_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """中序查找"""
    if not node:
        return None
    # 左边查找
    return_node = infix_order_search(node.left, no)
    if return_node:
        # 左边找到了节点，返回
        return return_node
    print("进入查找")
    if node.no == no:
        return node
    # 右边查找
    return infix_order_search(node.right, no)


def post_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """后序查找"""
    if not node:
        return None
    # 左边查找
    return_node = post_order_search(node.left, no)
    if return_node:
        # 左边找到了节点，返回
        return return_node
    # 右边查找
    return_node = post_order_search(node.right, no)
    if return_node is not None:
        # 右边找到了节点，返回
        return return_node
    print("进入查找")
    if node.no == no:
        return_node = node
    return return_node
