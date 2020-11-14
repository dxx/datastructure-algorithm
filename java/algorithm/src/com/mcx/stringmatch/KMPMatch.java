package com.mcx.stringmatch;

/**
 * KMP 匹配
 * 核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的
 * 具体实现就是通过一个 next() 函数实现，函数本身包含了模式串的局部匹配信息
 */
public class KMPMatch {

    public static int kmpSearch(String str, String match) {
        int[] next = getNext(match);

        for (int i = 0, j = 0; i < str.length(); i++) {
            // 算法核心点
            while (j > 0 && str.charAt(i) != match.charAt(j)) {
                // 根据部分匹配表，更新 j
                j = next[j - 1];
            }

            if (str.charAt(i) == match.charAt(j)) {
                j++;
            }
            // 判断是否找到了
            if (j == match.length()) {
                return i - (j - 1);
            }
        }
        return -1;
    }

    public static int[] getNext(String match) {
        int[] next = new int[match.length()];

        // 第一个字符的值为 0
        next[0] = 0;

        for (int i = 1, j = 0; i < match.length(); i++) {

            // 核心，比较直到相等
            while (j > 0 && match.charAt(i) != match.charAt(j)) {
                // 更新 j
                j = next[j - 1];
            }

            // 相等
            if (match.charAt(i) == match.charAt(j)) {
                j++;
            }
            next[i] = j;
        }
        return next;
    }

    public static void main(String[] args) {
        String str = "CBC DCABCABABCABD BBCCA";
        String match = "ABCABD";
        int index = kmpSearch(str, match);
        System.out.printf("%s 在 %s 中的位置为 %d", match, str, index);
    }

}
