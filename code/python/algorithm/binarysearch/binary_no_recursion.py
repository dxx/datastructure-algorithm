"""
二分法查找(非递归)
"""


def binary_search_no_recursion(nums: list[int], find_val: int) -> int:
    start = 0
    end = len(nums) - 1
    while start <= end:
        mid = (start + end) // 2
        if find_val < nums[mid]:  # 查找的值在左边
            end = mid - 1
        elif find_val > nums[mid]:  # 查找的值在右边
            start = mid + 1
        else:
            # 找到目标值的下标
            return mid
    return -1
