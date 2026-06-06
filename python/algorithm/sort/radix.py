"""
基数排序
1.创建一个二维数组，数组长度为 10，并初始化 10 个一维数组，每个一维数组的长度为待排序数组的长度
2.遍历待排序的数组，从最低位开始，求出每个元素的个位数作为二维数组的下标，将其放入到二维数组对应的数组中
3.从二维数组中依次取出所有元素放入原数组中
4.重复步骤 2，依次计算个位、十位、百位等，作为下标，直到待排序数组中的最大位数
"""


def radix_sort(nums: list[int] | None) -> None:
    if nums is None:
        return
    if len(nums) == 0:
        return

    max_value = nums[0]  # 最大位的元素
    for num in nums:
        if max_value < num:
            max_value = num
    max_length = len(str(max_value))

    # 桶数组
    # 初始化切片长度
    bucket = [[0 for _ in range(len(nums))] for _ in range(10)]
    order = [0 for _ in range(10)]  # 存放每个桶真实存放数据的长度

    n = 1  # 控制元素的位数
    for _ in range(max_length):
        for num in nums:
            bucket_index = num // n % 10  # 计算桶的下标
            bucket[bucket_index][order[bucket_index]] = num
            order[bucket_index] += 1  # 尾下标 + 1

        num_index = 0
        # 从桶数组中依次取出所有元素放入原数组
        for order_index in range(len(order)):
            bucket_length = order[order_index]
            if bucket_length != 0:
                for j in range(bucket_length):
                    nums[num_index] = bucket[order_index][j]
                    num_index += 1
                order[order_index] = 0  # 重置当前桶下标
        n *= 10


def main() -> None:
    nums = [5, 1, 7, 13, 21, 32, 9, 66, 8, 20]
    print("排序前: " + str(nums))
    radix_sort(nums)
    print("排序后: " + str(nums))


if __name__ == "__main__":
    main()
