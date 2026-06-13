"""
希尔排序
1.计算出步长 step，step = length / 2
2.从 step 开始，循环到 length
3.将循环开始时的元素和比当前大 step 的元素进行比较
4.发现逆序则进行交换
"""


def shell_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    length = len(nums)
    # 控制步长
    step = length // 2
    while step > 0:
        for i in range(step, length):
            j = i - step
            while j >= 0 and nums[j] > nums[j + step]:
                # 前面的数比后面的数大，进行交换
                swap(nums, j, j + step)
                j -= step
        step //= 2


def swap(nums: list[int], i: int, j: int) -> None:
    temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
