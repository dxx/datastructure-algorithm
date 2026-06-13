/**
 * 选择排序
 * 1.假定第一元素为最大或最小的元素
 * 2.找出最大或最小的元素的小标，循环 length - 1 次
 * 3.每次循环完成后将最大值或最小值和本次循环的第一个元素交换
 */

function selectSort(nums) {
  if (nums == null) {
    return;
  }
  let length = nums.length;
  for (let i = 0; i < length - 1; i++) {
    // 记录最小值的下标
    let minIndex = i;
    for (let j = i + 1; j < length; j++) {
      if (nums[minIndex] > nums[j]) {
         // 修改最小值下标
        minIndex = j;
      }
    }
    // 优化：判断是否需要交换
    if (minIndex != i) {
      swap(nums, i, minIndex);
    }
  }
}

function swap(nums, i, j) {
  let temp = nums[i];
  nums[i] = nums[j];
  nums[j] = temp;
}

function main() {
  let nums = [3, 5, 7, 1, 2, 4, 9, 6, 8];
  console.log("排序前: " + nums);
  selectSort(nums);
  console.log("排序后: " + nums);
}

main();
