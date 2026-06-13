/**
 * 堆排序
 * 堆排序是利用堆这种数据结构而设计的一种排序算法，堆排序是一种选择排序
 * 堆是一个具有特殊性质的完全二叉树，任意非叶子节点的值大于或等于左右子
 * 节点的值，或者任意非叶子节点的值小于或等于左右子节点的值
 *
 * 最后一个非叶子节点计算公式：length / 2 - 1
 * 第 n 个下标节点的左子节点计算公式：2 * n + 1
 * 第 n 个下标节点的右子节点计算公式：2 * n + 2
 */

function heapSort(nums) {
  if (!nums) {
    return;
  }
  // 调整所有叶子节点, 构造成一个大顶堆
  // 堆顶的根节点就是序列的最大值
  for (let i = Math.floor(nums.length / 2) - 1; i >= 0; i--) {
    adjustHeap(nums, i, nums.length);
  }
  // 将堆顶的根节点和叶子节点进交换，此时叶子节点就是最大值
  for (let i = nums.length - 1; i > 0; i--) {
    nums[0] = nums[0] + nums[i];
    nums[i] = nums[0] - nums[i];
    nums[0] = nums[0] - nums[i];
    // 对于剩余的元素重新构造成大顶堆
    adjustHeap(nums, 0, i);
  }
}

/**
 * 调整堆, 使其成为大顶堆
 * i: 当前需要调整的节点下标
 * count: 调整次数
 */
function adjustHeap(nums, i, count) {
  let temp = nums[i]; // 当前节点
  for (let j = 2 * i + 1; j < count; j = 2 * j + 1) {
    // 左子节点小于右子节点
    if (j + 1 < count && nums[j] < nums[j + 1]) {
      j++; // 指向右子节点
    }
    if (nums[j] > temp) {
      // 将节点赋值给父节点
      nums[i] = nums[j];
      i = j; // 修改成下一个子节点
    } else {
      // 跳出循环，因为调整顺序为从左至右，从下至上，子树是已经调整好的堆
      break;
    }
  }
  nums[i] = temp;
}

function main() {
  let nums = [1, 7, 5, 2, 8];

  console.log("排序前: " + nums);

  heapSort(nums);

  console.log("排序后: " + nums);
}

main();
