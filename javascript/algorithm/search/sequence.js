/**
 * 线性查找
 */

function sequenceSearch(nums, num) {
  if (nums == null) {
    return -1;
  }
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] == num) {
      return i;
    }
  }
  return -1;
  }
  
  function main() {
    let value = 8;
    let nums = [2, 5, 1, 7, 8, 16];
    let index = sequenceSearch(nums, value);
    if (index != -1) {
      console.log(value + " 在 nums 中的下标为: " + index);
    }
  }
  
  main();
  