"""
插值查找
基于二分法查找，步骤和二分法查找一样
中间下标计算公式: mid = start + (end - start) * (value - array[start]) / (array[end] - array[start])
"""


def insert_val_search(nums: list[int], start: int, end: int, find_val: int) -> int:
    if start > end or find_val < nums[start] or find_val > nums[end]:
        return -1
    if nums[end] == nums[start]:
        return start if nums[start] == find_val else -1
    # 根据 findVal 自适应计算中间下标
    mid = start + (end - start) * (find_val - nums[start]) // (nums[end] - nums[start])
    print("mid: " + str(mid))
    if find_val < nums[mid]:
        # 向左递归
        return insert_val_search(nums, start, mid - 1, find_val)
    if find_val > nums[mid]:
        # 向右递归
        return insert_val_search(nums, mid + 1, end, find_val)
    # 查找值和中间值相等，返回下标
    return mid


def main() -> None:
    nums = []
    # 填充 1 - 100
    for i in range(1, 101):
        nums.append(i)
    value = 58
    index = insert_val_search(nums, 0, len(nums) - 1, value)
    if index != -1:
        print(f"找到 {value}, 下标为 {index}")
    else:
        print(f"未找到 {value}")


if __name__ == "__main__":
    main()
