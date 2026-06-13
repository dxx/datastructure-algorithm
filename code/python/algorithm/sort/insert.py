"""
插入排序
1.从第二个元素开始循环，循环到元素末尾，左边分为有序列表，右边分为无序列表
2.将右边无序列表的第一个元素，标记为要插入的值（insertValue），并记录要插入的位置（insertIndex）
3.依次和左边的无序列表中的元素比较，如果顺序颠倒，则将有序列表中当前被比较的元素后移，同时修改 insertIndex
4.最后将 insertValue 插入到 insertIndex 位置
"""


def insert_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    for i in range(1, len(nums)):
        # 要插入的下标
        insert_index = i - 1
        # 保存插入的值
        insert_value = nums[i]
        # 从小到大
        while insert_index >= 0 and nums[insert_index] > insert_value:
            # 元素后移
            nums[insert_index + 1] = nums[insert_index]
            # 下标前移
            insert_index -= 1
        # 优化：判断如果当前位置发生移动，就插入
        if insert_index + 1 != i:
            # 插入
            nums[insert_index + 1] = insert_value
