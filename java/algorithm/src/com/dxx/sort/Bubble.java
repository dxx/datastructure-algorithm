package com.dxx.sort;

import java.util.Arrays;

/**
 * 冒泡排序
 * 1.从当前元素起，向前依次比较每一对相邻元素，若逆序则交换
 * 2.对所有元素均重复以上步骤，直至最后一个元素
 */
public class Bubble {

    public static void bubbleSort(int[] nums) {
        if (nums == null) {
            return;
        }
        int length = nums.length;
        // 外循环为排序趟数，length 个数进行 length - 1 趟
        for (int i = 0; i < length - 1; i++) {
            // 内循环为每趟比较的次数，第 i 趟比较 length - i 次
            for (int j = length - 1; j > i; j--) {
                // 相邻元素比较比较大小，然后交换位置
                if (nums[j] < nums[j - 1]) {
                    swap(nums, j, j - 1);
                }
            }
            System.out.printf("第 %d 趟排序结果:%s\n", i + 1, Arrays.toString(nums));
        }
    }

    /**
     * 优化
     * 在某次循环中，如果发现没有发生交换，则终止循环
     */
    public static void optimizeBubbleSort(int[] nums) {
        if (nums == null) {
            return;
        }
        int length = nums.length;
        boolean isChange = false; // 标记是否发生交换
        for (int i = 0; i < length - 1; i++) {
            for (int j = length - 1; j > i; j--) {
                if (nums[j] < nums[j - 1]) {
                    swap(nums, j, j - 1);
                    isChange = true; // 发生交换
                }
            }
            System.out.printf("第 %d 趟排序结果:%s\n", i + 1, Arrays.toString(nums));
            if (!isChange) {
                break; // 跳出循环，终止比较
            } else {
                isChange = false; // 重置
            }
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 5, 7, 3, 2, 4, 9, 6, 8};
        System.out.printf("交换前: %s\n", Arrays.toString(nums));
        bubbleSort(nums);
        System.out.printf("交换后: %s\n", Arrays.toString(nums));

        int[] nums2 = new int[]{1, 5, 7, 3, 2, 4, 9, 6, 8};
        System.out.printf("优化前: %s\n", Arrays.toString(nums2));
        optimizeBubbleSort(nums2);
        System.out.printf("优化后: %s\n",  Arrays.toString(nums2));
    }
}
