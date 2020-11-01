package com.mcx.search;

/**
 * 二分查找
 * 1.先找到中间值
 * 2.将中间值和查找值比较
 *   查找值小于中间值, 向左进行递归查找
 *   查找值大于中间值, 向右进行递归查找
 *   查找值和中间值相等，返回当前下标
 * 3.如果查找时，左边的小标大于右边的下标表示未找到，返回 -1
 * 注意：使用二分查找的前提是该数组是有序的
 */
public class Binary {

    public static int binarySearch(int[] nums, int start, int end, int findVal) {
        if (start > end) {
            // 表示未找到
            return -1;
        }
        int mid = (start + end) / 2;
        if (findVal < nums[mid]) {
            // 向左递归
            return binarySearch(nums, start, mid - 1, findVal);
        } else if (findVal > nums[mid]) {
            // 向右递归
            return binarySearch(nums, mid + 1, end, findVal);
        }
        // 查找值和中间值相等，返回下标
        return mid;
    }

    public static void main(String[] args) {
        int value = 100;
        int[] nums = new int[]{1, 8, 10, 89, 100, 100, 123};
        int index = binarySearch(nums, 0, nums.length, value);
        if (index != -1) {
            System.out.printf("找到 %d, 下标为 %d\n", value, index);
        } else {
            System.out.printf("未找到 %d\n", value);
        }
    }
}
