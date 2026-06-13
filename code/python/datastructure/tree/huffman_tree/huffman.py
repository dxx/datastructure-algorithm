"""
哈夫曼树
有 N 个权值作为 N 个叶子结点，构造一棵二叉树，如果该树的带权路径长度达到最小
称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。哈夫曼树是带权路
径长度最短的树，权值较大的结点离根较近。哈夫曼树又称为最优树。

构建步骤
1.将 w1、w2、…、wn 看成一个序列, 每个数据可以看做一个权值
2.将序列从小到大排序
3.选出两个根节点的权值最小的树合并，作为一棵新树的左、右子节点，且新树的根节点权值为其左、右子树根节点权值之和
4.从序列中删除选出的两个节点，并将新树加入序列
5.重复 2、3、4 步，直到序列中只剩一棵树为止，该树即为所求得的哈夫曼树
"""


class Node:
    def __init__(self, value: int) -> None:
        self.value = value
        self.left: Node | None = None
        self.right: Node | None = None


def create_huffman_tree(nums: list[int]) -> Node | None:
    if not nums:
        return None
    nodes = []
    for num in nums:
        nodes.append(Node(num))
    while len(nodes) > 1:
        # 排序
        nodes.sort(key=lambda node: node.value)
        left = nodes[0]  # 权值最小的元素
        right = nodes[1]  # 权值第二小的元素
        # 创建新的根节点
        root = Node(left.value + right.value)
        # 构建二叉树
        root.left = left
        root.right = right

        # 删除处理过的节点
        nodes.pop(0)
        nodes.pop(0)

        # 将二叉树加入到 nodes
        nodes.append(root)
    return nodes[0]


def pre_order(node: Node | None) -> None:
    if node is None:
        return
    print(node.value)
    pre_order(node.left)
    pre_order(node.right)
