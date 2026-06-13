## 递归

>各种语言实现代码：[Go](../code/golang/algorithm/recursion)   [Java](./java/algorithm/src/com/dxx/recursion)   [JavaScript](./javascript/algorithm/recursion)   [TypeScript](./typescript/algorithm/recursion)   [Python](./python/algorithm/recursion)   [Rust](./rust/algorithm/src/recursion)
>
>默认使用 **Python** 语言实现。

### 简介

递归是指在程序运行过程中调用本身的编程技巧。递归通常把一个大型复杂的问题层层转化为一个与原问题相似的规模较小的问题来求解，递归只需少量的代码就可描述出解题过程所需要的多次重复计算，大大地减少了程序的代码量。一般来说，递归需要有边界条件、递归前进代码块和递归返回代码块。当边界条件不满足时，递归前进，当边界条件满足时，递归返回。

### 示例

#### 计算阶乘

```python
def factorial(n: int) -> int:
    if n > 0:
        return n * factorial(n - 1)
    return 1
```

测试代码：

```python
class Test(unittest.TestCase):

    def test_factorial(self):
        res = factorial(5)
        print(res)  # 120
```

### 迷宫回溯

下图中黑色部分表示墙壁，白色部分表示可以走的通路。

![algorithm_migong](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_migong.png)

假设一个二维数组中 0 表示通道，1 表示为墙壁，给出一个起点，求起点到终点的最短路径。

思路分析：

使用递归算法，判断小球是否能够继续往下走，如果可以走就继续，不可用走就退回寻找其它通道。

需要指定一个走路策略即走路的方向是从下 -> 右 -> 上 -> 左 ，还是从上 -> 右 -> 下 -> 左。

测试每一种走路策略，对比找出最少的步骤就是最短路径。

假设使用下 -> 右 -> 上 -> 左的策略，实现函数如下：

```python
# 终点
X = 6
Y = 6


# 如果小球走到 6,6 的位置，表示已经走完
# 假设 0-未走过，1-墙，2-走过，3-走不通
# 规定策略：下 -> 右 -> 上 -> 左
# 如果走不通再回溯
def walk(mi_gong_map: list[list[int]], i: int, j: int) -> bool:
    # 到达终点
    if mi_gong_map[X][Y] == 2:
        return True
    if mi_gong_map[i][j] == 0:
        # 设置为已走过
        mi_gong_map[i][j] = 2
        # 下 -> 右 -> 上 -> 左
        if walk(mi_gong_map, i + 1, j):  # 向下走
            return True
        if walk(mi_gong_map, i, j + 1):  # 向右走
            return True
        if walk(mi_gong_map, i - 1, j):  # 向上走
            return True
        if walk(mi_gong_map, i, j - 1):  # 向左走
            return True
        # 走过但路不通
        mi_gong_map[i][j] = 3
    # 可能为 1, 2 , 3
    return False
```

测试代码：

```python
def print_map(mi_gong_map: list[list[int]]) -> None:
    result = ""
    for row in mi_gong_map:
        for value in row:
            result += " " + str(value)
        result += "\n"
    print(result)


class Test(unittest.TestCase):

    def test_walk(self):
        # 初始化地图，0 表示通道，1 表示墙
        mi_gong_map = [
            [1, 1, 1, 1, 1, 1, 1, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 0, 0, 0, 0, 0, 0, 1],
            [1, 1, 1, 1, 1, 1, 1, 1],
        ]

        print("探路之前:")
        print_map(mi_gong_map)
        # 开始探路, 起点为1,1
        walk(mi_gong_map, 1, 1)
        print("探路之后:")
        print_map(mi_gong_map)
```

运行：

```shell
❯ python -m unittest test_mi_gong.Test.test_walk 
探路之前:
 1 1 1 1 1 1 1 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 1 1 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 1 1 1 1 1 1 1

探路之后:
 1 1 1 1 1 1 1 1
 1 2 0 0 0 0 0 1
 1 2 2 2 0 0 0 1
 1 1 1 2 0 0 0 1
 1 0 0 2 0 0 0 1
 1 0 0 2 0 0 0 1
 1 0 0 2 2 2 2 1
 1 1 1 1 1 1 1 1
```

改变走路策略为从上 -> 右 -> 下 -> 左：

```python
# 规定策略：上 -> 右 -> 下 -> 左
def walk2(mi_gong_map: list[list[int]], i: int, j: int) -> bool:
    # 到达终点
    if mi_gong_map[X][Y] == 2:
        return True
    if mi_gong_map[i][j] == 0:
        # 设置为已走过
        mi_gong_map[i][j] = 2
        # 上 -> 右 -> 下 -> 左
        if walk2(mi_gong_map, i - 1, j):  # 向上走
            return True
        if walk2(mi_gong_map, i, j + 1):  # 向右走
            return True
        if walk2(mi_gong_map, i + 1, j):  # 向下走
            return True
        if walk2(mi_gong_map, i, j - 1):  # 向左走
            return True
        # 走过但路不通
        mi_gong_map[i][j] = 3
    # 可能为 1, 2 , 3
    return False
```

运行测试代码：

```shell
❯ python -m unittest test_mi_gong.Test.test_walk2
探路之前:
 1 1 1 1 1 1 1 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 1 1 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 0 0 0 0 0 0 1
 1 1 1 1 1 1 1 1

探路之后:
 1 1 1 1 1 1 1 1
 1 2 2 2 2 2 2 1
 1 0 0 0 0 0 2 1
 1 1 1 0 0 0 2 1
 1 0 0 0 0 0 2 1
 1 0 0 0 0 0 2 1
 1 0 0 0 0 0 2 1
 1 1 1 1 1 1 1 1
```

求最短路径时，设定几种走路策略，再循环调用走路函数，将每一步走过的结果保存到切片中，判断出包含最少元素的切片，该切片中的元素连成的点就最短路径。

### 八皇后

在 8×8 格的国际象棋上摆放 8 个皇后，使其不能互相攻击，即任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。

![algorithm_eight_queen](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_eight_queen.png)

思路分析：

先将第一个皇后摆放在第一行第一列，接着将第二个皇后先摆在第二行第一列，如果和之前摆放过的皇后冲突再摆放在第二行第二列，还是冲突继续在摆放在下一列，依次类推，整个过程递归进行，当最后一个皇后摆放完成后，回溯到上一个皇后，继续摆放。

使用一个一维数组 positions 存储象棋摆放的位置，下标表示第几个象棋，数组的值表示第几列。将当前皇后和之前摆放过的皇后依次比较，摆放第 n 个皇后时，**positions[n] == positions[i]** 表示在同一列，**|n - i| == |positions[n] - positions[i]|** 表示在同一斜线。

实现如下：

```python
class EightQueen:
    def __init__(self) -> None:
        self.positions = [0 for _ in range(8)]

    def put_queen(self, n: int) -> None:
        # 最后一个皇后已经放置完成
        if n == len(self.positions):
            # 打印当前摆放的位置
            print(self.positions)
            return
        for i in range(len(self.positions)):
            # i=0 时，假设当前皇后可以放在第一列
            # 如果不能放，将进行下一次循环，当前皇后放在下一个位置
            self.positions[n] = i
            # 判断是否可以放
            if self.is_can_put(n):
                # 放置下一个皇后
                self.put_queen(n + 1)

    # 判断当前皇后是否和已经摆放过的皇后冲突
    def is_can_put(self, n: int) -> bool:
        positions = self.positions
        for i in range(n):
            # positions[n] == positions[i] 表示在同一列
            # abs(n - i) == abs(positions[n] - positions[i]) 表示同一斜线
            if positions[n] == positions[i] or abs(n - i) == abs(positions[n] - positions[i]):
                return False
        return True
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_eight_queen(self):
        eight_queen = EightQueen()
        eight_queen.put_queen(0)
```

运行：

```shell
❯ python -m unittest test_eight_queen.Test.test_eight_queen
[0, 4, 7, 5, 2, 6, 1, 3]
[0, 5, 7, 2, 6, 3, 1, 4]
[0, 6, 3, 5, 7, 1, 4, 2]
[0, 6, 4, 7, 1, 3, 5, 2]
[1, 3, 5, 7, 2, 0, 6, 4]
[1, 4, 6, 0, 2, 7, 5, 3]
[1, 4, 6, 3, 0, 7, 5, 2]
[1, 5, 0, 6, 3, 7, 2, 4]
[1, 5, 7, 2, 0, 3, 6, 4]
[1, 6, 2, 5, 7, 4, 0, 3]
[1, 6, 4, 7, 0, 3, 5, 2]
[1, 7, 5, 0, 2, 4, 6, 3]
[2, 0, 6, 4, 7, 1, 3, 5]
[2, 4, 1, 7, 0, 6, 3, 5]
[2, 4, 1, 7, 5, 3, 6, 0]
[2, 4, 6, 0, 3, 1, 7, 5]
[2, 4, 7, 3, 0, 6, 1, 5]
[2, 5, 1, 4, 7, 0, 6, 3]
[2, 5, 1, 6, 0, 3, 7, 4]
[2, 5, 1, 6, 4, 0, 7, 3]
[2, 5, 3, 0, 7, 4, 6, 1]
[2, 5, 3, 1, 7, 4, 6, 0]
[2, 5, 7, 0, 3, 6, 4, 1]
[2, 5, 7, 0, 4, 6, 1, 3]
[2, 5, 7, 1, 3, 0, 6, 4]
[2, 6, 1, 7, 4, 0, 3, 5]
[2, 6, 1, 7, 5, 3, 0, 4]
[2, 7, 3, 6, 0, 5, 1, 4]
[3, 0, 4, 7, 1, 6, 2, 5]
[3, 0, 4, 7, 5, 2, 6, 1]
[3, 1, 4, 7, 5, 0, 2, 6]
[3, 1, 6, 2, 5, 7, 0, 4]
[3, 1, 6, 2, 5, 7, 4, 0]
[3, 1, 6, 4, 0, 7, 5, 2]
[3, 1, 7, 4, 6, 0, 2, 5]
[3, 1, 7, 5, 0, 2, 4, 6]
[3, 5, 0, 4, 1, 7, 2, 6]
[3, 5, 7, 1, 6, 0, 2, 4]
[3, 5, 7, 2, 0, 6, 4, 1]
[3, 6, 0, 7, 4, 1, 5, 2]
[3, 6, 2, 7, 1, 4, 0, 5]
[3, 6, 4, 1, 5, 0, 2, 7]
[3, 6, 4, 2, 0, 5, 7, 1]
[3, 7, 0, 2, 5, 1, 6, 4]
[3, 7, 0, 4, 6, 1, 5, 2]
[3, 7, 4, 2, 0, 6, 1, 5]
[4, 0, 3, 5, 7, 1, 6, 2]
[4, 0, 7, 3, 1, 6, 2, 5]
[4, 0, 7, 5, 2, 6, 1, 3]
[4, 1, 3, 5, 7, 2, 0, 6]
[4, 1, 3, 6, 2, 7, 5, 0]
[4, 1, 5, 0, 6, 3, 7, 2]
[4, 1, 7, 0, 3, 6, 2, 5]
[4, 2, 0, 5, 7, 1, 3, 6]
[4, 2, 0, 6, 1, 7, 5, 3]
[4, 2, 7, 3, 6, 0, 5, 1]
[4, 6, 0, 2, 7, 5, 3, 1]
[4, 6, 0, 3, 1, 7, 5, 2]
[4, 6, 1, 3, 7, 0, 2, 5]
[4, 6, 1, 5, 2, 0, 3, 7]
[4, 6, 1, 5, 2, 0, 7, 3]
[4, 6, 3, 0, 2, 7, 5, 1]
[4, 7, 3, 0, 2, 5, 1, 6]
[4, 7, 3, 0, 6, 1, 5, 2]
[5, 0, 4, 1, 7, 2, 6, 3]
[5, 1, 6, 0, 2, 4, 7, 3]
[5, 1, 6, 0, 3, 7, 4, 2]
[5, 2, 0, 6, 4, 7, 1, 3]
[5, 2, 0, 7, 3, 1, 6, 4]
[5, 2, 0, 7, 4, 1, 3, 6]
[5, 2, 4, 6, 0, 3, 1, 7]
[5, 2, 4, 7, 0, 3, 1, 6]
[5, 2, 6, 1, 3, 7, 0, 4]
[5, 2, 6, 1, 7, 4, 0, 3]
[5, 2, 6, 3, 0, 7, 1, 4]
[5, 3, 0, 4, 7, 1, 6, 2]
[5, 3, 1, 7, 4, 6, 0, 2]
[5, 3, 6, 0, 2, 4, 1, 7]
[5, 3, 6, 0, 7, 1, 4, 2]
[5, 7, 1, 3, 0, 6, 4, 2]
[6, 0, 2, 7, 5, 3, 1, 4]
[6, 1, 3, 0, 7, 4, 2, 5]
[6, 1, 5, 2, 0, 3, 7, 4]
[6, 2, 0, 5, 7, 4, 1, 3]
[6, 2, 7, 1, 4, 0, 5, 3]
[6, 3, 1, 4, 7, 0, 2, 5]
[6, 3, 1, 7, 5, 0, 2, 4]
[6, 4, 2, 0, 5, 7, 1, 3]
[7, 1, 3, 0, 6, 4, 2, 5]
[7, 1, 4, 2, 0, 6, 3, 5]
[7, 2, 0, 5, 1, 4, 6, 3]
[7, 3, 0, 2, 5, 1, 6, 4]
```
