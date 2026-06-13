package com.dxx.search;

/**
 * 插值查找
 * 基于二分法查找，步骤和二分法查找一样
 * 中间下标计算公式: mid = start + (end - start) * (value - array[start]) / (array[end] - array[start])
 */
public class InsertValue {

    public static int insertValSearch(int[] nums, int start, int end, int findVal) {
        if (start > end) {
            return -1;
        }
        // 根据 findVal 自适应计算中间下标
        int mid = start + (end - start) * (findVal - nums[start]) / (nums[end] - nums[start]);
        System.out.printf("mid: %d\n", mid);
        if (findVal < nums[mid]) {
            // 向左递归
            return insertValSearch(nums, start, mid - 1, findVal);
        } else if (findVal > nums[mid]) {
            // 向右递归
            return insertValSearch(nums, mid + 1, end, findVal);
        } else {
            // 查找值和中间值相等，返回下标
            return mid;
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[100];
        // 填充 1 - 100
        for (int i = 1; i <= 100; i++) {
            nums[i - 1] = i;
        }
        int value = 58;
        int index = insertValSearch(nums, 0, nums.length - 1, value);
        if (index != -1) {
            System.out.printf("找到 %d, 下标为 %d\n", value, index);
        } else {
            System.out.printf("未找到 %d\n", value);
        }
    }
}
