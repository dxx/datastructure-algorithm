package com.mcx.sort;

import java.util.Arrays;

/**
 * 基数排序
 * 1.创建一个二维数组，数组长度为 10，并初始化 10 个一维数组，每个一维数组的长度为待排序数组的长度
 * 2.遍历待排序的数组，从最低位开始，求出每个元素的个位数作为二维数组的下标，将其放入到二维数组对应的数组中
 * 3.从二维数组中依次取出所有元素放入原数组中
 * 4.重复步骤 2，依次计算个位、十位、百位等，作为下标，直到待排序数组中的最大位数
 */
public class Radix {

    public static void radixSort(int[] nums) {
        if (nums == null) {
            return;
        }

        int max = nums[0]; // 最大位的元素
        for (int num : nums) {
            if (max < num) {
                max = num;
            }
        }
        int maxLength = String.valueOf(max).length();

        int[][] bucket = new int[10][nums.length]; // 桶数组
        int[] order = new int[10]; // 存放每个桶真实存放数据的长度

        int n = 1; // 控制元素的位数
        for (int i = 0; i < maxLength; i++) {
            for (int num : nums) {
                int bucketIndex = num / n % 10; // 计算桶的下标
                bucket[bucketIndex][order[bucketIndex]] = num;
                order[bucketIndex]++; // 尾下标 + 1
            }

            int numIndex = 0;
            // 从桶数组中依次取出所有元素放入原数组
            for (int orderIndex = 0; orderIndex < order.length; orderIndex++) {
                int bucketLength = order[orderIndex];
                if (bucketLength != 0) {
                    for (int j = 0; j < bucketLength; j++) {
                        nums[numIndex] = bucket[orderIndex][j];
                        numIndex++;
                    }
                    order[orderIndex] = 0; // 重置当前桶下标
                }
            }

            n *= 10;
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[]{5, 1, 7, 13, 21, 32, 9, 66, 8, 20};
        System.out.printf("排序前: %s\n", Arrays.toString(nums));
        radixSort(nums);
        System.out.printf("排序后: %s\n", Arrays.toString(nums));
    }
}
