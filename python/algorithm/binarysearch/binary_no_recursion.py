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


def main() -> None:
    value = 100
    nums = [1, 8, 10, 89, 100, 100, 123]
    index = binary_search_no_recursion(nums, value)
    if index != -1:
        print(f"找到 {value}, 下标为 {index}")
    else:
        print(f"未找到 {value}")


if __name__ == "__main__":
    main()
