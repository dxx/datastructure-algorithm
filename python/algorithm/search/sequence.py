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


def main() -> None:
    value = 8
    nums = [2, 5, 1, 7, 8, 16]
    index = sequence_search(nums, value)
    if index != -1:
        print(f"{value} 在 nums 中的下标为: {index}")


if __name__ == "__main__":
    main()
