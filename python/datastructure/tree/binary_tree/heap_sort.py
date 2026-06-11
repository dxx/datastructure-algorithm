"""
堆排序
堆排序是利用堆这种数据结构而设计的一种排序算法，堆排序是一种选择排序
堆是一个具有特殊性质的完全二叉树，任意非叶子节点的值大于或等于左右子
节点的值，或者任意非叶子节点的值小于或等于左右子节点的值

最后一个非叶子节点计算公式：length / 2 - 1
第 n 个下标节点的左子节点计算公式：2 * n + 1
第 n 个下标节点的右子节点计算公式：2 * n + 2
"""


def heap_sort(nums: list[int]) -> None:
    if not nums:
        return
    # 调整所有叶子节点, 构造成一个大顶堆
    # 堆顶的根节点就是序列的最大值
    for i in range(len(nums) // 2 - 1, -1, -1):
        adjust_heap(nums, i, len(nums))
    # 将堆顶的根节点和叶子节点进交换，此时叶子节点就是最大值
    for i in range(len(nums) - 1, 0, -1):
        nums[0], nums[i] = nums[i], nums[0]
        # 对于剩余的元素重新构造成大顶堆
        adjust_heap(nums, 0, i)


def adjust_heap(nums: list[int], i: int, count: int) -> None:
    """
    调整堆, 使其成为大顶堆
    i: 当前需要调整的节点下标
    count: 调整次数
    """
    temp = nums[i]  # 当前节点
    j = 2 * i + 1
    while j < count:
        # 左子节点小于右子节点
        if j + 1 < count and nums[j] < nums[j + 1]:
            j += 1  # 指向右子节点
        if nums[j] > temp:
            # 将节点赋值给父节点
            nums[i] = nums[j]
            i = j  # 修改成下一个子节点
        else:
            # 跳出循环，因为调整顺序为从左至右，从下至上，子树是已经调整好的堆
            break
        j = 2 * j + 1
    nums[i] = temp
