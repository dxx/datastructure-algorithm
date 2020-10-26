/**
 * 基数排序
 * 1.创建一个二维数组，数组长度为 10，并初始化 10 个一维数组，每个一维数组的长度为待排序数组的长度
 * 2.遍历待排序的数组，从最低位开始，求出每个元素的个位数作为二维数组的下标，将其放入到二维数组对应的数组中
 * 3.从二维数组中依次取出所有元素放入原数组中
 * 4.重复步骤 2，依次计算个位、十位、百位等，作为下标，直到待排序数组中的最大位数
 */

function radixSort(nums) {
  if (nums == null) {
    return;
  }

  let max = nums[0]; // 最大位的元素
  for (let i = 0; i < nums.length; i++) {
    let num = nums[i];
    if (max < num) {
      max = num;
    }
  }
  let maxLength = new String(max).length;

  let bucket = new Array(10); // 桶数组
  for (let i = 0; i < 10; i++) {
    bucket[i] = new Array(nums.length); // 初始化切片长度
    bucket[i].fill(0);
  }
  let order = new Array(10); // 存放每个桶真实存放数据的长度
  order.fill(0);

  let n = 1; // 控制元素的位数
  for (let i = 0; i < maxLength; i++) {
    for (let j = 0; j < nums.length; j++){
      let num = nums[j];
      let bucketIndex = Math.floor(num / n) % 10; // 计算桶的下标
      bucket[bucketIndex][order[bucketIndex]] = num;
      order[bucketIndex]++; // 尾下标 + 1
    }

    let numIndex = 0;
    // 从桶数组中依次取出所有元素放入原数组
    for (let orderIndex = 0; orderIndex < order.length; orderIndex++) {
      let bucketLength = order[orderIndex];
      if (bucketLength != 0) {
        for (let j = 0; j < bucketLength; j++) {
          nums[numIndex] = bucket[orderIndex][j];
          numIndex++;
        }
        order[orderIndex] = 0; // 重置当前桶下标
      }
    }

    n *= 10;
  }
}

function main() {
  let nums = [5, 1, 7, 13, 21, 32, 9, 66, 8, 20];
  console.log("排序前: " + nums);
  radixSort(nums);
  console.log("排序后: " + nums);
}

main();
