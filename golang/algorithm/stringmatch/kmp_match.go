package main

import "fmt"

// KMP 匹配
// 核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的
// 具体实现就是通过一个 next() 函数实现，函数本身包含了模式串的局部匹配信息

func kmpSearch(str, match string) int {
    next := getNext(match)
    fmt.Println(next)
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
