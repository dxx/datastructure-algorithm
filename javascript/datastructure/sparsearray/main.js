/**
 * 稀疏数组
 * 在二维数组中，如果值为 0 的元素数目远远大于非 0 元素的数目，并且非 0 元素的分布没有规律，则该数组被称为稀疏数组
 * 稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：
 * 原数组中存在大量的无用的数据，占据了存储空间，真正有用的数据却很少
 * 压缩后存储可以节省存储空间，在数据序列化到磁盘时，压缩存储可以提高 IO 效率
 *
 * 二维数组
 * 0	0	3	0	0
 * 0	0	0	6	0
 * 0	1	0	0	0
 * 0	0	0	5	0
 * 0	0	0	0	0
 *
 * 稀疏数组
 * row	col	val
 * 5	5	4
 * 0	2	3
 * 1	3	6
 * 2	1	1
 * 3	3	5
 */

const fs = require("fs");

const sparseArrayFileName = "./sparse.data";

/**
* 二维数组转稀疏数组
*/
function toSparseArray(array) {
  // 统计非 0 的数量
  let count = 0;
  for (let i = 0; i < array.length; i++) {
  for (let j = 0; j < array[i].length; j++) {
    if (array[i][j] != 0) {
		count++;
      }
    }
  }

  let sparseArray = new Array(count + 1);
  for (let i = 0; i < sparseArray.length; i++) {
    // 每行 3 列
    sparseArray[i] = [0, 0, 0];
  }

  // 存储第一行信息，从左到右存储的依次是 行，列，非 0 的个数
  sparseArray[0] = [array.length, array[0].length, count];
  for (let i = 0, row = 1; i < array.length; i++) {
    for (let j = 0; j < array[i].length; j++) {
      if (array[i][j] != 0) {
      sparseArray[row][0] = i;
      sparseArray[row][1] = j;
      sparseArray[row][2] = array[i][j];
      row++;
      }
    }
  }
  return sparseArray;
}

/**
 * 稀疏数组转二维数组
 */
function toArray(sparseArray) {
  let row = sparseArray[0][0];
  let col = sparseArray[0][1];
  let array = new Array(row);
  // 初始化 0 值
  for (let i = 0; i < array.length; i++) {
    array[i] = new Array(col).fill(0);
  }
  let val;
  // 从稀疏数组中第二行开始读取数据
  for (let i = 1; i < sparseArray.length; i++) {
    row = sparseArray[i][0];
    col = sparseArray[i][1];
    val = sparseArray[i][2];
    array[row][col] = val;
  }
  return array;
}

/**
 * 存储稀疏数组
 */
function storageSparseArray(sparseArray) {
  let data = "";
  // 存储矩阵格式
  for (let i = 0; i < sparseArray.length; i++) {
    for (let j = 0; j < sparseArray[i].length; j++) {
      data += sparseArray[i][j] + "\t";
    }
    data += "\n";
  }
  fs.writeFileSync(sparseArrayFileName, data);
}

/**
 * 读取稀疏数组
 */
function readSparseArray() {
  let fileLines  = fs.readFileSync(sparseArrayFileName).toString("UTF-8").split("\n");
  // 读取第一行
  let strs = fileLines[0].split("\t");
  // 第一行信息
  let row = Number(strs[0]);
  let col = Number(strs[1]);
  let val = Number(strs[2]);
  
  let sparseArray = new Array(val + 1);
  sparseArray[0] = [row, col, val];
  // 从第二行开始
  for (let i = 1; i < fileLines.length; i++) {
    if (fileLines[i]) {
      strs = fileLines[i].split("\t");
      
      row = Number(strs[0]);
      col = Number(strs[1]);
      val = Number(strs[2]);
      sparseArray[i] = [row, col, val];
    }
  }
  return sparseArray;
}

function printArray(array) {
  let str = "";
  for (let i = 0; i < array.length; i++) {
    for (let j = 0; j < array[i].length; j++) {
      str += array[i][j] + "\t";
    }
    str += "\n";
  }
  console.log(str);
}

function main() {
  // 定义一个 5x5 的二维数组
  let array = new Array(5);
  for (let i = 0; i < array.length; i++) {
    array[i] = [0, 0, 0, 0, 0];
  }
  // 初始化 3，6， 1，5
  array[0][2] = 3;
  array[1][3] = 6;
  array[2][1] = 1;
  array[3][3] = 5;

  console.log("原二维数组：");
  printArray(array);

  // 转成稀疏数组
  let sparseArray = toSparseArray(array);

  console.log("转换后的稀疏数组：");
  printArray(sparseArray);

  // 存储稀疏数组
  storageSparseArray(sparseArray);

  // 读取稀疏数组
  sparseArray = readSparseArray();
  console.log("读取的稀疏数组：");
  printArray(sparseArray);

  // 转成二维数组
  array = toArray(sparseArray);
  console.log("转换后的二维数组：");
  printArray(array);
}

main();
