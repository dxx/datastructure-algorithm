## KMP 算法

>各种语言实现代码：[Go](./golang/algorithm/stringmatch)   Java(待实现)   JavaScript(待实现)
>
>默认使用 **Go** 语言实现。

### 简介

KMP 算法是一种改进的字符串匹配算法，KMP 算法是由 D.E.Knuth，J.H.Morris 和 V.R.Pratt 三人提出的，因此人们称它为克努特—莫里斯—普拉特操作（简称 KMP 算法）。KMP 算法的核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的。具体实现就是通过一个 next() 函数实现，函数本身包含了模式串的局部匹配信息。

### 暴力匹配

有一个文本串 S 和一个模式串 P，现在要查找 P 在 S 中的位置。如果用暴力匹配的思路，并假设现在文本串 S 匹配到 i 位置，模式串 P 匹配到 j 位置，则有：

* 如果当前字符匹配成功（即 S[i] == P[j]），则 i++，j ++，继续匹配下一个字符。
* 如果失配（即 S[i] != P[j]），重置 i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回退，j 被置为 0。

代码如下：

```go
func violenceSearch(str, match string) int {
    strBytes := []byte(str)
    matchBytes := []byte(match)

    strLength := len(strBytes)
    matchLength := len(matchBytes)

    if strLength < matchLength {
        return -1
    }

    i, j := 0, 0
    // 当字符串长度并且子字符串长度超出范围
    for i < strLength && j < matchLength {
        // 相等继续匹配
        if strBytes[i] == matchBytes[j] {
            i++
            j++
        } else {
            // i 重置到上一次第一个相等字符的位置 + 1
            i = i - (j - 1)
            // j 重置为 0
            j = 0
        }
    }

    if j == matchLength { // 匹配到，然后返回索引
        return i - j
    }
    return -1
}

func main() {
    str := "CBC DCABCABABCABD BBCCA"
    match := "ABCABD"
    index := violenceSearch(str, match)
    fmt.Printf("%s 在 %s 中的位置为 %d", match, str, index)
}
```

输出：

```
ABCABD 在 CBC DCABCABABCABD BBCCA 中的位置为 11
```

### KMP匹配

上述暴力匹配时，每次匹配不成功都要回溯到和第一个相同字符的下一个位置。即 ABCBCD 中的 A 字符匹配到strBytes[6] 位置的 A 时，继续匹配 B，直到匹配 ABCABD 中的 D 时，D 和 strBytes[12] 不相等，进行回溯到 strBytes[7]，显然 strBytes[7] 和 ABCABD 中的第一字符 A 不相等，只有 strBytes[9] 才和 A 相等，而 KMP 算法就可以让下标从 9 开始匹配，尽量让下一次移动到有效位置进行匹配。

图解分析：

首先比较文本串和模式串的第一字符。

![algorithm_kmp_1](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_1.png)

不相等，使用后一个字符比较。

![algorithm_kmp_2](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_2.png)

还是不相等，继续后移一位比较，直到模式串的第一字符和文本串中的字符相等。

![algorithm_kmp_3](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_3.png)

接着比较文本串和模式串的下一个字符。

![algorithm_kmp_4](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_4.png)

直到遇到文本串中有一个字符和模式串的当前字符串不相等。

![algorithm_kmp_5](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_5.png)

如果是暴力匹配，则将文本串中和模式串中第一字符相等的下一个字符开始继续比较。

![algorithm_kmp_6](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_6.png)

这是不明智的，因为此时 B C 已经比较过了。KMP 算法的想法是，设法利用已知信息，不要把"搜索位置"移回已经比较过的位置，继续把它向后移，这样就提高了效率。

模式串 `ABCABD` 计算出部分匹配表，匹配表如下：

| 字符   | A    | B    | C    | A    | B    | D    |
| ------ | ---- | ---- | ---- | ---- | ---- | ---- |
| 匹配值 | 0    | 0    | 0    | 1    | 2    | 0    |

> 部分匹配值计算过程见后面。

当 A 和 D 不匹配时，ABCAB 是匹配的，根据匹配表，最后一个 B 对应的值为 2，按照下面的公式：

移动位数 = 已匹配的字符数 - 最后一个字符对应的匹配值

5 - 2 = 3，故此时将模式串后移 3 位，按照之前的逻辑继续匹配，直到 A 和 C 不匹配。

![algorithm_kmp_7](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_7.png)

按照移动位数计算公式，此时后移 2 位，整个模式串匹配成功。

![algorithm_kmp_8](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/algorithm_kmp_8.png)

**部分匹配值就是前缀和后缀的最长共有元素的长度**。假设一个字符串 "hello"，它的前缀有 h、he、hel、hell，它的后缀有 ello、llo、lo、o。

以上述模式字符串 ABCAB 为例：

* A 没有前缀和后缀，共有元素长度为 0。
* AB 的前缀有 A，后缀有 B，共有元素长度为 0。
* ABC 的前缀有 A、AB，后缀有 BC、C，共有元素长度为 0。
* ABCA 的前缀有 A、AB、ABC，后缀有 BCA、CA、A，共有元素长度为 1。
* ABCAB 的前缀有 A、AB、ABC、ABCA，后缀有 BCAB、CAB、AB、B，共有元素长度为 2。
* ABCABD 的前缀有 A、AB、ABC、ABCA、ABCAB，后缀有 BCABD、CABD、ABD、BD、D，共有元素长度为 0。

所以 ABCABD 中每个字符对于的匹配值分别为 0、0、0、1、2、0。

代码实现：

```go
func kmpSearch(str, match string) int {
    next := getNext(match)

    for i, j := 0, 0; i < len(str); i++ {
        // 算法核心点
        for j > 0 && str[i] != match[j] {
            // 根据部分匹配表，更新 j
            j = next[j - 1]
        }

        if str[i] == match[j] {
            j++
        }
        // 判断是否找到了
        if j == len(match) {
            return i - (j - 1)
        }
    }
    return -1
}

// 计算出匹配的部分信息
func getNext(match string) []int {
    next := make([]int, len(match))

    // 第一个字符的值为 0
    next[0] = 0

    for i, j := 1, 0; i < len(match); i++ {

        // 核心，比较直到相等
        for j > 0 && match[i] != match[j] {
            // 更新 j
            j = next[j - 1]
        }

        // 相等
        if match[i] == match[j] {
            j++
        }
        next[i] = j
    }
    return next
}

func main() {
    str := "CBC DCABCABABCABD BBCCA"
    match := "ABCABD"
    index := kmpSearch(str, match)
    fmt.Printf("%s 在 %s 中的位置为 %d", match, str, index)
}
```

输出：

```
ABCABD 在 CBC DCABCABABCABD BBCCA 中的位置为 11
```
