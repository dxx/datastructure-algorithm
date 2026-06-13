package com.dxx.recursion;

/**
 * 递归
 * 递归是指在程序运行过程中调用本身的编程技巧
 */
public class Recursion {
    /**
     * 计算阶乘
     */
    public static int factorial(int n) {
        if (n > 0) {
            return n * factorial(n - 1);
        }
        return 1;
    }

    public static void main(String[] args) {
        int res = factorial(5);
        System.out.println(res); // 120
    }
}
