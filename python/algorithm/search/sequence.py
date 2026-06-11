"""
线性查找
"""


def sequence_search(nums: list[int] | None, num: int) -> int:
    if nums is None:
        return -1
    for i, value in enumerate(nums):
        if value == num:
            return i
    return -1
