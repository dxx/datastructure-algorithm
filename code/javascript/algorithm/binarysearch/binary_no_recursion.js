/**
 * 二分法查找(非递归)
 */

function binarySearchNoRecursion(nums, findVal) {
  let start = 0;
  let end = nums.length - 1;
  while (start <= end) {
    let mid = (start + end) / 2;
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

function main() {
  let value = 100;
  let nums = [1, 8, 10, 89, 100, 100, 123];
  let index = binarySearchNoRecursion(nums, value);
  if (index != -1) {
    console.log("找到 " + value + ", 下标为 " + index);
  } else {
    console.log("未找到 " + value);
  }
}

main();
