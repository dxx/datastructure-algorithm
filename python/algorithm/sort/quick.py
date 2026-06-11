"""
快速排序
1.获取最中间的元素，然后分别将左边和右边的元素分别和最中间的元素进行比较
2.左边找出比中间大的元素，右边找出比中间小的元素，交换左右的位置
3.一次比较完成后，将左边和右边分别递归重复以上操作
"""


def quick_sort(nums: list[int], start: int, end: int) -> None:
    left = start
    right = end
    # 获取最中间的元素
    center_value = nums[(start + end) // 2]
    while left < right:
        # 从左边找出比中间元素大的元素值
        while nums[left] < center_value:
            left += 1
        # 从右边找出比中间元素小的值
        while nums[right] > center_value:
            right -= 1
        if left >= right:
            break
        # 交换左边和右边的值
        swap(nums, left, right)

        # 交换后，nums[left] 的值等于 centerValue，r 前移
        if nums[left] == center_value:
            right -= 1
        # 交换后，nums[right] 的值等于 centerValue，l 后移
        if nums[right] == center_value:
            left += 1
    # 避免死递归
    if left == right:
        left += 1
        right -= 1
    if start < right:
        quick_sort(nums, start, right)
    if end > left:
        quick_sort(nums, left, end)


def swap(nums: list[int], i: int, j: int) -> None:
    temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
