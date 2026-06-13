/**
 * 希尔排序
 * 1.计算出步长 step，step = length / 2
 * 2.从 step 开始，循环到 length
 * 3.将循环开始时的元素和比当前大 step 的元素进行比较
 * 4.发现逆序则进行交换
 */

function shellSort(nums) {
  if (nums == null) {
      return;
  }
  let length = nums.length;
  // 控制步长
  for (let step = Math.floor(length / 2); step > 0; step = Math.floor(step / 2)) {
    for (let i = step; i < length; i++) {
      for (let j = i - step; j >= 0 && nums[j] > nums[j + step]; j -= step) {
        // 前面的数比后面的数大，进行交换
        swap(nums, j, j + step);
      }
    }
  }
}

function swap(nums, i, j) {
  let temp = nums[i];
  nums[i] = nums[j];
  nums[j] = temp;
}

function main() {
  let nums = [5, 1, 7, 3, 2, 4, 9, 6, 8];
  console.log("排序前: " + nums);
  shellSort(nums);
  console.log("排序后: " + nums);
}

main();
