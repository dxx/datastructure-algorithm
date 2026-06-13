"""
选择排序
1.假定第一元素为最大或最小的元素
2.找出最大或最小的元素的小标，循环 length - 1 次
3.每次循环完成后将最大值或最小值和本次循环的第一个元素交换
"""


def select_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    length = len(nums)
    for i in range(length - 1):
        # 记录最小值的下标
        min_index = i
        for j in range(i + 1, length):
            if nums[min_index] > nums[j]:
                # 修改最小值下标
                min_index = j
        # 优化：判断是否需要交换
        if min_index != i:
            swap(nums, i, min_index)


def swap(nums: list[int], i: int, j: int) -> None:
    temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
