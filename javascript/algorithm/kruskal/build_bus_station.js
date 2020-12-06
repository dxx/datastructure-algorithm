/**
 * 克鲁斯卡尔(Kruskal)算法解决建设公交站问题
 * Kruskal 算法适用于稀疏图
 * 思想：按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路
 */

/**
 * 最小生成树
 */
function MinTree(vertexes, edges) {
  let numOfVertex = vertexes.length;
  let initialVertexes = new Array(numOfVertex);
  let initialMatrix = new Array(numOfVertex);
  for (let i = 0; i < numOfVertex; i++) {
    initialVertexes[i] = vertexes[i];
    initialMatrix[i] = new Array(numOfVertex);
    for (let j = 0; j < numOfVertex; j++) {
      initialMatrix[i][j] = edges[i][j];
    }
  }
  let numOfEdge = 0;
  for (let i = 0; i < numOfVertex; i++) {
    // 已经统计过的边不统计
    for (let j = i + 1; j < numOfVertex; j++) {
      numOfEdge++;
    }
  }
  let graph = new Graph(initialVertexes, initialMatrix, numOfEdge);
  this.graph = graph;
}

function Graph(vertexes, matrix, numOfEdge) {
  this.vertexes = vertexes; // 顶点
  this.matrix = matrix; // 领接矩阵，代表边的值
  this.numOfEdge = numOfEdge; // 边的条数
}
/**
 * 边
 */
function Edge(start, end, weight) {
  this.start = start; // 起始顶点
  this.end = end; // 结束顶点
  this.weight = weight; // 边的权值
}

MinTree.prototype.kruskal = function() {
  // 保存最小生成树的边
  let edges = new Array();

  // 保存已存在最小生成树中每个顶点对应的在树中的终点下标
  let endPosits = new Array(this.graph.numOfEdge);

  // 获取所有边
  let allEdges = this.getEdges();
  console.log("======边排序前======");
  console.log(allEdges);

  // 对边按照权值从小到大进行排序
  allEdges.sort((a, b) => a.weight - b.weight);

  console.log("======边排序后======");
  console.log(allEdges);

  // 遍历所以的边
  for (let i = 0; i < allEdges.length; i++) {
    let v = allEdges[i];
    // 获取边的起始顶点下标
    let startPosit = this.getVertexPosit(v.start);
    // 获取边的结束顶点下标
    let endPosit = this.getVertexPosit(v.end);

    // 获取起始顶点的终点下标
    let endPosit1 = this.getEndPosit(endPosits, startPosit);
    // 获取结束顶点的终点下标
    let endPosit2 = this.getEndPosit(endPosits, endPosit);
    // 判断是否形成回路
    if (endPosit1 != endPosit2) { // 没有构成回路
      // 设置 endPosit1 的终点下标 endPosit2
      // [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
      endPosits[endPosit1] = endPosit2;
      edges.push(v); // 将改变加入最小生成树
    }
  }

  // 输出最小生成树
  for (let i = 0; i < edges.length; i++) {
    let edge = edges[i];
    console.log("边: " + edge.start +"-" + edge.end + " => " + edge.weight);
  }
}

/**
 * 获取顶点在顶点集合中的位置
 * @param vertex 顶点
 */
MinTree.prototype.getVertexPosit = function(vertex) {
  for (let i = 0; i < this.graph.vertexes.length; i++) {
    if (vertex == this.graph.vertexes[i]) {
      return i;
    }
  }
  return -1;
}

/**
 * 获取指定下标顶点的终点下标
 * posits: 存放顶点和对应终点下标，posits 的下标表示顶点下标，值表示对应顶点的终点下标
 * i: 目标顶点下标
 * posits = [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0], i = 4 返回 5
 */
MinTree.prototype.getEndPosit = function(posits, i) {
  while (posits[i]) {
    i = posits[i];
  }
  return i;
}

/**
 * 获取所有的边
 */
MinTree.prototype.getEdges = function() {
  let edges = new Array();
  let numOfVertex = this.graph.vertexes.length;
  for (let i = 0; i < numOfVertex; i++) {
    for (let j = i + 1; j < numOfVertex; j++) {
      // 不连通的跳过
      if (this.graph.matrix[i][j] == Number.MAX_VALUE) {
          continue;
      }
      // 创建边
      let edge = new Edge(this.graph.vertexes[i],
        this.graph.vertexes[j],
        this.graph.matrix[i][j]);
      edges.push(edge);
    }
  }
  return edges;
}

MinTree.prototype.showGraph = function() {
  for (let i = 0; i < this.graph.matrix.length; i++) {
    let edges = this.graph.matrix[i];
    let str = "[ ";
    for (let j = 0; j < edges.length; j++) {
      str += edges[j] + " ";
    }
    str += "]";
    console.log(str);
  }
}

function main() {
  let vertexes = ["A", "B", "C", "D", "E", "F", "G"];
  // 0-表示自己跟自己不连通，intMax-表示跟其它顶点不连通
  let edges = [
    [0, 12, Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 16, 14],
    [12, 0, 10, Number.MAX_VALUE, Number.MAX_VALUE, 7, Number.MAX_VALUE],
    [Number.MAX_VALUE, 10, 0, 3, 5, 6, Number.MAX_VALUE],
    [Number.MAX_VALUE, Number.MAX_VALUE, 3, 0, 4, Number.MAX_VALUE, Number.MAX_VALUE],
    [Number.MAX_VALUE, Number.MAX_VALUE, 5, 4, 0, 2, 8],
    [16, 7, 6, Number.MAX_VALUE, 2, 0, 9],
    [14, Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 8, 9, 0]
  ];
  let minTree = new MinTree(vertexes, edges);
  // minTree.showGraph();
  minTree.kruskal();
}

main();
