/**
 * 马踏棋盘算法
 * 国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，
 * 按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最
 * 终使得“马”走遍棋盘 64 个方格
 */

/**
 * 棋盘
 */
function Chessboard(row, col) {
  let length = row * col;
  this.row = row; // 表示棋盘的行数
  this.col = col; // 表示棋盘的列数
  this.visited = new Array(length); // 标记点是否被访问过
  this.visited.fill(false);
  this.steps = []; // 存放步数
  for (let i = 0; i < row; i++) {
    this.steps[i] = [];
    for (let j = 0; j < col; j++) {
      this.steps[i][j] = 0;
    }
  }
  this.finished = false; // 表示是否已经走完
}

/**
 * 位置
 */
function Point(x, y) {
  this.x = x; // X 下标
  this.y = y; // Y 下标
}

/**
 * 移动
 * startX 起始横坐标。从 0 开始
 * startY 起始纵坐标。从 0 开始
 */
Chessboard.prototype.move = function(startX, startY) {
  console.log("开始走马踏棋");
  let startTime = new Date().getTime();

  this.traversal(startX, startY, 1);

  console.log("马踏棋结束");
  let endTime = new Date().getTime();
  console.log("耗时:%fs", (endTime - startTime) / 1000);

  for (let i = 0; i < this.row; i++) {
    let str = "";
    for (let j = 0; j < this.col; j++) {
      str += this.steps[i][j] + " ";
    }
    console.log(str);
  }
}

/**
 * 骑士周游算法
 */
Chessboard.prototype.traversal = function(x, y, step) {
  // 将当前位置标记已访问。y = 4, col = 8, x = 4 => 4 * 8 + 4 = 36
  this.visited[y * this.col + x] = true;
  // 记录步数
  this.steps[y][x] = step;
  // 获取下一步可以走的所有位置
  let points = this.nextPoints(new Point(x, y));
  // 排序优化，优先选择下一步最少可走数目的位置，减少回溯，体现出贪心算法的特点
  // 去掉此方法，算法的耗时会很久，根据走法不同，结果也会不同
  points.sort((a, b) => this.nextPoints(a).length - this.nextPoints(b).length);

  while (points.length != 0) {
      // 取出第一个点
      let p = points.shift();
      // 该点未被访问过
      if (!this.visited[p.y * this.col + p.x]) {
          // 继续往下走
          this.traversal(p.x, p.y, step + 1);
      }
  }
  // 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置
  if (step < this.row * this.col && !this.finished) {
      this.visited[y * this.col + x] = false;
      this.steps[y][x] = 0;
  } else {
      this.finished = true;
  }
}

/**
 * 获取当位置的下一步可走位置的集合，最多可有 8 个位置
 */
Chessboard.prototype.nextPoints = function(point) {
  let points = [];
  // 判断是否可以走 1 的位置
  if (point.x - 1 >= 0 && point.y - 2 >= 0) {
    points.push(new Point(point.x - 1, point.y - 2));
  }
  // 判断是否可以走 2 的位置
  if (point.x + 1 < this.col && point.y - 2 >= 0) {
    points.push(new Point(point.x + 1, point.y - 2));
  }
  // 判断是否可以走 3 的位置
  if (point.x + 2 < this.col && point.y - 1 >= 0) {
    points.push(new Point(point.x + 2, point.y - 1));
  }
  // 判断是否可以走 4 的位置
  if (point.x + 2 < this.col && point.y + 1 < this.row) {
    points.push(new Point(point.x + 2, point.y + 1));
  }
  // 判断是否可以走 5 的位置
  if (point.x + 1 < this.col && point.y + 2 < this.row) {
    points.push(new Point(point.x + 1, point.y + 2));
  }
  // 判断是否可以走 6 的位置
  if (point.x - 1 >= 0 && point.y + 2 < this.row) {
    points.push(new Point(point.x - 1, point.y + 2));
  }
  // 判断是否可以走 7 的位置
  if (point.x - 2 >= 0 && point.y + 1 < this.row) {
    points.push(new Point(point.x - 2, point.y + 1));
  }
  // 判断是否可以走 8 的位置
  if (point.x - 2 >= 0 && point.y - 1 >= 0) {
    points.push(new Point(point.x - 2, point.y - 1));
  }
  return points;
}

function main() {
  let chessboard = new Chessboard(8, 8);
  chessboard.move(4, 4);
}

 main();
 