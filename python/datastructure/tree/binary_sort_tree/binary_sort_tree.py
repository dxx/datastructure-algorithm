"""
二叉排序树
二叉排序树（Binary Sort Tree），又称二叉查找树（Binary Search Tree），也叫二叉搜索树
在一般情况下，查询效率比链表结构要高。对于任何一个非叶子节点，它的左子节点的值小于自身的值，
它的右子节点的值大于自身的值，具有这样性质的二叉树称为二叉排序树
"""


class BinaryTreeNode:
    def __init__(self, no):
        self.no = no
        self.left = None
        self.right = None


class BinarySortTree:
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

    def search(self, no):
        if self.root is None:
            return None
        # 查找的节点就是根节点
        if self.root.no == no:
            return [None, self.root]
        return self._recursion_search(self.root, no)

    def _recursion_search(self, node, no):
        """
        递归查找指定节点
        返回查找到的父节点和查找到的节点
        """
        if node is None:
            return None
        if node.left is not None and node.left.no == no:
            return [node, node.left]
        if node.right is not None and node.right.no == no:
            return [node, node.right]
        # 判断是往左边还是往右边查找
        if no < node.no:
            return self._recursion_search(node.left, no)
        return self._recursion_search(node.right, no)

    def delete(self, no):
        """
        删除节点
        1.节点是叶子节点直接删除
        2.节点是子节点且只有一颗子树，左子树或右子树。如果被删除的节点是父节点的
          左子节点，将父节点的左子节点指向该删除节点的子树，如果是父节点的右子节
          点，则将父节点的右子节点指向该删除节点的子树
        3.节点是子节点且只有两颗子树。从被删除节点的右子节点的左子树中找到最小值的节点，将
          其删除，然后将该节点的值赋值给被删除的节点
        """
        binary_tree_nodes = self.search(no)
        # 没有找到要删除的节点
        if not binary_tree_nodes:
            return
        parent_node = binary_tree_nodes[0]
        node = binary_tree_nodes[1]
        # 当前节点为叶子节点
        if node.left is None and node.right is None:
            # 被删除的节点为根节点
            if parent_node is None:
                self.root = None
                return
            # 当前节点为父节点的左子节点
            if parent_node.left is not None and parent_node.left.no == no:
                parent_node.left = None
            # 当前节点为父节点的右子节点
            if parent_node.right is not None and parent_node.right.no == no:
                parent_node.right = None
            return
        # 当前节点有两颗子树
        if node.left is not None and node.right is not None:
            # 把右子节点作为根节点，从左边开始遍历到最后一个叶子节点
            left_child_node = node.right
            while left_child_node.left is not None:
                left_child_node = left_child_node.left
            # 删除最小的叶子节点
            self.delete(left_child_node.no)

            # 替换掉被删除节点的值
            node.no = left_child_node.no
        else:  # 当前节点只有一颗子树
            replace_node = None
            if node.left is not None:
                replace_node = node.left
            if node.right is not None:
                replace_node = node.right

            # 父节点为 null，表示根节点
            if parent_node is None:
                self.root = replace_node
                return
            # 当前节点为父节点的左子节点
            if parent_node.left is not None and parent_node.left.no == no:
                parent_node.left = replace_node
            # 当前节点为父节点的右子节点
            if parent_node.right is not None and parent_node.right.no == no:
                parent_node.right = replace_node

    def infix_order(self):
        self._infix_order(self.root)

    def _infix_order(self, node):
        if node is None:
            return
        self._infix_order(node.left)
        print("no:" + str(node.no))
        self._infix_order(node.right)


def main():
    nos = [8, 5, 10, 3, 6, 9, 12, 2]
    binary_sort_tree = BinarySortTree()
    for no in nos:
        binary_sort_tree.add(BinaryTreeNode(no))
    print("======中序遍历======")
    binary_sort_tree.infix_order()

    binary_sort_tree.delete(6)

    print("======删除叶子节点 6======")
    binary_sort_tree.infix_order()

    binary_sort_tree.delete(5)


    print("======删除只有一颗子树的节点 5======")
    binary_sort_tree.infix_order()

    binary_sort_tree.delete(10)

    print("======删除有两颗子树的节点 10======")
    binary_sort_tree.infix_order()


if __name__ == "__main__":
    main()
