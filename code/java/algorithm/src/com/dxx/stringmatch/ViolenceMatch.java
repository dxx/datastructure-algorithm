package com.dxx.stringmatch;

/**
 * 暴力匹配
 * 如果当前字符匹配成功（即 S[i] == P[j]），则 i++，j ++，继续匹配下一个字符
 * 如果失配（即 S[i] != P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为 0
 */
public class ViolenceMatch {

    public static int violenceSearch(String str, String match) {
        byte[] strBytes = str.getBytes();
        byte[] matchBytes = match.getBytes();

        int strLength = strBytes.length;
        int matchLength = matchBytes.length;

        if (strLength < matchLength) {
            return -1;
        }

        int i = 0, j= 0;
        // 当字符串长度并且子字符串长度超出范围
        while (i < strLength && j < matchLength) {
            // 相等继续匹配
            if (strBytes[i] == matchBytes[j]) {
                i++;
                j++;
            } else {
                // i 重置到上一次第一个相等字符的位置 + 1
                i = i - (j - 1);
                // j 重置为 0
                j = 0;
            }
        }

        if (j == matchLength) { // 匹配到，然后返回索引
            return i - j;
        }
        return -1;
    }

    public static void main(String[] args) {
        String str = "CBC DCABCABABCABD BBCCA";
        String match = "ABCABD";
        int index = violenceSearch(str, match);
        System.out.printf("%s 在 %s 中的位置为 %d", match, str, index);
    }
}
