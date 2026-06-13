"""
二叉树顺序存储
将二叉树存储在一个数组中，通过存储元素的下标反映元素之间的父子关系
用一组连续的存储单元存放二又树中的结点元素，一般按照二叉树结点自上向下、自左向右的顺序存储
使用此存储方式，结点的前驱和后继不一定是它们在逻辑上的邻接关系，非常适用于满二又树和完全二又树
采用顺序存储能够最大地节省存储空间，可以利用数组元素下标值确定结点在二叉树中的位置以及结点之间的关系

计算顺序存储二叉树节点下标的方法如下:
第 n 个节点的左子节点下标为 2*n+1
第 n 个节点的右子节点下标为 2*n+2
第 n 个节点的父节点下标为 (n-1)/2
n 表示二叉树中第几个节点，对应该节点在数组中的位置
位置从 0 开始，顺序为从上之下，从左至右
"""


class SeqBinaryTree:
    def __init__(self, array: list[int]) -> None:
        self.array = array

    def pre_order(self) -> None:
        self._pre_order_from_index(0)

    def _pre_order_from_index(self, index: int) -> None:
        if not self.array:
            return
        length = len(self.array)
        if length == 0 or index >= length:
            return
        # 当前节点
        print(self.array[index])
        # 左子节点下标
        left_index = 2 * index + 1
        # 右子节点下标
        right_index = 2 * index + 2
        # 向左遍历
        self._pre_order_from_index(left_index)
        # 向右遍历
        self._pre_order_from_index(right_index)
