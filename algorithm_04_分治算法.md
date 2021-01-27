## 分治算法

>各种语言实现代码：[Go](./golang/algorithm/divide)   [Java](./java/algorithm/src/com/mcx/divide)   [JavaScript](./javascript/algorithm/divide)
>
>默认使用 **Go** 语言实现。

### 简介

分治算法的基本思想是将一个规模为 N 的问题分解为 K 个规模较小的子问题，这些子问题相互独立且与原问题性质相同。求出子问题的解，再找到合适的方法，把它们组合成求整个问题的解法。如果这些子问题还较大，难以解决，可以再把它们分成几个更小的子问题，以此类推，直至可以直接求出解为止。

解题的一般步骤：

1. 分解，将要解决的问题划分成若干规模较小的同类问题。
2. 求解，当子问题划分得足够小时，用较简单的方法解决。
3. 合并，按原问题的要求，将子问题的解逐层合并构成原问题的解。

### 汉诺塔

汉诺塔（又称河内塔）问题是源于印度一个古老传说的益智玩具。大梵天创造世界的时候做了三根金刚石柱子，在一根柱子上从下往上按照大小顺序摞着 64 片黄金圆盘。大梵天命令婆罗门把圆盘从下面开始按大小顺序重新摆放在另一根柱子上。并且规定，在小圆盘上不能放大圆盘，在三根柱子之间一次只能移动一个圆盘。

拆解思路：

假设有 A、B、C 三个柱子，要把 A 柱子上的所有盘放到 C 柱子上。

如果只有一个盘，直接将盘从 A 柱子移动到 C 柱子。

![algorithm_hannotower](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower.png)

如果有两个或以上的盘，将盘分成两份，最下面一个盘一份和上面所有的面盘一份。

![algorithm_hannotower_1](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_1.png)

先将上面所有的盘移动到 B 柱子，然后将下面的盘从 A 柱子移动到 C 柱子，再将 B 柱子上的所有盘移动到 C 柱子上。

![algorithm_hannotower_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_2.png)

![algorithm_hannotower_3](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_3.png)

![algorithm_hannotower_4](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_4.png)

![algorithm_hannotower_5](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_5.png)

![algorithm_hannotower_6](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_6.png)

![algorithm_hannotower_7](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_7.png)

![algorithm_hannotower_8](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_hannotower_8.png)

代码实现：

```go
func hannotower(num int, a, b, c byte) {
    if num == 1 { // 只有一个盘
        fmt.Printf("第 1 个盘从 %s 到 %s\n", string(a), string(c))
        return
    }
    // 先将除了最后一个盘之外的所有盘从 a 移动到 b
    hannotower(num - 1, a, c, b)
    // 最后一个盘从 a 移动到 c
    fmt.Printf("第 %d 个盘从 %s 到 %s\n", num, string(a), string(c))
    // 再将 b 柱子上的所有盘移动到 c 柱子上
    hannotower(num - 1, b, a, c)
}

func main() {
    hannotower(3, 'A', 'B', 'C')
}
```

输出：

```
第 1 个盘从 A 到 C
第 2 个盘从 A 到 B
第 1 个盘从 C 到 B
第 3 个盘从 A 到 C
第 1 个盘从 B 到 A
第 2 个盘从 B 到 C
第 1 个盘从 A 到 C
```
