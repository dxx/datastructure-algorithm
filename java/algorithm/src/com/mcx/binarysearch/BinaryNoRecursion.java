package com.mcx.binarysearch;

/**
 * 二分法查找(非递归)
 */
public class BinaryNoRecursion {

    public static int binarySearchNoRecursion(int[] nums, int findVal) {
        int start = 0;
        int end = nums.length - 1;
        while (start <= end) {
            int mid = (start + end) / 2;
            if (findVal < nums[mid]) { // 查找的值在左边
                end = mid - 1;
            } else if (findVal > nums[mid]) { // 查找的值在右边
                start = mid + 1;
            } else {
                // 找到目标值的下标
                return mid;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        int value = 100;
        int[] nums = new int[]{1, 8, 10, 89, 100, 100, 123};
        int index = binarySearchNoRecursion(nums, value);
        if (index != -1) {
            System.out.printf("找到 %d, 下标为 %d\n", value, index);
        } else {
            System.out.printf("未找到 %d\n", value);
        }
    }
}
