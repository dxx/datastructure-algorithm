package main

import "fmt"

// 暴力匹配
// 如果当前字符匹配成功（即 S[i] == P[j]），则 i++，j ++，继续匹配下一个字符
// 如果失配（即 S[i] != P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为 0
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
