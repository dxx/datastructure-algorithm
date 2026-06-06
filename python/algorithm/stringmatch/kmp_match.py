"""
KMP 匹配
核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的
具体实现就是通过一个 next() 函数实现，函数本身包含了模式串的局部匹配信息
"""


def kmp_search(string: str, match: str) -> int:
    next_values = get_next(match)

    j = 0
    for i in range(len(string)):
        # 算法核心点
        while j > 0 and string[i] != match[j]:
            # 根据部分匹配表，更新 j
            j = next_values[j - 1]

        if string[i] == match[j]:
            j += 1
        # 判断是否找到了
        if j == len(match):
            return i - (j - 1)
    return -1


def get_next(match: str) -> list[int]:
    next_values = [0 for _ in range(len(match))]
    # 第一个字符的值为 0
    next_values[0] = 0

    j = 0
    for i in range(1, len(match)):
        # 核心，比较直到相等
        while j > 0 and match[i] != match[j]:
            # 更新 j
            j = next_values[j - 1]

        # 相等
        if match[i] == match[j]:
            j += 1
        next_values[i] = j
    return next_values


def main() -> None:
    string = "CBC DCABCABABCABD BBCCA"
    match = "ABCABD"
    index = kmp_search(string, match)
    print(match + " 在 " + string + " 中的位置为 " + str(index))


if __name__ == "__main__":
    main()
