/**
 * 插值查找
 * 基于二分法查找，步骤和二分法查找一样
 * 中间下标计算公式: mid = start + (end - start) * (value - array[start]) / (array[end] - array[start])
 */

function insertValSearch(nums, start, end, findVal) {
  if (start > end) {
    return -1;
  }
  // 根据 findVal 自适应计算中间下标
  let mid = start + (end - start) * (findVal - nums[start]) / (nums[end] - nums[start]);
  console.log("mid: " + mid);
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

function main() {
  let nums = [];
  // 填充 1 - 100
  for (let i = 1; i <= 100; i++) {
    nums.push(i);
  }
  let value = 58;
  let index = insertValSearch(nums, 0, nums.length - 1, value);
  if (index != -1) {
    console.log("找到 " + value + ", 下标为 " + index);
  } else {
    console.log("未找到 " + value);
  }
}

main();
