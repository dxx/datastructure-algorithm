"""
冒泡排序
1.从当前元素起，向前依次比较每一对相邻元素，若逆序则交换
2.对所有元素均重复以上步骤，直至最后一个元素
"""


def bubble_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    length = len(nums)
    # 外循环为排序趟数，length 个数进行 length - 1 趟
    for i in range(length - 1):
        # 内循环为每趟比较的次数，第 i 趟比较 length - i 次
        for j in range(length - 1, i, -1):
            # 相邻元素比较比较大小，然后交换位置
            if nums[j] < nums[j - 1]:
                swap(nums, j, j - 1)
        print("第 " + str(i + 1) + " 趟排序结果:" + str(nums))


def optimize_bubble_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    length = len(nums)
    is_change = False  # 标记是否发生交换
    for i in range(length - 1):
        for j in range(length - 1, i, -1):
            if nums[j] < nums[j - 1]:
                swap(nums, j, j - 1)
                is_change = True  # 发生交换
        print("第 " + str(i + 1) + " 趟排序结果:" + str(nums))
        if not is_change:
            break  # 跳出循环，终止比较
        is_change = False  # 重置


def swap(nums: list[int], i: int, j: int) -> None:
    temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
