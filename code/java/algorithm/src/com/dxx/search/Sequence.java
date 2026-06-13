package com.dxx.search;

/**
 * 线性查找
 */
public class Sequence {

    public static int sequenceSearch(int[] nums, int num) {
        if (nums == null) {
            return -1;
        }
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == num) {
                return i;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        int value = 8;
        int[] nums = new int[]{2, 5, 1, 7, 8, 16};
        int index = sequenceSearch(nums, value);
        if (index != -1) {
            System.out.printf("%d 在 nums 中的下标为: %d", value, index);
        }
    }
}
