/**
 * 图
 */

function Graph(num) {
  this.vertexes = new Array(num);
  this.matrix = new Array(num);
  this.matrix.fill(0);
  for (let i = 0; i < this.matrix.length; i++) {
    this.matrix[i] = new Array(num);
    this.matrix[i].fill(0);
  }
  this.numOfVertex = 0;
  this.numOfEdge = 0;
}

/**
 * 添加顶点
 */
Graph.prototype.addVertex = function(vertex) {
  this.vertexes[this.numOfVertex++] = vertex;
}

/**
 * 添加边
 * i1: 第一个顶点下标
 * i2: 第二个顶点下标
 * weight: 权值. 0-表示不通, 1-表示通
 */
Graph.prototype.addEdge = function(i1, i2, weight) {
  // 在二维数组中设置权值，因为无方向图，所以两个位置都需要设置
  this.matrix[i1][i2] = weight;
  this.matrix[i2][i1] = weight;
  this.numOfEdge++;
}

/**
 * 获取顶点数量
 */
Graph.prototype.getNumOfVertex = function() {
  return this.numOfVertex;
}

/**
 * 获取边的数量
 */
Graph.prototype.getNumOfEdge = function() {
  return this.numOfEdge;
}

/**
 * 深度优先遍历
 */
Graph.prototype.dfs = function() {
  let isVisited = new Array(this.getNumOfVertex());
  // 遍历所有的顶点，进行深度优先遍历
  for (let i = 0; i < this.getNumOfVertex(); i++) {
    if (!isVisited[i]) {
      this.dfsRecursion(isVisited, i);
    }
  }
}

/**
 * 递归遍历
 * 1.访问初始顶点 v，并标记顶点 v 为已访问
 * 2.查找顶点 v 的第一个邻接顶点 w
 * 3.如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问
 * 4.如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3
 * 5.查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3
 */
Graph.prototype.dfsRecursion = function (isVisited, v) {
  console.log(this.vertexes[v]);

  // 标记已被访问
  isVisited[v] = true;

  // 获取第一个邻接顶点
  let w = this.getFirstVertex(v);
  // 存在则继续调用
  while (w != -1) {
    // 未被访问
    if (!isVisited[w]) {
      // 继续遍历
      this.dfsRecursion(isVisited, w);
    }
    // 查找顶点 v 的 w 邻接顶点的下一个邻接顶点
    w = this.getNextVertex(v, w);
  }
}

/**
 * 广度优先遍历
 */
Graph.prototype.bfs = function() {
  let isVisited = new Array(this.getNumOfVertex());
  // 遍历所有的顶点，进行广度优先遍历
  for (let i = 0; i < this.getNumOfVertex(); i++) {
    if (!isVisited[i]) {
      this.bfs2(isVisited, i);
    }
  }
}

/**
 * 1.访问初始顶点 v 并标记顶点 v 为已访问。
 * 2.顶点 v 入队列。
 * 3.当队列非空时，继续执行，否则结束。
 * 4.出队列，取得队头结点 u。
 * 5.查找结点 u 的第一个邻接顶点 w。
 * 6.若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
 * 7.若顶点 w 尚未被访问，则访问顶点 w 并标记为已访问。
 * 8.将顶点 w 入队列。
 * 9.查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。
 */
Graph.prototype.bfs2 = function(isVisited, v) {
  console.log(this.vertexes[v]);

  let queue = new Array();
  // 标记已被访问
  isVisited[v] = true;
  // 将顶点入队列
  queue.unshift(v);
  while (queue.length != 0) {
    // 取出头结点下标
    let u = queue.pop();
    // 获取第一个邻接节点的下标
    let w = this.getFirstVertex(u);
    while (w != -1) {
      // 未被访问
      if (!isVisited[w]) {
        console.log(this.vertexes[w]);
        // 标记已被访问
        isVisited[w] = true;
        // 入队列
        queue.unshift(w);
      }
      // 获取顶点 u 的继 w 邻接顶点后的下一个邻接顶点
      w = this.getNextVertex(u, w);
    }
  }
}

/**
 * 获取第一个邻接顶点下标
 */
Graph.prototype.getFirstVertex = function(i) {
  for (let j = 0; j < this.getNumOfVertex(); j++) {
    if (this.matrix[i][j] > 0) {
      return j;
    }
  }
  return -1;
}

/**
 * 获取下一个邻接顶点下标
 */
Graph.prototype.getNextVertex = function(i1, i2){
  for (let j = i2 + 1; j < this.getNumOfVertex(); j++) {
    if (this.matrix[i1][j] > 0) {
      return j;
    }
  }
  return -1;
}

/**
 * 显示邻接矩阵
 */
Graph.prototype.showEdges = function() {
  let str = "";
  for (let i = 0; i < this.matrix.length; i++) {
    str = "[";
    for (let j = 0; j < this.matrix[i].length; j++) {
      str += " " + this.matrix[i][j] + " ";
    }
    str += "]";
    console.log(str);
  }
}

function main() {
  let vertexes = ["A", "B", "C", "D", "E"];
  let graph = new Graph(5);

  for (let i = 0; i < vertexes.length; i++) {
    graph.addVertex(vertexes[i]);
  }

  // A-B
  graph.addEdge(0, 1, 1);
  // A-C
  graph.addEdge(0, 2, 1);
  // B-C
  graph.addEdge(1, 2, 1);
  // B-E
  graph.addEdge(1, 4, 1);
  // C-D
  graph.addEdge(2, 3, 1);

  graph.showEdges();

  console.log("======深度优先遍历======");
  graph.dfs();

  console.log();

  console.log("======广度优先遍历======");
  graph.bfs();
}

main();
