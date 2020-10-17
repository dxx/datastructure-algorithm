package com.mcx.sort;

import java.util.Arrays;

/**
 * 插入排序
 * 1.从第二个元素开始循环，循环到元素末尾，左边分为有序列表，右边分为无序列表
 * 2.将右边无序列表的第一个元素，标记为要插入的值（insertValue），并记录要插入的位置（insertIndex）
 * 3.依次和左边的无序列表中的元素比较，如果顺序颠倒，则将有序列表中当前被比较的元素后移，同时修改 insertIndex
 * 4.最后将 insertValue 插入到 insertIndex 位置
 */
public class Insert {

    public static void insertSort(int[] nums) {
        if (nums == null) {
            return;
        }
        for (int i = 1; i < nums.length; i++) {
            // 要插入的下标
            int insertIndex = i - 1;
            // 保存插入的值
            int insertValue = nums[i];
            // 从小到大
            while (insertIndex >= 0 && nums[insertIndex] > insertValue) {
                // 元素后移
                nums[insertIndex + 1] = nums[insertIndex];
                // 下标前移
                insertIndex--;
            }
            // 优化：判断如果当前位置发生移动，就插入
            if (insertIndex + 1 != i) {
                // 插入
                nums[insertIndex + 1] = insertValue;
            }
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[]{5, 1, 7, 3, 2, 4, 9, 6, 8};
        System.out.printf("排序前: %s\n", Arrays.toString(nums));
        insertSort(nums);
        System.out.printf("排序后: %s\n", Arrays.toString(nums));
    }
}
