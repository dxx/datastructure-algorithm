## 查找

>各种语言实现代码：[Go](../code/golang/algorithm/search)   [Java](../code/java/algorithm/src/com/dxx/search)   [JavaScript](../code/javascript/algorithm/search)   [TypeScript](../code/typescript/algorithm/search)   [Python](../code/python/algorithm/search)   [Rust](../code/rust/algorithm/src/search)
>
>默认使用 **Python** 语言实现。

### 简介

在一些有序或无序的数据元素中，通过一定的方法找出与给定关键字相同的数据元素的过程叫做查找。也就是根据给定的某个值，在查找表中找出一个关键字等于给定值的记录或数据元素。

### 线性查找

按照给定数据元素的顺序，依次和给定的值比较，满足相等时就返回。

示例：

从元素为 2, 5, 1, 7, 8, 16 的数组中查找 8 ，返回其位置。

代码实现：

```python
def sequence_search(nums: list[int] | None, num: int) -> int:
    if nums is None:
        return -1
    for i, value in enumerate(nums):
        if value == num:
            return i
    return -1
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_sequence_search(self):
        value = 8
        nums = [2, 5, 1, 7, 8, 16]
        index = sequence_search(nums, value)
        if index != -1:
            print(f"{value} 在 nums 中的下标为: {index}")
```

运行：

```shell
❯ python -m unittest test_sequence.Test.test_sequence_search
8 在 nums 中的下标为: 4
```

### 二分法查找

二分法查找适用于数据量较大时，但是数据需要先排好顺序。确定该区间的中间位置 K，将查找的值 T 与 array[k]比较，若相等，查找成功返回此位置；否则确定新的查找区域，继续二分查找。

步骤：

1. 先找到中间值。

2. 将中间值和查找值比较。

   查找值小于中间值, 向左进行递归查找。

   查找值大于中间值, 向右进行递归查找。

   查找值和中间值相等，返回当前下标。

3. 如果查找时，左边的小标大于右边的下标表示未找到，返回 -1。

**注意**：使用二分查找的前提是该数组是有序的。

示例：

从数组 1, 8, 10, 89, 100, 100, 123 中查找 100，返回其下标。

画图分析：

![algorithm_search_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_search_01.png)

代码实现：

```python
def binary_search(nums: list[int], start: int, end: int, find_val: int) -> int:
    if start > end:
        # 表示未找到
        return -1
    mid = (start + end) // 2
    if find_val < nums[mid]:
        # 向左递归
        return binary_search(nums, start, mid - 1, find_val)
    if find_val > nums[mid]:
        # 向右递归
        return binary_search(nums, mid + 1, end, find_val)
    # 查找值和中间值相等，返回下标
    return mid
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_binary_search(self):
        value = 100
        nums = [1, 8, 10, 89, 100, 100, 123]
        index = binary_search(nums, 0, len(nums) - 1, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
```

运行：

```shell
❯ python -m unittest test_binary.Test.test_binary_search
找到 100, 下标为 5
```

### 插值查找

插值查找，有序表的一种查找方式。插值查找是根据查找关键字与查找表中最大最小记录关键字比较后的查找方法。插值查找基于二分法查找，将查找点的选择改进为自适应选择，提高查找效率。**插值查找除要求查找表是顺序存储的有序表外，还要求数据元素的关键字在查找表中均匀分布**，这样就可以按比例插值。

二分法查找中间值下标计算公式为：

mid = (start + end) / 2 = start + (end - start) / 2

插值查找将中间值的下表改为自适应的方式，其计算公式为：

**mid = start + (end - start) * (value - array\[start\]) / (array\[end\] - array\[start\]) **

`start：起始下标`

`end：末尾下标`

`value：要查找的值`

`array：数组`

步骤：

插值查找步骤和二分法查找步骤一样。

示例：

有一个数组，元素为 1, 2, 3, ..., 99, 100，找出 58 所在的下标。

代码实现：

```python
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
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_insert_val_search(self):
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
```

运行：

```shell
❯ python -m unittest test_insert_value.Test.test_insert_val_search
mid: 57
找到 58, 下标为 57
```

可以看到，第一次 mid 直接就计算出值为 57，调用一次就能找到下标。

### 斐波那契查找

斐波那契搜索就是在二分查找的基础上根据斐波那契数列进行分割的。在斐波那契数列找一个等于略大于查找表中元素个数的数 F[n]，将原查找表扩展为长度为 F\[n\](如果要补充元素，则补充重复最后一个元素，直到满足F[n]个元素)，完成后进行斐波那契分割，即 F[n] 个元素分割为前半部分 F[n-1] 个元素，后半部分 F[n-2] 个元素，找出要查找的元素在那一部分并递归，直到找到为止。

斐波那契数列中间值的计算公式为：

**mid = start + F[k - 1] - 1**

`F：斐波那契数列`

示例：

从数组 1, 8, 10, 89, 100, 100, 123 中查找 100，返回其下标。

代码实现：

```python
MAX_SIZE = 20  # 斐波那契数组长度


# 获取斐波那契数列
def fibonacci() -> list[int]:
    fib = [0 for _ in range(MAX_SIZE)]
    fib[0] = 1
    fib[1] = 1
    for i in range(2, MAX_SIZE):
        fib[i] = fib[i - 1] + fib[i - 2]
    return fib


# 斐波那契查找
# 中间值计算公式: mid = start + F[k - 1] - 1
def fibonacci_search(nums: list[int], find_val: int) -> int:
    start = 0
    end = len(nums) - 1
    k = 0  # 斐波那契分割数的下标
    f = fibonacci()  # 斐波那契数列

    # 获取斐波那契分割数的下标
    while end > f[k] - 1:
        k += 1
    # 创建一个切片，长度为斐波那契分割数的值
    dst_length = f[k]
    # 将要查找的数组复制到 dst 中
    dst = nums[:]
    # 将最后一个元素填充到 dst 元素为 0 的位置
    # [1, 8, 10, 89, 100, 100, 123] => [1 8 10 89 100 100 123 123]
    for _ in range(end + 1, dst_length):
        dst.append(nums[end])

    # 使用循环来寻找 findVal
    while start <= end:
        # 中间值
        mid = start + f[k - 1] - 1
        if find_val < dst[mid]:  # 向左查找
            end = mid - 1
            # 全部元素 = 前面的元素 + 后边元素
            # 即f[k] = f[k-1] + f[k-2]
            # 因为前面有 f[k-1]个元素,所以可以继续拆分 f[k-1] = f[k-2] + f[k-3]
            # 即 在 f[k-1] 的前面继续查找 k--
            # 即下次循环 mid = f[k-1-1]-1
            k -= 1
        elif find_val > dst[mid]:  # 向右查找
            start = mid + 1
            # 全部元素 = 前面的元素 + 后边元素
            # f[k] = f[k-1] + f[k-2]
            # 因为后面我们有f[k-2] 所以可以继续拆分 f[k-1] = f[k-3] + f[k-4]
            # 即在f[k-2] 的前面进行查找 k -=2
            # 即下次循环 mid = f[k - 1 - 2] - 1
            k -= 2
        else:
            if mid <= end:
                return mid  # 若相等则说明 mid 即为查找到的位置
            return end  # 若 mid > end 则说明是扩展的数值，返回 end
    return -1
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_fibonacci_search(self):
        value = 100
        nums = [1, 8, 10, 89, 100, 100, 123]
        index = fibonacci_search(nums, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
```

运行：

```shell
❯ python -m unittest test_fibonacci.Test.test_fibonacci_search
找到 100, 下标为 4
```

## 二分法查找(非递归)

>各种语言实现代码：[Go](../code/golang/algorithm/binarysearch)   [Java](../code/java/algorithm/src/com/dxx/binarysearch)   [JavaScript](../code/javascript/algorithm/binarysearch)   [TypeScript](../code/typescript/algorithm/binarysearch)   [Python](../code/python/algorithm/binarysearch)   [Rust](../code/rust/algorithm/src/binary_search)
>
>默认使用 **Python** 语言实现。

### 简介

二分法查找也可以使用非递归的方式来实现，思想和二分法查找一样。确定该区间的中间位置 K，将查找的值 T 与 array[k] 比较，若相等，查找成功返回此位置；否则确定新的查找区域，继续二分查找。

> 注意二分法查找，需要数组是有序的。

### 实现

代码实现：

```python
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
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_binary_search_no_recursion(self):
        value = 100
        nums = [1, 8, 10, 89, 100, 100, 123]
        index = binary_search_no_recursion(nums, value)
        if index != -1:
            print(f"找到 {value}, 下标为 {index}")
        else:
            print(f"未找到 {value}")
```

运行：

```shell
❯ python -m unittest test_binary_no_recursion.Test.test_binary_search_no_recursion
找到 100, 下标为 5
```
