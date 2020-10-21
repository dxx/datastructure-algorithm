package com.mcx.sort;

import java.util.Arrays;

/**
 * 希尔排序
 * 1.计算出步长 step，step = length / 2
 * 2.从 step 开始，循环到 length
 * 3.将循环开始时的元素和比当前大 step 的元素进行比较
 * 4.发现逆序则进行交换
 */
// 交换法
public class Shell {

    public static void shellSort(int[] nums) {
        if (nums == null) {
            return;
        }
        int length = nums.length;
        // 控制步长
        for (int step = length / 2; step > 0; step /= 2) {
            for (int i = step; i < length; i++) {
                for (int j = i - step; j >= 0 && nums[j] > nums[j + step]; j -= step) {
                    // 前面的数比后面的数大，进行交换
                    swap(nums, j, j + step);
                }
            }
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{5, 1, 7, 3, 2, 4, 9, 6, 8};
        System.out.printf("排序前: %s\n", Arrays.toString(nums));
        shellSort(nums);
        System.out.printf("排序后: %s\n", Arrays.toString(nums));
    }
}
