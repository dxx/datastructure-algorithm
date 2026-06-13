/**
 * 冒泡排序
 * 1.从当前元素起，向前依次比较每一对相邻元素，若逆序则交换
 * 2.对所有元素均重复以上步骤，直至最后一个元素
 */

function bubbleSort(nums) {
  if (nums == null) {
    return;
  }
  let length = nums.length;
  // 外循环为排序趟数，length 个数进行 length - 1 趟
  for (let i = 0; i < length - 1; i++) {
    // 内循环为每趟比较的次数，第 i 趟比较 length - i 次
    for (let j = length - 1; j > i; j--) {
      // 相邻元素比较比较大小，然后交换位置
      if (nums[j] < nums[j - 1]) {
        swap(nums, j, j - 1);
      }
    }
    console.log("第 " + (i + 1) + " 趟排序结果:" + nums);
  }
}

function optimizeBubbleSort(nums) {
  if (nums == null) {
    return;
  }
  let length = nums.length;
  let isChange = false; // 标记是否发生交换
  for (let i = 0; i < length - 1; i++) {
    for (let j = length - 1; j > i; j--) {
      if (nums[j] < nums[j - 1]) {
        swap(nums, j, j - 1);
        isChange = true; // 发生交换
      }
    }
    console.log("第 " + (i + 1) + " 趟排序结果:" + nums);
    if (!isChange) {
      break; // 跳出循环，终止比较
    } else {
      isChange = false; // 重置
    }
  }
}

function swap(nums, i, j) {
  let temp = nums[i];
  nums[i] = nums[j];
  nums[j] = temp;
}

function main() {
  let nums = [1, 5, 7, 3, 2, 4, 9, 6, 8];
  console.log("交换前: " + nums);
  bubbleSort(nums);
  console.log("交换后: " + nums);

  let nums2 = [1, 5, 7, 3, 2, 4, 9, 6, 8];
  console.log("优化前: " + nums2);
  optimizeBubbleSort(nums2);
  console.log("优化后: " + nums2);
}

main();
