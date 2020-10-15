package com.mcx.sort;

import java.util.Arrays;

/**
 * 选择排序
 * 1.假定第一元素为最大或最小的元素
 * 2.找出最大或最小的元素的小标，循环 length - 1 次
 * 3.每次循环完成后将最大值或最小值和本次循环的第一个元素交换
 */
public class Select {

    public static void selectSort(int[] nums) {
        if (nums == null) {
            return;
        }
        int length = nums.length;
        for (int i = 0; i < length - 1; i++) {
            // 记录最小值的下标
            int minIndex = i;
            for (int j = i + 1; j < length; j++) {
                if (nums[minIndex] > nums[j]) {
                    // 修改最小值下标
                    minIndex = j;
                }
            }
            // 优化：判断是否需要交换
            if (minIndex != i) {
                swap(nums, i, minIndex);
            }
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{3, 5, 7, 1, 2, 4, 9, 6, 8};
        System.out.printf("排序前: %s\n", Arrays.toString(nums));
        selectSort(nums);
        System.out.printf("排序后: %s\n", Arrays.toString(nums));
    }
}
